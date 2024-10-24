// utils/github_api.go
package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const GitHubAPI = "https://api.github.com"

// Repo 表示 GitHub 上的一个仓库
type Repo struct {
	Name        string `json:"name"`
	Stars       int    `json:"stargazers_count"`
	Forks       int    `json:"forks_count"`
	OpenIssues  int    `json:"open_issues_count"`
	Contributor string `json:"owner.login"`
}

// User 表示 GitHub 用户
type User struct {
	Login    string `json:"login"`
	Location string `json:"location"`
}

// GetUserRepos 获取指定用户的仓库
func GetUserRepos(username string) ([]Repo, error) {
	resp, err := http.Get(fmt.Sprintf("%s/users/%s/repos", GitHubAPI, username))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get repos: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var repos []Repo
	if err := json.Unmarshal(body, &repos); err != nil {
		return nil, err
	}
	return repos, nil
}

// GetUserFollowers 获取指定用户的 followers
func GetUserFollowers(username string) ([]User, error) {
	resp, err := http.Get(fmt.Sprintf("%s/users/%s/followers", GitHubAPI, username))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get followers: %s", resp.Status)
	}

	var followers []User
	if err := json.NewDecoder(resp.Body).Decode(&followers); err != nil {
		return nil, err
	}
	return followers, nil
}

// GetUserFollowing 获取指定用户的 following
func GetUserFollowing(username string) ([]User, error) {
	resp, err := http.Get(fmt.Sprintf("%s/users/%s/following", GitHubAPI, username))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get following: %s", resp.Status)
	}

	var following []User
	if err := json.NewDecoder(resp.Body).Decode(&following); err != nil {
		return nil, err
	}
	return following, nil
}

// GetUserProfile 获取指定用户的个人信息
func GetUserProfile(username string) (User, error) {
	resp, err := http.Get(fmt.Sprintf("%s/users/%s", GitHubAPI, username))
	if err != nil {
		return User{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return User{}, fmt.Errorf("failed to get user profile: %s", resp.Status)
	}

	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return User{}, err
	}
	return user, nil
}
