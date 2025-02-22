/*
Copyright 2022-2023 Nutanix, Inc.

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

package common

// Constants are defined in lexographical order
const (
	DATABASE_CR_STATUS_DELETING     = "DELETING"
	DATABASE_CR_STATUS_EMPTY        = ""
	DATABASE_CR_STATUS_READY        = "READY"
	DATABASE_CR_STATUS_PROVISIONING = "PROVISIONING"

	DATABASE_DEFAULT_INSTANCE_NAME = "database_instance_name"

	DATABASE_DEFAULT_PORT_MONGODB  = 27017
	DATABASE_DEFAULT_PORT_MSSQL    = 1433
	DATABASE_DEFAULT_PORT_MYSQL    = 3306
	DATABASE_DEFAULT_PORT_POSTGRES = 5432

	DATABASE_ENGINE_TYPE_GENERIC  = "Generic"
	DATABASE_ENGINE_TYPE_MONGODB  = "mongodb_database"
	DATABASE_ENGINE_TYPE_MSSQL    = "sqlserver_database"
	DATABASE_ENGINE_TYPE_MYSQL    = "mysql_database"
	DATABASE_ENGINE_TYPE_ORACLE   = "oracle_database"
	DATABASE_ENGINE_TYPE_POSTGRES = "postgres_database"

	DATABASE_TYPE_GENERIC  = "generic"
	DATABASE_TYPE_MONGODB  = "mongodb"
	DATABASE_TYPE_MSSQL    = "mssql"
	DATABASE_TYPE_MYSQL    = "mysql"
	DATABASE_TYPE_ORACLE   = "oracle"
	DATABASE_TYPE_POSTGRES = "postgres"

	FINALIZER_DATABASE_INSTANCE = "ndb.nutanix.com/finalizerdatabaseinstance"
	FINALIZER_DATABASE_SERVER   = "ndb.nutanix.com/finalizerdatabaseserver"

	NDB_PARAM_USERNAME       = "username"
	NDB_PARAM_PASSWORD       = "password"
	NDB_PARAM_SSH_PUBLIC_KEY = "ssh_public_key"

	PROFILE_DEFAULT_OOB_SMALL_COMPUTE = "DEFAULT_OOB_SMALL_COMPUTE"

	PROFILE_MAP_PARAM = "profileMap"

	PROFILE_STATUS_READY = "READY"

	PROFILE_TYPE_COMPUTE                     = "Compute"
	PROFILE_TYPE_DATABASE_PARAMETER          = "Database_Parameter"
	PROFILE_TYPE_DATABASE_PARAMETER_INSTANCE = "Database_Parameter_Instance"
	PROFILE_TYPE_NETWORK                     = "Network"
	PROFILE_TYPE_SOFTWARE                    = "Software"

	PROPERTY_NAME_VM_IP = "vm_ip"

	SECRET_DATA_KEY_CA_CERTIFICATE = "ca_certificate"
	SECRET_DATA_KEY_PASSWORD       = "password"
	SECRET_DATA_KEY_SSH_PUBLIC_KEY = "ssh_public_key"
	SECRET_DATA_KEY_USERNAME       = "username"

	SLA_NAME_NONE = "NONE"

	TIMEZONE_UTC = "UTC"

	TOPOLOGY_ALL      = "ALL"
	TOPOLOGY_CLUSTER  = "cluster"
	TOPOLOGY_DATABASE = "database"
	TOPOLOGY_INSTANCE = "instance"
	TOPOLOGY_SINGLE   = "single"
)
