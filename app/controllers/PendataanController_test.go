package controllers

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestPendataanRepo_CreateAlat(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		repo *PendataanRepo
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.repo.CreateAlat(tt.args.c)
		})
	}
}
