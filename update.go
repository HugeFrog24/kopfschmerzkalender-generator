package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/blang/semver/v4"
)

var (
	CurrentVersion     = semver.MustParse("1.0.1")
	ErrUpdateCancelled = errors.New("update cancelled by user")
)

func GetCurrentVersion() semver.Version {
	return CurrentVersion
}

type GithubRelease struct {
	TagName string `json:"tag_name"`
	HTMLURL string `json:"html_url"`
	Assets  []struct {
		BrowserDownloadURL string `json:"browser_download_url"`
	} `json:"assets"`
}

func CheckForUpdates(cancelChan <-chan struct{}) (semver.Version, string, error) {
	log.Println("Checking for updates...")

	respChan := make(chan *http.Response)
	errChan := make(chan error)

	go func() {
		resp, err := http.Get(fmt.Sprintf("https://api.github.com/repos/%s/releases/latest", strings.TrimPrefix(GithubRepoURL, "https://github.com/")))
		if err != nil {
			errChan <- err
			return
		}
		respChan <- resp
	}()

	select {
	case <-cancelChan:
		return semver.Version{}, "", ErrUpdateCancelled
	case err := <-errChan:
		log.Printf("Error fetching latest release: %v", err)
		return semver.Version{}, "", err
	case resp := <-respChan:
		defer resp.Body.Close()

		log.Println("Successfully fetched latest release information")
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Error reading response body: %v", err)
			return semver.Version{}, "", err
		}

		var release GithubRelease
		err = json.Unmarshal(body, &release)
		if err != nil {
			log.Printf("Error unmarshalling JSON: %v", err)
			return semver.Version{}, "", err
		}

		latestVersion, err := semver.Parse(release.TagName[1:]) // Remove 'v' prefix
		if err != nil {
			log.Printf("Error parsing version: %v", err)
			return semver.Version{}, "", err
		}

		log.Printf("Latest version: %s", latestVersion)

		downloadURL := release.HTMLURL
		if len(release.Assets) > 0 {
			downloadURL = release.Assets[0].BrowserDownloadURL
		}
		log.Printf("Download URL: %s", downloadURL)

		return latestVersion, downloadURL, nil
	}
}

func DownloadAndInstallUpdate(url string, progressCallback func(float64), cancelChan <-chan struct{}) error {
	log.Println("Starting download and installation of update...")

	execPath, err := os.Executable()
	if err != nil {
		log.Printf("Error getting executable path: %v", err)
		return err
	}
	log.Printf("Current executable path: %s", execPath)

	updatePath := execPath + ".new"
	log.Printf("Update will be downloaded to: %s", updatePath)

	if err := downloadFile(url, updatePath, progressCallback, cancelChan); err != nil {
		if err == ErrUpdateCancelled {
			log.Println("Update download cancelled by user")
			return err
		}
		log.Printf("Error downloading update: %v", err)
		return err
	}
	log.Println("Update downloaded successfully")

	oldPath := execPath + ".old"
	log.Printf("Renaming current executable to: %s", oldPath)
	if err := os.Rename(execPath, oldPath); err != nil {
		log.Printf("Error renaming current executable: %v", err)
		return err
	}

	log.Printf("Renaming new executable to: %s", execPath)
	if err := os.Rename(updatePath, execPath); err != nil {
		log.Printf("Error renaming new executable: %v", err)
		// If renaming fails, try to restore the old executable
		log.Println("Attempting to restore old executable")
		os.Rename(oldPath, execPath)
		return err
	}

	log.Println("Update installed successfully")
	return nil
}

func downloadFile(url, filepath string, progressCallback func(float64), cancelChan <-chan struct{}) error {
	log.Printf("Downloading file from %s to %s", url, filepath)

	respChan := make(chan *http.Response)
	errChan := make(chan error)

	go func() {
		resp, err := http.Get(url)
		if err != nil {
			errChan <- err
			return
		}
		respChan <- resp
	}()

	var resp *http.Response
	select {
	case <-cancelChan:
		return ErrUpdateCancelled
	case err := <-errChan:
		log.Printf("Error fetching file: %v", err)
		return err
	case resp = <-respChan:
		// Continue with the download
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		log.Printf("Error creating file: %v", err)
		return err
	}
	defer out.Close()

	counter := &WriteCounter{
		Total:             resp.ContentLength,
		ProgressCallback:  progressCallback,
		LastProgressValue: -1,
	}

	done := make(chan error, 1)
	go func() {
		_, err := io.Copy(out, io.TeeReader(resp.Body, counter))
		done <- err
	}()

	select {
	case <-cancelChan:
		return ErrUpdateCancelled
	case err := <-done:
		if err != nil {
			log.Printf("Error copying file contents: %v", err)
			return err
		}
	}

	log.Println("File downloaded successfully")
	return nil
}

type WriteCounter struct {
	Total             int64
	Current           int64
	ProgressCallback  func(float64)
	LastProgressValue float64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Current += int64(n)
	progress := float64(wc.Current) / float64(wc.Total)
	if progress-wc.LastProgressValue >= 0.01 { // Update every 1%
		wc.ProgressCallback(progress)
		wc.LastProgressValue = progress
	}
	return n, nil
}
