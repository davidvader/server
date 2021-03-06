// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package sqlite

const (
	// ListBuilds represents a query to
	// list all builds in the database.
	ListBuilds = `
SELECT *
FROM builds;
`

	// ListRepoBuilds represents a query to list
	// all builds for a repo_id in the database.
	ListRepoBuilds = `
SELECT *
FROM builds
WHERE repo_id = ?
ORDER BY id DESC
LIMIT ?
OFFSET ?;
`

	// ListRepoBuildsByEvent represents a query to select
	// a build for a repo_id with a specific event type
	// in the database.
	ListRepoBuildsByEvent = `
SELECT *
FROM builds
WHERE repo_id = ?
AND event = ?
ORDER BY number DESC
LIMIT ?
OFFSET ?;
`

	// SelectRepoBuild represents a query to select
	// a build for a repo_id in the database.
	SelectRepoBuild = `
SELECT *
FROM builds
WHERE repo_id = ?
AND number = ?
LIMIT 1;
`

	// SelectLastRepoBuild represents a query to select
	// the last build for a repo_id in the database.
	SelectLastRepoBuild = `
SELECT *
FROM builds
WHERE repo_id = ?
ORDER BY number DESC
LIMIT 1;
`
	// SelectLastRepoBuildByBranch represents a query to
	// select the last build for a repo_id and branch name
	// in the database
	SelectLastRepoBuildByBranch = `
SELECT *
FROM builds
WHERE repo_id = ?
AND branch = ?
ORDER BY number DESC
LIMIT 1;
`

	// SelectBuildsCount represents a query to select
	// the count of builds in the database.
	SelectBuildsCount = `
SELECT count(*) as count
FROM builds;
`

	// SelectRepoBuildCount represents a query to select
	// the count of builds for a repo_id in the database.
	SelectRepoBuildCount = `
SELECT count(*) as count
FROM builds
WHERE repo_id = ?;
`

	// SelectRepoBuildCountByEvent represents a query to select
	// the count of builds for by repo and event type in the database.
	SelectRepoBuildCountByEvent = `
SELECT count(*) as count
FROM builds
WHERE repo_id = ?
AND event = ?;
`

	// SelectBuildsCountByStatus represents a query to select
	// the count of builds for a status in the database.
	SelectBuildsCountByStatus = `
SELECT count(*) as count
FROM builds
WHERE status = ?;
`

	// DeleteBuild represents a query to
	// remove a build from the database.
	DeleteBuild = `
DELETE
FROM builds
WHERE id = ?;
`
)

// createBuildService is a helper function to return
// a service for interacting with the builds table.
func createBuildService() *Service {
	return &Service{
		List: map[string]string{
			"all":         ListBuilds,
			"repo":        ListRepoBuilds,
			"repoByEvent": ListRepoBuildsByEvent,
		},
		Select: map[string]string{
			"repo":                SelectRepoBuild,
			"last":                SelectLastRepoBuild,
			"lastByBranch":        SelectLastRepoBuildByBranch,
			"count":               SelectBuildsCount,
			"countByStatus":       SelectBuildsCountByStatus,
			"countByRepo":         SelectRepoBuildCount,
			"countByRepoAndEvent": SelectRepoBuildCountByEvent,
		},
		Delete: DeleteBuild,
	}
}
