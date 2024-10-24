// services/evaluation.go
package services

import (
	"DevRank/models"
)

func CalculateTalentRank(developer models.Developer) float64 {
	// 根据项目重要性和贡献度计算 TalentRank
	rank := 0.0
	for _, repo := range developer.Repos {
		rank += float64(repo.Stars) * 0.5      // 假设星级占50%
		rank += float64(repo.Forks) * 0.3      // 假设 Forks 占30%
		rank += float64(repo.OpenIssues) * 0.2 // 假设 Issues 占20%
	}
	return rank
}
