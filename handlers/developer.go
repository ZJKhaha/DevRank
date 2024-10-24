package handlers

import (
	"DevRank/models"
	"DevRank/services"
	"DevRank/utils"
	"encoding/json"
	"net/http"
)

func GetDeveloper(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")

	// 获取用户的仓库
	repos, err := utils.GetUserRepos(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 获取用户的个人信息
	userProfile, err := utils.GetUserProfile(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 首先使用 profile 中的 location
	guessedLocation := userProfile.Location
	originalLocation := userProfile.Location // 原始 location

	// 如果 location 为空，则通过 followers 和 following 推测
	if guessedLocation == "" {
		followers, err := utils.GetUserFollowers(username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		following, err := utils.GetUserFollowing(username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// 统计 location 出现次数
		locationCount := make(map[string]int)
		for _, user := range followers {
			if user.Location != "" {
				locationCount[user.Location]++
			}
		}
		for _, user := range following {
			if user.Location != "" {
				locationCount[user.Location]++
			}
		}

		// 找到出现次数最多的 location
		maxCount := 0
		for location, count := range locationCount {
			if count > maxCount {
				maxCount = count
				guessedLocation = location
			}
		}
	}

	// 将 utils.Repo 转换为 models.Repo
	var modelRepos []models.Repo
	for _, repo := range repos {
		modelRepo := models.Repo{
			Name:       repo.Name,
			Stars:      repo.Stars,
			Forks:      repo.Forks,
			OpenIssues: repo.OpenIssues,
		}
		modelRepos = append(modelRepos, modelRepo)
	}

	// 创建 Developer 实例
	developer := models.Developer{
		Username:         username,
		Repos:            modelRepos,
		Nation:           guessedLocation,  // 设置国籍
		OriginalLocation: originalLocation, // 原始 location
	}

	// 计算 TalentRank
	developer.TalentRank = services.CalculateTalentRank(developer)

	// 返回 JSON 响应
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(developer)
}
