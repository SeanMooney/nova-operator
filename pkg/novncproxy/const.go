/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package novncproxy

const (
	// MergedServiceConfigPath - The location of the merged configuration file
	// for the service
	MergedServiceConfigPath = "/var/lib/openstack/config/nova-novncproxy-config.json"
	//NoVNCProxyPort - The port the nova-noVNCProxyPort service is exposed on
	NoVNCProxyPort = 6080
	// ServiceName - The name of the service exposed to k8s
	ServiceName = "nova-novncproxy"
	Host        = "::0"
)
