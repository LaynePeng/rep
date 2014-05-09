package models

type StagingRequestFromCC struct {
	AppId                          string                `json:"app_id"`
	TaskId                         string                `json:"task_id"`
	Stack                          string                `json:"stack"`
	AppBitsDownloadUri             string                `json:"app_bits_download_uri"`
	BuildArtifactsCacheDownloadUri string                `json:"build_artifacts_cache_download_uri,omitempty"`
	FileDescriptors                int                   `json:"file_descriptors"`
	MemoryMB                       int                   `json:"memory_mb"`
	DiskMB                         int                   `json:"disk_mb"`
	Buildpacks                     []Buildpack           `json:"buildpacks"`
	Environment                    []EnvironmentVariable `json:"environment"`
}

type Buildpack struct {
	Name string `json:"name"`
	Key  string `json:"key"`
	Url  string `json:"url"`
}

type EnvironmentVariable struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type StagingInfo struct {
	BuildpackKey      string `yaml:"-" json:"buildpack_key,omitempty"`
	DetectedBuildpack string `yaml:"detected_buildpack" json:"detected_buildpack"`
	StartCommand      string `yaml:"start_command" json:"-"`
}

type StagingResponseForCC struct {
	AppId             string `json:"app_id,omitempty"`
	TaskId            string `json:"task_id,omitempty"`
	BuildpackKey      string `json:"buildpack_key,omitempty"`
	DetectedBuildpack string `json:"detected_buildpack,omitempty"`
	Error             string `json:"error,omitempty"`
}

type StagingTaskAnnotation struct {
	AppId  string `json:"app_id"`
	TaskId string `json:"task_id"`
}