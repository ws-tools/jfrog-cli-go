package tests

import "github.com/jfrogdev/jfrog-cli-go/utils/ioutils"

const (
	Repo1 = "jfrog-cli-tests-repo"
	Repo2 = "jfrog-cli-tests-repo1"
	Out = "out"
	DownloadSpec = "download_spec.json"
	SimpleUploadSpec = "simple_upload_spec.json"
	UploadSpec = "upload_spec.json"
	DeleteSpec = "delete_spec.json"
	DeleteComplexSpec = "delete_complex_spec.json"
	MoveCopyDeleteSpec = "move_copy_delete_spec.json"
	PrepareCopy = "prepare_copy.json"
	Search = "search.json"
	SearchMoveDeleteRepoSpec = "search_move_delete_repo_spec.json"
	MoveRepositoryConfig = "move_repository_config.json"
	SpecsTestRepositoryConfig = "specs_test_repository_config.json"

	RepoDetailsUrl = "api/repositories/"
	SpecsCommand = "jfrog rt %v --spec=%v --url=%v --user=%v --password=%v"
)

var SimpleUploadExpected = []string{
	Repo1 + "/flat_recursive/a3.in",
	Repo1 + "/flat_recursive/a1.in",
	Repo1 + "/flat_recursive/a2.in",
	Repo1 + "/flat_recursive/b2.in",
	Repo1 + "/flat_recursive/b3.in",
	Repo1 + "/flat_recursive/b1.in",
	Repo1 + "/flat_recursive/c2.in",
	Repo1 + "/flat_recursive/c1.in",
	Repo1 + "/flat_recursive/c3.in",
}

var MassiveMoveExpected = []string{
	Repo2 + "/3_only_flat_recursive_target/a3.in",
	Repo2 + "/3_only_flat_recursive_target/b3.in",
	Repo2 + "/3_only_flat_recursive_target/c3.in",
	Repo2 + "/defaults_recursive_nonFlat_target/defaults_recursive_nonFlat_source/a/a1.in",
	Repo2 + "/defaults_recursive_nonFlat_target/defaults_recursive_nonFlat_source/a/a2.in",
	Repo2 + "/defaults_recursive_nonFlat_target/defaults_recursive_nonFlat_source/a/a3.in",
	Repo2 + "/defaults_recursive_nonFlat_target/defaults_recursive_nonFlat_source/a/b/b1.in",
	Repo2 + "/defaults_recursive_nonFlat_target/defaults_recursive_nonFlat_source/a/b/b2.in",
	Repo2 + "/defaults_recursive_nonFlat_target/defaults_recursive_nonFlat_source/a/b/b3.in",
	Repo2 + "/defaults_recursive_nonFlat_target/defaults_recursive_nonFlat_source/a/b/c/c1.in",
	Repo2 + "/defaults_recursive_nonFlat_target/defaults_recursive_nonFlat_source/a/b/c/c2.in",
	Repo2 + "/defaults_recursive_nonFlat_target/defaults_recursive_nonFlat_source/a/b/c/c3.in",
	Repo2 + "/flat_nonrecursive_target/a1.in",
	Repo2 + "/flat_nonrecursive_target/a2.in",
	Repo2 + "/flat_nonrecursive_target/a3.in",
	Repo2 + "/flat_recursive_target/a1.in",
	Repo2 + "/flat_recursive_target/a2.in",
	Repo2 + "/flat_recursive_target/a3.in",
	Repo2 + "/flat_recursive_target/b1.in",
	Repo2 + "/flat_recursive_target/b2.in",
	Repo2 + "/flat_recursive_target/b3.in",
	Repo2 + "/flat_recursive_target/c1.in",
	Repo2 + "/flat_recursive_target/c2.in",
	Repo2 + "/flat_recursive_target/c3.in",
	Repo2 + "/no_target/a/a1.in",
	Repo2 + "/no_target/a/a2.in",
	Repo2 + "/no_target/a/a3.in",
	Repo2 + "/no_target/a/b/b1.in",
	Repo2 + "/no_target/a/b/b2.in",
	Repo2 + "/no_target/a/b/b3.in",
	Repo2 + "/no_target/a/b/c/c1.in",
	Repo2 + "/no_target/a/b/c/c2.in",
	Repo2 + "/no_target/a/b/c/c3.in",
	Repo2 + "/nonflat_nonrecursive_target/nonflat_nonrecursive_source/a/a1.in",
	Repo2 + "/nonflat_nonrecursive_target/nonflat_nonrecursive_source/a/a2.in",
	Repo2 + "/nonflat_nonrecursive_target/nonflat_nonrecursive_source/a/a3.in",
	Repo2 + "/nonflat_recursive_target/nonflat_recursive_source/a/a1.in",
	Repo2 + "/nonflat_recursive_target/nonflat_recursive_source/a/a2.in",
	Repo2 + "/nonflat_recursive_target/nonflat_recursive_source/a/a3.in",
	Repo2 + "/nonflat_recursive_target/nonflat_recursive_source/a/b/b1.in",
	Repo2 + "/nonflat_recursive_target/nonflat_recursive_source/a/b/b2.in",
	Repo2 + "/nonflat_recursive_target/nonflat_recursive_source/a/b/b3.in",
	Repo2 + "/nonflat_recursive_target/nonflat_recursive_source/a/b/c/c1.in",
	Repo2 + "/nonflat_recursive_target/nonflat_recursive_source/a/b/c/c2.in",
	Repo2 + "/nonflat_recursive_target/nonflat_recursive_source/a/b/c/c3.in",
	Repo2 + "/pattern_placeholder_target/a/a1.in",
	Repo2 + "/pattern_placeholder_target/a/a2.in",
	Repo2 + "/pattern_placeholder_target/a/a3.in",
	Repo2 + "/pattern_placeholder_target/a/b/b1.in",
	Repo2 + "/pattern_placeholder_target/a/b/b2.in",
	Repo2 + "/pattern_placeholder_target/a/b/b3.in",
	Repo2 + "/pattern_placeholder_target/a/b/c/c1.in",
	Repo2 + "/pattern_placeholder_target/a/b/c/c2.in",
	Repo2 + "/pattern_placeholder_target/a/b/c/c3.in",
	Repo2 + "/properties_target/properties_source/a/a1.in",
	Repo2 + "/properties_target/properties_source/a/a2.in",
	Repo2 + "/properties_target/properties_source/a/a3.in",
	Repo2 + "/properties_target/properties_source/a/b/b1.in",
	Repo2 + "/properties_target/properties_source/a/b/b2.in",
	Repo2 + "/properties_target/properties_source/a/b/b3.in",
	Repo2 + "/properties_target/properties_source/a/b/c/c1.in",
	Repo2 + "/properties_target/properties_source/a/b/c/c2.in",
	Repo2 + "/properties_target/properties_source/a/b/c/c3.in",
	Repo2 + "/rename_target/RENAMED.in",
	Repo2 + "/simple_placeholder_target/a/a1.in",
	Repo2 + "/simple_target/a1.in",
}

var MassiveDownload = []string{
	Out + ioutils.GetFileSeperator() + "a1.in",
	Out + ioutils.GetFileSeperator() + "a2.in",
	Out + ioutils.GetFileSeperator() + "a3.in",
	Out + ioutils.GetFileSeperator() + "download",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "3_only_flat_recursive" + ioutils.GetFileSeperator() + "a3.in",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "3_only_flat_recursive" + ioutils.GetFileSeperator() + "b3.in",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "3_only_flat_recursive" + ioutils.GetFileSeperator() + "c3.in",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "aql" + ioutils.GetFileSeperator() + "downloadTestResources" + ioutils.GetFileSeperator() + "a",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "aql" + ioutils.GetFileSeperator() + "downloadTestResources" + ioutils.GetFileSeperator() + "a" + ioutils.GetFileSeperator() + "a1.in",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "aql_flat" + ioutils.GetFileSeperator() + "a1.in",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "aql_flat" + ioutils.GetFileSeperator() + "a2.in",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "aql_flat" + ioutils.GetFileSeperator() + "a3.in",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "aql_flat" + ioutils.GetFileSeperator() + "b1.in",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "aql_flat" + ioutils.GetFileSeperator() + "b2.in",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "aql_flat" + ioutils.GetFileSeperator() + "b3.in",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "aql_flat" + ioutils.GetFileSeperator() + "c1.in",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "aql_flat" + ioutils.GetFileSeperator() + "c2.in",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "aql_flat" + ioutils.GetFileSeperator() + "c3.in",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "defaults_recursive_nonFlat" + ioutils.GetFileSeperator() + "downloadTestResources" + ioutils.GetFileSeperator() + "a",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "defaults_recursive_nonFlat" + ioutils.GetFileSeperator() + "downloadTestResources" + ioutils.GetFileSeperator() + "a" + ioutils.GetFileSeperator() + "a1.in",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "defaults_recursive_nonFlat" + ioutils.GetFileSeperator() + "downloadTestResources" + ioutils.GetFileSeperator() + "a" + ioutils.GetFileSeperator() + "a2.in",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "defaults_recursive_nonFlat" + ioutils.GetFileSeperator() + "downloadTestResources" + ioutils.GetFileSeperator() + "a" + ioutils.GetFileSeperator() + "a3.in",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "defaults_recursive_nonFlat" + ioutils.GetFileSeperator() + "downloadTestResources" + ioutils.GetFileSeperator() + "a" + ioutils.GetFileSeperator() + "b",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "defaults_recursive_nonFlat" + ioutils.GetFileSeperator() + "downloadTestResources" + ioutils.GetFileSeperator() + "a" + ioutils.GetFileSeperator() + "b" + ioutils.GetFileSeperator() + "b1.in",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "defaults_recursive_nonFlat" + ioutils.GetFileSeperator() + "downloadTestResources" + ioutils.GetFileSeperator() + "a" + ioutils.GetFileSeperator() + "b" + ioutils.GetFileSeperator() + "b2.in",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "defaults_recursive_nonFlat" + ioutils.GetFileSeperator() + "downloadTestResources" + ioutils.GetFileSeperator() + "a" + ioutils.GetFileSeperator() + "b" + ioutils.GetFileSeperator() + "b3.in",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "defaults_recursive_nonFlat" + ioutils.GetFileSeperator() + "downloadTestResources" + ioutils.GetFileSeperator() + "a" + ioutils.GetFileSeperator() + "b" + ioutils.GetFileSeperator() + "c",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "defaults_recursive_nonFlat" + ioutils.GetFileSeperator() + "downloadTestResources" + ioutils.GetFileSeperator() + "a" + ioutils.GetFileSeperator() + "b" + ioutils.GetFileSeperator() + "c" + ioutils.GetFileSeperator() + "c1.in",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "defaults_recursive_nonFlat" + ioutils.GetFileSeperator() + "downloadTestResources" + ioutils.GetFileSeperator() + "a" + ioutils.GetFileSeperator() + "b" + ioutils.GetFileSeperator() + "c" + ioutils.GetFileSeperator() + "c2.in",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "defaults_recursive_nonFlat" + ioutils.GetFileSeperator() + "downloadTestResources" + ioutils.GetFileSeperator() + "a" + ioutils.GetFileSeperator() + "b" + ioutils.GetFileSeperator() + "c" + ioutils.GetFileSeperator() + "c3.in",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "flat_nonrecursive" + ioutils.GetFileSeperator() + "a1.in",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "flat_nonrecursive" + ioutils.GetFileSeperator() + "a2.in",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "flat_nonrecursive" + ioutils.GetFileSeperator() + "a3.in",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "flat_recursive" + ioutils.GetFileSeperator() + "a1.in",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "flat_recursive" + ioutils.GetFileSeperator() + "a2.in",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "flat_recursive" + ioutils.GetFileSeperator() + "a3.in",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "flat_recursive" + ioutils.GetFileSeperator() + "b1.in",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "flat_recursive" + ioutils.GetFileSeperator() + "b2.in",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "flat_recursive" + ioutils.GetFileSeperator() + "b3.in",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "flat_recursive" + ioutils.GetFileSeperator() + "c1.in",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "flat_recursive" + ioutils.GetFileSeperator() + "c2.in",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "flat_recursive" + ioutils.GetFileSeperator() + "c3.in",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "nonflat_nonrecursive" + ioutils.GetFileSeperator() + "downloadTestResources" + ioutils.GetFileSeperator() + "a" + ioutils.GetFileSeperator() + "a1.in",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "nonflat_nonrecursive" + ioutils.GetFileSeperator() + "downloadTestResources" + ioutils.GetFileSeperator() + "a" + ioutils.GetFileSeperator() + "a2.in",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "nonflat_nonrecursive" + ioutils.GetFileSeperator() + "downloadTestResources" + ioutils.GetFileSeperator() + "a" + ioutils.GetFileSeperator() + "a3.in",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "nonflat_recursive" + ioutils.GetFileSeperator() + "downloadTestResources" + ioutils.GetFileSeperator() + "a",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "nonflat_recursive" + ioutils.GetFileSeperator() + "downloadTestResources" + ioutils.GetFileSeperator() + "a" + ioutils.GetFileSeperator() + "a1.in",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "nonflat_recursive" + ioutils.GetFileSeperator() + "downloadTestResources" + ioutils.GetFileSeperator() + "a" + ioutils.GetFileSeperator() + "a2.in",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "nonflat_recursive" + ioutils.GetFileSeperator() + "downloadTestResources" + ioutils.GetFileSeperator() + "a" + ioutils.GetFileSeperator() + "a3.in",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "nonflat_recursive" + ioutils.GetFileSeperator() + "downloadTestResources" + ioutils.GetFileSeperator() + "a" + ioutils.GetFileSeperator() + "b",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "nonflat_recursive" + ioutils.GetFileSeperator() + "downloadTestResources" + ioutils.GetFileSeperator() + "a" + ioutils.GetFileSeperator() + "b" + ioutils.GetFileSeperator() + "b1.in",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "nonflat_recursive" + ioutils.GetFileSeperator() + "downloadTestResources" + ioutils.GetFileSeperator() + "a" + ioutils.GetFileSeperator() + "b" + ioutils.GetFileSeperator() + "b2.in",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "nonflat_recursive" + ioutils.GetFileSeperator() + "downloadTestResources" + ioutils.GetFileSeperator() + "a" + ioutils.GetFileSeperator() + "b" + ioutils.GetFileSeperator() + "b3.in",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "nonflat_recursive" + ioutils.GetFileSeperator() + "downloadTestResources" + ioutils.GetFileSeperator() + "a" + ioutils.GetFileSeperator() + "b" + ioutils.GetFileSeperator() + "c" + ioutils.GetFileSeperator() + "c1.in",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "nonflat_recursive" + ioutils.GetFileSeperator() + "downloadTestResources" + ioutils.GetFileSeperator() + "a" + ioutils.GetFileSeperator() + "b" + ioutils.GetFileSeperator() + "c" + ioutils.GetFileSeperator() + "c2.in",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "nonflat_recursive" + ioutils.GetFileSeperator() + "downloadTestResources" + ioutils.GetFileSeperator() + "a" + ioutils.GetFileSeperator() + "b" + ioutils.GetFileSeperator() + "c" + ioutils.GetFileSeperator() + "c3.in",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "properties" + ioutils.GetFileSeperator() + "downloadTestResources" + ioutils.GetFileSeperator() + "a" + ioutils.GetFileSeperator() + "a1.in",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "properties" + ioutils.GetFileSeperator() + "downloadTestResources" + ioutils.GetFileSeperator() + "a" + ioutils.GetFileSeperator() + "a2.in",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "properties" + ioutils.GetFileSeperator() + "downloadTestResources" + ioutils.GetFileSeperator() + "a" + ioutils.GetFileSeperator() + "a3.in",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "properties" + ioutils.GetFileSeperator() + "downloadTestResources" + ioutils.GetFileSeperator() + "a" + ioutils.GetFileSeperator() + "b" + ioutils.GetFileSeperator() + "b1.in",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "properties" + ioutils.GetFileSeperator() + "downloadTestResources" + ioutils.GetFileSeperator() + "a" + ioutils.GetFileSeperator() + "b" + ioutils.GetFileSeperator() + "b2.in",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "properties" + ioutils.GetFileSeperator() + "downloadTestResources" + ioutils.GetFileSeperator() + "a" + ioutils.GetFileSeperator() + "b" + ioutils.GetFileSeperator() + "b3.in",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "properties" + ioutils.GetFileSeperator() + "downloadTestResources" + ioutils.GetFileSeperator() + "a" + ioutils.GetFileSeperator() + "b" + ioutils.GetFileSeperator() + "c" + ioutils.GetFileSeperator() + "c1.in",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "properties" + ioutils.GetFileSeperator() + "downloadTestResources" + ioutils.GetFileSeperator() + "a" + ioutils.GetFileSeperator() + "b" + ioutils.GetFileSeperator() + "c" + ioutils.GetFileSeperator() + "c2.in",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "properties" + ioutils.GetFileSeperator() + "downloadTestResources" + ioutils.GetFileSeperator() + "a" + ioutils.GetFileSeperator() + "b" + ioutils.GetFileSeperator() + "c" + ioutils.GetFileSeperator() + "c3.in",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "rename" + ioutils.GetFileSeperator() + "a1.out",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "simple" + ioutils.GetFileSeperator() + "a1.in",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "simple_placeholder" + ioutils.GetFileSeperator() + "a",
	Out + ioutils.GetFileSeperator() + "download" + ioutils.GetFileSeperator() + "simple_placeholder" + ioutils.GetFileSeperator() + "a" + ioutils.GetFileSeperator() + "a1.in",
}

var MassiveUpload = []string{
	Repo1 + "/spec-copy-test/3_only_flat_recursive/a3.in",
	Repo1 + "/spec-copy-test/3_only_flat_recursive/b3.in",
	Repo1 + "/spec-copy-test/3_only_flat_recursive/c3.in",
	Repo1 + "/spec-copy-test/copy-to-existing/a1.in",
	Repo1 + "/spec-copy-test/copy-to-existing/a2.in",
	Repo1 + "/spec-copy-test/copy-to-existing/a3.in",
	Repo1 + "/spec-copy-test/copy-to-existing/b1.in",
	Repo1 + "/spec-copy-test/copy-to-existing/b2.in",
	Repo1 + "/spec-copy-test/copy-to-existing/b3.in",
	Repo1 + "/spec-copy-test/copy-to-existing/c1.in",
	Repo1 + "/spec-copy-test/copy-to-existing/c2.in",
	Repo1 + "/spec-copy-test/copy-to-existing/c3.in",
	Repo1 + "/spec-copy-test/defaults_recursive_nonFlat/a1.in",
	Repo1 + "/spec-copy-test/defaults_recursive_nonFlat/a2.in",
	Repo1 + "/spec-copy-test/defaults_recursive_nonFlat/a3.in",
	Repo1 + "/spec-copy-test/defaults_recursive_nonFlat/b1.in",
	Repo1 + "/spec-copy-test/defaults_recursive_nonFlat/b2.in",
	Repo1 + "/spec-copy-test/defaults_recursive_nonFlat/b3.in",
	Repo1 + "/spec-copy-test/defaults_recursive_nonFlat/c1.in",
	Repo1 + "/spec-copy-test/defaults_recursive_nonFlat/c2.in",
	Repo1 + "/spec-copy-test/defaults_recursive_nonFlat/c3.in",
	Repo1 + "/spec-copy-test/flat_nonrecursive/a1.in",
	Repo1 + "/spec-copy-test/flat_nonrecursive/a2.in",
	Repo1 + "/spec-copy-test/flat_nonrecursive/a3.in",
	Repo1 + "/spec-copy-test/flat_recursive/a1.in",
	Repo1 + "/spec-copy-test/flat_recursive/a2.in",
	Repo1 + "/spec-copy-test/flat_recursive/a3.in",
	Repo1 + "/spec-copy-test/flat_recursive/b1.in",
	Repo1 + "/spec-copy-test/flat_recursive/b2.in",
	Repo1 + "/spec-copy-test/flat_recursive/b3.in",
	Repo1 + "/spec-copy-test/flat_recursive/c1.in",
	Repo1 + "/spec-copy-test/flat_recursive/c2.in",
	Repo1 + "/spec-copy-test/flat_recursive/c3.in",
	Repo1 + "/spec-copy-test/nonflat_nonrecursive/testsdata/a/a1.in",
	Repo1 + "/spec-copy-test/nonflat_nonrecursive/testsdata/a/a2.in",
	Repo1 + "/spec-copy-test/nonflat_nonrecursive/testsdata/a/a3.in",
	Repo1 + "/spec-copy-test/nonflat_recursive/testsdata/a/a1.in",
	Repo1 + "/spec-copy-test/nonflat_recursive/testsdata/a/a2.in",
	Repo1 + "/spec-copy-test/nonflat_recursive/testsdata/a/a3.in",
	Repo1 + "/spec-copy-test/nonflat_recursive/testsdata/a/b/b1.in",
	Repo1 + "/spec-copy-test/nonflat_recursive/testsdata/a/b/b2.in",
	Repo1 + "/spec-copy-test/nonflat_recursive/testsdata/a/b/b3.in",
	Repo1 + "/spec-copy-test/nonflat_recursive/testsdata/a/b/c/c1.in",
	Repo1 + "/spec-copy-test/nonflat_recursive/testsdata/a/b/c/c2.in",
	Repo1 + "/spec-copy-test/nonflat_recursive/testsdata/a/b/c/c3.in",
	Repo1 + "/spec-copy-test/properties/testsdata/a/a1.in",
	Repo1 + "/spec-copy-test/properties/testsdata/a/a2.in",
	Repo1 + "/spec-copy-test/properties/testsdata/a/a3.in",
	Repo1 + "/spec-copy-test/properties/testsdata/a/b/b1.in",
	Repo1 + "/spec-copy-test/properties/testsdata/a/b/b2.in",
	Repo1 + "/spec-copy-test/properties/testsdata/a/b/b3.in",
	Repo1 + "/spec-copy-test/properties/testsdata/a/b/c/c1.in",
	Repo1 + "/spec-copy-test/properties/testsdata/a/b/c/c2.in",
	Repo1 + "/spec-copy-test/properties/testsdata/a/b/c/c3.in",
	Repo1 + "/spec-copy-test/simple/a1.in",
}

var PropsExpected = []string{
	Repo1 + "/spec-copy-test/properties/testsdata/a/a1.in",
	Repo1 + "/spec-copy-test/properties/testsdata/a/b/b1.in",
	Repo1 + "/spec-copy-test/properties/testsdata/a/a2.in",
	Repo1 + "/spec-copy-test/properties/testsdata/a/b/b2.in",
	Repo1 + "/spec-copy-test/properties/testsdata/a/b/c/c1.in",
	Repo1 + "/spec-copy-test/properties/testsdata/a/a3.in",
	Repo1 + "/spec-copy-test/properties/testsdata/a/b/b3.in",
	Repo1 + "/spec-copy-test/properties/testsdata/a/b/c/c2.in",
	Repo1 + "/spec-copy-test/properties/testsdata/a/b/c/c3.in",
}

var Delete1 = []string{
	Repo2 + "/3_only_flat_recursive_target/a3.in",
	Repo2 + "/3_only_flat_recursive_target/b3.in",
	Repo2 + "/3_only_flat_recursive_target/c3.in",
	Repo2 + "/defaults_recursive_nonFlat_target/defaults_recursive_nonFlat_source/a/a1.in",
	Repo2 + "/defaults_recursive_nonFlat_target/defaults_recursive_nonFlat_source/a/a2.in",
	Repo2 + "/defaults_recursive_nonFlat_target/defaults_recursive_nonFlat_source/a/a3.in",
	Repo2 + "/defaults_recursive_nonFlat_target/defaults_recursive_nonFlat_source/a/b/b1.in",
	Repo2 + "/defaults_recursive_nonFlat_target/defaults_recursive_nonFlat_source/a/b/b2.in",
	Repo2 + "/defaults_recursive_nonFlat_target/defaults_recursive_nonFlat_source/a/b/b3.in",
	Repo2 + "/defaults_recursive_nonFlat_target/defaults_recursive_nonFlat_source/a/b/c/c1.in",
	Repo2 + "/defaults_recursive_nonFlat_target/defaults_recursive_nonFlat_source/a/b/c/c2.in",
	Repo2 + "/defaults_recursive_nonFlat_target/defaults_recursive_nonFlat_source/a/b/c/c3.in",
	Repo2 + "/flat_nonrecursive_target/a1.in",
	Repo2 + "/flat_nonrecursive_target/a2.in",
	Repo2 + "/flat_nonrecursive_target/a3.in",
	Repo2 + "/flat_recursive_target/a1.in",
	Repo2 + "/flat_recursive_target/a2.in",
	Repo2 + "/flat_recursive_target/a3.in",
	Repo2 + "/flat_recursive_target/b1.in",
	Repo2 + "/flat_recursive_target/b2.in",
	Repo2 + "/flat_recursive_target/b3.in",
	Repo2 + "/flat_recursive_target/c1.in",
	Repo2 + "/flat_recursive_target/c2.in",
	Repo2 + "/flat_recursive_target/c3.in",
	Repo2 + "/no_target/a/a1.in",
	Repo2 + "/no_target/a/a2.in",
	Repo2 + "/no_target/a/a3.in",
	Repo2 + "/no_target/a/b/b1.in",
	Repo2 + "/no_target/a/b/b2.in",
	Repo2 + "/no_target/a/b/b3.in",
	Repo2 + "/no_target/a/b/c/c1.in",
	Repo2 + "/no_target/a/b/c/c2.in",
	Repo2 + "/no_target/a/b/c/c3.in",
	Repo2 + "/nonflat_nonrecursive_target/nonflat_nonrecursive_source/a/a1.in",
	Repo2 + "/nonflat_nonrecursive_target/nonflat_nonrecursive_source/a/a2.in",
	Repo2 + "/nonflat_nonrecursive_target/nonflat_nonrecursive_source/a/a3.in",
	Repo2 + "/nonflat_recursive_target/nonflat_recursive_source/a/a1.in",
	Repo2 + "/nonflat_recursive_target/nonflat_recursive_source/a/a2.in",
	Repo2 + "/nonflat_recursive_target/nonflat_recursive_source/a/a3.in",
	Repo2 + "/pattern_placeholder_target/a/a1.in",
	Repo2 + "/pattern_placeholder_target/a/a2.in",
	Repo2 + "/pattern_placeholder_target/a/a3.in",
	Repo2 + "/pattern_placeholder_target/a/b/b1.in",
	Repo2 + "/pattern_placeholder_target/a/b/b2.in",
	Repo2 + "/pattern_placeholder_target/a/b/b3.in",
	Repo2 + "/pattern_placeholder_target/a/b/c/c1.in",
	Repo2 + "/pattern_placeholder_target/a/b/c/c2.in",
	Repo2 + "/pattern_placeholder_target/a/b/c/c3.in",
	Repo2 + "/properties_target/properties_source/a/a1.in",
	Repo2 + "/properties_target/properties_source/a/a2.in",
	Repo2 + "/properties_target/properties_source/a/a3.in",
	Repo2 + "/properties_target/properties_source/a/b/b1.in",
	Repo2 + "/properties_target/properties_source/a/b/b2.in",
	Repo2 + "/properties_target/properties_source/a/b/b3.in",
	Repo2 + "/properties_target/properties_source/a/b/c/c1.in",
	Repo2 + "/properties_target/properties_source/a/b/c/c2.in",
	Repo2 + "/properties_target/properties_source/a/b/c/c3.in",
	Repo2 + "/rename_target/RENAMED.in",
	Repo2 + "/simple_placeholder_target/a/a1.in",
	Repo2 + "/simple_target/a1.in",
}

var DeleteDisplyedFiles = []string{
	Repo2 + "/3_only_flat_recursive_source/a/a1.in",
	Repo2 + "/3_only_flat_recursive_source/a/a2.in",
	Repo2 + "/3_only_flat_recursive_source/a/a3.in",
	Repo2 + "/flat_recursive_source/a/b/b1.in",
	Repo2 + "/flat_recursive_source/a/b/b2.in",
	Repo2 + "/flat_recursive_source/a/b/b3.in",
	Repo2 + "/flat_recursive_source/a/b/c/c1.in",
	Repo2 + "/flat_recursive_source/a/b/c/c2.in",
	Repo2 + "/flat_recursive_source/a/b/c/c3.in",
	Repo2 + "/defaults_recursive_nonFlat_source/a/a1.in",
	Repo2 + "/defaults_recursive_nonFlat_source/a/a2.in",
	Repo2 + "/defaults_recursive_nonFlat_source/a/a3.in",
	Repo2 + "/defaults_recursive_nonFlat_source/a/b/b1.in",
	Repo2 + "/defaults_recursive_nonFlat_source/a/b/b2.in",
	Repo2 + "/defaults_recursive_nonFlat_source/a/b/b3.in",
	Repo2 + "/defaults_recursive_nonFlat_source/a/b/c/c1.in",
	Repo2 + "/defaults_recursive_nonFlat_source/a/b/c/c2.in",
	Repo2 + "/defaults_recursive_nonFlat_source/a/b/c/c3.in",
	Repo2 + "/flat_nonrecursive_source/a/b/",
}