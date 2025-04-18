package services

import (
	"os"
	"testing"

	internal "drupdater/internal"
	"drupdater/internal/codehosting"
	"drupdater/internal/utils"

	object "github.com/go-git/go-git/v5/plumbing/object"
	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

func TestGetFixedAdvisories(t *testing.T) {
	ws := &WorkflowSecurityUpdateService{}

	tests := []struct {
		name     string
		before   []Advisory
		after    []Advisory
		expected []Advisory
	}{
		{
			name: "No advisories fixed",
			before: []Advisory{
				{CVE: "CVE-1234", Title: "Vulnerability 1"},
				{CVE: "CVE-5678", Title: "Vulnerability 2"},
			},
			after: []Advisory{
				{CVE: "CVE-1234", Title: "Vulnerability 1"},
				{CVE: "CVE-5678", Title: "Vulnerability 2"},
			},
			expected: []Advisory{},
		},
		{
			name: "Some advisories fixed",
			before: []Advisory{
				{CVE: "CVE-1234", Title: "Vulnerability 1"},
				{CVE: "CVE-5678", Title: "Vulnerability 2"},
			},
			after: []Advisory{
				{CVE: "CVE-1234", Title: "Vulnerability 1"},
			},
			expected: []Advisory{
				{CVE: "CVE-5678", Title: "Vulnerability 2"},
			},
		},
		{
			name: "All advisories fixed",
			before: []Advisory{
				{CVE: "CVE-1234", Title: "Vulnerability 1"},
				{CVE: "CVE-5678", Title: "Vulnerability 2"},
			},
			after: []Advisory{},
			expected: []Advisory{
				{CVE: "CVE-1234", Title: "Vulnerability 1"},
				{CVE: "CVE-5678", Title: "Vulnerability 2"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := ws.GetFixedAdvisories(tt.before, tt.after)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestSecurityUpdateStartUpdate(t *testing.T) {
	logger := zap.NewNop()
	installer := NewMockInstallerService(t)
	updater := NewMockUpdaterService(t)
	repositoryService := NewMockRepositoryService(t)
	vcsProviderFactory := codehosting.NewMockVcsProviderFactory(t)
	vcsProvider := codehosting.NewMockPlatform(t)
	repository := internal.NewMockRepository(t)
	commandExecutor := utils.NewMockCommandExecutor(t)
	composerService := NewMockComposerService(t)
	config := internal.Config{
		RepositoryURL: "https://example.com/repo.git",
		Branch:        "main",
		Token:         "token",
		Sites:         []string{"site1", "site2"},
		DryRun:        false,
	}

	workflowService := newWorkflowSecurityUpdateService(logger, installer, updater, repositoryService, vcsProviderFactory, config, commandExecutor, composerService)

	worktree := internal.NewMockWorktree(t)
	worktree.On("Checkout", mock.Anything).Return(nil)

	installer.On("InstallDrupal", config.RepositoryURL, config.Branch, config.Token, config.Sites).Return(nil)
	repositoryService.On("CloneRepository", config.RepositoryURL, config.Branch, config.Token).Return(repository, worktree, "/tmp", nil)
	repositoryService.On("GetHeadCommit", repository).Return(&object.Commit{}, nil)
	repositoryService.On("BranchExists", mock.Anything, "security-update-ddd").Return(false, nil)

	updater.On("UpdateDependencies", "/tmp", []string{"package1"}, mock.Anything, true).Return(DependencyUpdateReport{}, nil)
	updater.On("UpdateDrupal", "/tmp", mock.Anything, config.Sites).Return(UpdateHooksPerSite{}, nil)
	vcsProviderFactory.On("Create", "https://example.com/repo.git", "token").Return(vcsProvider)

	fixture, _ := os.ReadFile("testdata/security_update.md")
	vcsProvider.On("CreateMergeRequest", mock.Anything, string(fixture), mock.Anything, config.Branch).Return(codehosting.MergeRequest{}, nil)
	repository.On("Push", mock.Anything).Return(nil)
	commandExecutor.On("GenerateDiffTable", mock.Anything, mock.Anything, true).Return("Dummy Table", nil)
	composerService.On("RunComposerAudit", "/tmp").Return(ComposerAudit{
		Advisories: []Advisory{
			{CVE: "CVE-1234", Title: "Vul 1", Severity: "high    ", Link: "https://example.com", PackageName: "package1"},
			{CVE: "CVE-5678", Title: "Vul 2", Severity: "high    ", Link: "https://example.com", PackageName: "package1"},
		},
	}, nil)
	composerService.On("GetComposerLockHash", "/tmp").Return("ddd", nil)

	err := workflowService.StartUpdate()

	assert.NoError(t, err)
	installer.AssertExpectations(t)
	repositoryService.AssertExpectations(t)
	updater.AssertExpectations(t)
	vcsProviderFactory.AssertExpectations(t)
	vcsProvider.AssertExpectations(t)
}

func TestSecurityUpdateStartUpdateWithDryRun(t *testing.T) {
	logger := zap.NewNop()
	installer := NewMockInstallerService(t)
	updater := NewMockUpdaterService(t)
	repositoryService := NewMockRepositoryService(t)
	vcsProviderFactory := codehosting.NewMockVcsProviderFactory(t)
	vcsProvider := codehosting.NewMockPlatform(t)
	repository := internal.NewMockRepository(t)
	commandExecutor := utils.NewMockCommandExecutor(t)
	composerService := NewMockComposerService(t)
	config := internal.Config{
		RepositoryURL: "https://example.com/repo.git",
		Branch:        "main",
		Token:         "token",
		Sites:         []string{"site1", "site2"},
		DryRun:        true,
	}

	workflowService := newWorkflowSecurityUpdateService(logger, installer, updater, repositoryService, vcsProviderFactory, config, commandExecutor, composerService)

	worktree := internal.NewMockWorktree(t)
	worktree.On("Checkout", mock.Anything).Return(nil)
	installer.On("InstallDrupal", config.RepositoryURL, config.Branch, config.Token, config.Sites).Return(nil)
	repositoryService.On("BranchExists", mock.Anything, "security-update-ddd").Return(false, nil)
	repositoryService.On("CloneRepository", config.RepositoryURL, config.Branch, config.Token).Return(repository, worktree, "/tmp", nil)
	repositoryService.On("GetHeadCommit", repository).Return(&object.Commit{}, nil)
	updater.On("UpdateDependencies", "/tmp", []string{"package1"}, mock.Anything, true).Return(DependencyUpdateReport{}, nil)
	updater.On("UpdateDrupal", "/tmp", mock.Anything, config.Sites).Return(UpdateHooksPerSite{}, nil)
	commandExecutor.On("GenerateDiffTable", mock.Anything, mock.Anything, true).Return("foo", nil)
	composerService.On("RunComposerAudit", "/tmp").Return(ComposerAudit{
		Advisories: []Advisory{
			{CVE: "CVE-1234", Title: "Vul 1", Severity: "high    ", Link: "https://example.com", PackageName: "package1"},
			{CVE: "CVE-5678", Title: "Vul 2", Severity: "high    ", Link: "https://example.com", PackageName: "package1"},
		},
	}, nil)
	composerService.On("GetComposerLockHash", "/tmp").Return("ddd", nil)

	err := workflowService.StartUpdate()

	assert.NoError(t, err)
	installer.AssertExpectations(t)
	repositoryService.AssertExpectations(t)
	updater.AssertExpectations(t)
	vcsProviderFactory.AssertExpectations(t)
	vcsProvider.AssertExpectations(t)
}
