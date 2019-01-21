package generated

import (
	"github.com/hashicorp/terraform/terraform"
)

// Goodnight sweet linter, I'm sorry

type Aws_vpc_endpoint_route_table_association struct {
	Route_table_id  string
	Vpc_endpoint_id string
}

func Aws_vpc_endpoint_route_table_associationMapper(r *Aws_vpc_endpoint_route_table_association) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["route_table_id"] = r.Route_table_id
	config["vpc_endpoint_id"] = r.Vpc_endpoint_id
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_autoscaling_notification struct {
	Topic_arn string
}

func Aws_autoscaling_notificationMapper(r *Aws_autoscaling_notification) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["topic_arn"] = r.Topic_arn
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_iam_group_policy_attachment struct {
	Policy_arn string
	Group      string
}

func Aws_iam_group_policy_attachmentMapper(r *Aws_iam_group_policy_attachment) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["policy_arn"] = r.Policy_arn
	config["group"] = r.Group
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_lightsail_static_ip_attachment struct {
	Static_ip_name string
	Instance_name  string
}

func Aws_lightsail_static_ip_attachmentMapper(r *Aws_lightsail_static_ip_attachment) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["static_ip_name"] = r.Static_ip_name
	config["instance_name"] = r.Instance_name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_opsworks_haproxy_layer struct {
	Install_updates_on_boot     *bool
	Drain_elb_on_shutdown       *bool
	Auto_assign_public_ips      *bool
	Auto_healing                *bool
	Use_ebs_optimized_instances *bool
	Healthcheck_url             *string
	Custom_json                 *string
	Name                        *string
	Stats_user                  *string
	Auto_assign_elastic_ips     *bool
	Elastic_load_balancer       *string
	Stats_password              string
	Healthcheck_method          *string
	Stats_enabled               *bool
	Stats_url                   *string
	Custom_instance_profile_arn *string
	Stack_id                    string
}

func Aws_opsworks_haproxy_layerMapper(r *Aws_opsworks_haproxy_layer) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Install_updates_on_boot != nil {
		config["install_updates_on_boot"] = *r.Install_updates_on_boot
	}
	if r.Drain_elb_on_shutdown != nil {
		config["drain_elb_on_shutdown"] = *r.Drain_elb_on_shutdown
	}
	if r.Auto_assign_public_ips != nil {
		config["auto_assign_public_ips"] = *r.Auto_assign_public_ips
	}
	if r.Auto_healing != nil {
		config["auto_healing"] = *r.Auto_healing
	}
	if r.Use_ebs_optimized_instances != nil {
		config["use_ebs_optimized_instances"] = *r.Use_ebs_optimized_instances
	}
	if r.Healthcheck_url != nil {
		config["healthcheck_url"] = *r.Healthcheck_url
	}
	if r.Custom_json != nil {
		config["custom_json"] = *r.Custom_json
	}
	if r.Name != nil {
		config["name"] = *r.Name
	}
	if r.Stats_user != nil {
		config["stats_user"] = *r.Stats_user
	}
	if r.Auto_assign_elastic_ips != nil {
		config["auto_assign_elastic_ips"] = *r.Auto_assign_elastic_ips
	}
	if r.Elastic_load_balancer != nil {
		config["elastic_load_balancer"] = *r.Elastic_load_balancer
	}
	config["stats_password"] = r.Stats_password
	if r.Healthcheck_method != nil {
		config["healthcheck_method"] = *r.Healthcheck_method
	}
	if r.Stats_enabled != nil {
		config["stats_enabled"] = *r.Stats_enabled
	}
	if r.Stats_url != nil {
		config["stats_url"] = *r.Stats_url
	}
	if r.Custom_instance_profile_arn != nil {
		config["custom_instance_profile_arn"] = *r.Custom_instance_profile_arn
	}
	config["stack_id"] = r.Stack_id
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_sqs_queue struct {
	Policy                      *string
	Fifo_queue                  *bool
	Content_based_deduplication *bool
	Name                        *string
	Kms_master_key_id           *string
	Name_prefix                 *string
	Redrive_policy              *string
	Arn                         string
	Tags                        *map[string]interface{}
}

func Aws_sqs_queueMapper(r *Aws_sqs_queue) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Policy != nil {
		config["policy"] = *r.Policy
	}
	if r.Fifo_queue != nil {
		config["fifo_queue"] = *r.Fifo_queue
	}
	if r.Content_based_deduplication != nil {
		config["content_based_deduplication"] = *r.Content_based_deduplication
	}
	if r.Name != nil {
		config["name"] = *r.Name
	}
	if r.Kms_master_key_id != nil {
		config["kms_master_key_id"] = *r.Kms_master_key_id
	}
	if r.Name_prefix != nil {
		config["name_prefix"] = *r.Name_prefix
	}
	if r.Redrive_policy != nil {
		config["redrive_policy"] = *r.Redrive_policy
	}
	config["arn"] = r.Arn
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_cloudformation_stack struct {
	Template_body    *string
	Name             string
	Outputs          map[string]interface{}
	Policy_url       *string
	Iam_role_arn     *string
	Disable_rollback *bool
	On_failure       *string
	Parameters       *map[string]interface{}
	Policy_body      *string
	Template_url     *string
	Tags             *map[string]interface{}
}

func Aws_cloudformation_stackMapper(r *Aws_cloudformation_stack) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Template_url != nil {
		config["template_url"] = *r.Template_url
	}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	if r.Template_body != nil {
		config["template_body"] = *r.Template_body
	}
	config["name"] = r.Name
	config["outputs"] = r.Outputs
	if r.Policy_url != nil {
		config["policy_url"] = *r.Policy_url
	}
	if r.Parameters != nil {
		config["parameters"] = *r.Parameters
	}
	if r.Policy_body != nil {
		config["policy_body"] = *r.Policy_body
	}
	if r.Iam_role_arn != nil {
		config["iam_role_arn"] = *r.Iam_role_arn
	}
	if r.Disable_rollback != nil {
		config["disable_rollback"] = *r.Disable_rollback
	}
	if r.On_failure != nil {
		config["on_failure"] = *r.On_failure
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_rds_cluster struct {
	Database_name                       *string
	Engine_mode                         *string
	Skip_final_snapshot                 *bool
	Master_username                     *string
	Preferred_maintenance_window        *string
	Engine_version                      *string
	Cluster_resource_id                 string
	Reader_endpoint                     string
	Final_snapshot_identifier           *string
	Snapshot_identifier                 *string
	Source_region                       *string
	Arn                                 string
	Deletion_protection                 *bool
	Global_cluster_identifier           *string
	Storage_encrypted                   *bool
	Replication_source_identifier       *string
	Cluster_identifier                  *string
	Db_cluster_parameter_group_name     *string
	Kms_key_id                          *string
	Tags                                *map[string]interface{}
	Endpoint                            string
	Engine                              *string
	Preferred_backup_window             *string
	Cluster_identifier_prefix           *string
	Hosted_zone_id                      string
	Master_password                     *string
	Apply_immediately                   *bool
	Db_subnet_group_name                *string
	Iam_database_authentication_enabled *bool
}

func Aws_rds_clusterMapper(r *Aws_rds_cluster) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Db_subnet_group_name != nil {
		config["db_subnet_group_name"] = *r.Db_subnet_group_name
	}
	if r.Iam_database_authentication_enabled != nil {
		config["iam_database_authentication_enabled"] = *r.Iam_database_authentication_enabled
	}
	if r.Database_name != nil {
		config["database_name"] = *r.Database_name
	}
	if r.Engine_mode != nil {
		config["engine_mode"] = *r.Engine_mode
	}
	if r.Skip_final_snapshot != nil {
		config["skip_final_snapshot"] = *r.Skip_final_snapshot
	}
	if r.Master_username != nil {
		config["master_username"] = *r.Master_username
	}
	if r.Preferred_maintenance_window != nil {
		config["preferred_maintenance_window"] = *r.Preferred_maintenance_window
	}
	if r.Engine_version != nil {
		config["engine_version"] = *r.Engine_version
	}
	config["cluster_resource_id"] = r.Cluster_resource_id
	config["reader_endpoint"] = r.Reader_endpoint
	if r.Final_snapshot_identifier != nil {
		config["final_snapshot_identifier"] = *r.Final_snapshot_identifier
	}
	if r.Snapshot_identifier != nil {
		config["snapshot_identifier"] = *r.Snapshot_identifier
	}
	if r.Source_region != nil {
		config["source_region"] = *r.Source_region
	}
	config["arn"] = r.Arn
	if r.Deletion_protection != nil {
		config["deletion_protection"] = *r.Deletion_protection
	}
	if r.Global_cluster_identifier != nil {
		config["global_cluster_identifier"] = *r.Global_cluster_identifier
	}
	if r.Storage_encrypted != nil {
		config["storage_encrypted"] = *r.Storage_encrypted
	}
	if r.Replication_source_identifier != nil {
		config["replication_source_identifier"] = *r.Replication_source_identifier
	}
	if r.Cluster_identifier != nil {
		config["cluster_identifier"] = *r.Cluster_identifier
	}
	if r.Db_cluster_parameter_group_name != nil {
		config["db_cluster_parameter_group_name"] = *r.Db_cluster_parameter_group_name
	}
	if r.Kms_key_id != nil {
		config["kms_key_id"] = *r.Kms_key_id
	}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["endpoint"] = r.Endpoint
	if r.Engine != nil {
		config["engine"] = *r.Engine
	}
	if r.Preferred_backup_window != nil {
		config["preferred_backup_window"] = *r.Preferred_backup_window
	}
	if r.Cluster_identifier_prefix != nil {
		config["cluster_identifier_prefix"] = *r.Cluster_identifier_prefix
	}
	config["hosted_zone_id"] = r.Hosted_zone_id
	if r.Master_password != nil {
		config["master_password"] = *r.Master_password
	}
	if r.Apply_immediately != nil {
		config["apply_immediately"] = *r.Apply_immediately
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_storagegateway_gateway struct {
	Arn                 string
	Gateway_timezone    string
	Gateway_type        *string
	Medium_changer_type *string
	Tape_drive_type     *string
	Activation_key      *string
	Gateway_id          string
	Gateway_ip_address  *string
	Gateway_name        string
	Smb_guest_password  *string
}

func Aws_storagegateway_gatewayMapper(r *Aws_storagegateway_gateway) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["gateway_name"] = r.Gateway_name
	if r.Smb_guest_password != nil {
		config["smb_guest_password"] = *r.Smb_guest_password
	}
	if r.Activation_key != nil {
		config["activation_key"] = *r.Activation_key
	}
	config["gateway_id"] = r.Gateway_id
	if r.Gateway_ip_address != nil {
		config["gateway_ip_address"] = *r.Gateway_ip_address
	}
	if r.Medium_changer_type != nil {
		config["medium_changer_type"] = *r.Medium_changer_type
	}
	if r.Tape_drive_type != nil {
		config["tape_drive_type"] = *r.Tape_drive_type
	}
	config["arn"] = r.Arn
	config["gateway_timezone"] = r.Gateway_timezone
	if r.Gateway_type != nil {
		config["gateway_type"] = *r.Gateway_type
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_autoscaling_schedule struct {
	End_time               *string
	Recurrence             *string
	Arn                    string
	Scheduled_action_name  string
	Autoscaling_group_name string
	Start_time             *string
}

func Aws_autoscaling_scheduleMapper(r *Aws_autoscaling_schedule) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.End_time != nil {
		config["end_time"] = *r.End_time
	}
	if r.Recurrence != nil {
		config["recurrence"] = *r.Recurrence
	}
	config["arn"] = r.Arn
	config["scheduled_action_name"] = r.Scheduled_action_name
	config["autoscaling_group_name"] = r.Autoscaling_group_name
	if r.Start_time != nil {
		config["start_time"] = *r.Start_time
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_inspector_resource_group struct {
	Tags map[string]interface{}
	Arn  string
}

func Aws_inspector_resource_groupMapper(r *Aws_inspector_resource_group) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["tags"] = r.Tags
	config["arn"] = r.Arn
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_waf_size_constraint_set struct {
	Name string
}

func Aws_waf_size_constraint_setMapper(r *Aws_waf_size_constraint_set) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_wafregional_rule struct {
	Name        string
	Metric_name string
}

func Aws_wafregional_ruleMapper(r *Aws_wafregional_rule) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	config["metric_name"] = r.Metric_name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_dx_connection_association struct {
	Lag_id        string
	Connection_id string
}

func Aws_dx_connection_associationMapper(r *Aws_dx_connection_association) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["connection_id"] = r.Connection_id
	config["lag_id"] = r.Lag_id
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_sns_topic_policy struct {
	Arn    string
	Policy string
}

func Aws_sns_topic_policyMapper(r *Aws_sns_topic_policy) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["arn"] = r.Arn
	config["policy"] = r.Policy
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_ebs_volume struct {
	Snapshot_id       *string
	Resource_type     *string
	Arn               string
	Availability_zone string
	Encrypted         *bool
	Kms_key_id        *string
	Tags              *map[string]interface{}
}

func Aws_ebs_volumeMapper(r *Aws_ebs_volume) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Snapshot_id != nil {
		config["snapshot_id"] = *r.Snapshot_id
	}
	if r.Resource_type != nil {
		config["resource_type"] = *r.Resource_type
	}
	config["arn"] = r.Arn
	config["availability_zone"] = r.Availability_zone
	if r.Encrypted != nil {
		config["encrypted"] = *r.Encrypted
	}
	if r.Kms_key_id != nil {
		config["kms_key_id"] = *r.Kms_key_id
	}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_load_balancer_backend_server_policy struct {
	Load_balancer_name string
}

func Aws_load_balancer_backend_server_policyMapper(r *Aws_load_balancer_backend_server_policy) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["load_balancer_name"] = r.Load_balancer_name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_api_gateway_client_certificate struct {
	Description             *string
	Created_date            string
	Expiration_date         string
	Pem_encoded_certificate string
}

func Aws_api_gateway_client_certificateMapper(r *Aws_api_gateway_client_certificate) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["pem_encoded_certificate"] = r.Pem_encoded_certificate
	if r.Description != nil {
		config["description"] = *r.Description
	}
	config["created_date"] = r.Created_date
	config["expiration_date"] = r.Expiration_date
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_dax_parameter_group struct {
	Name        string
	Description *string
}

func Aws_dax_parameter_groupMapper(r *Aws_dax_parameter_group) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	if r.Description != nil {
		config["description"] = *r.Description
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_neptune_cluster_parameter_group struct {
	Arn         string
	Name        *string
	Name_prefix *string
	Family      string
	Description *string
	Tags        *map[string]interface{}
}

func Aws_neptune_cluster_parameter_groupMapper(r *Aws_neptune_cluster_parameter_group) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["arn"] = r.Arn
	if r.Name != nil {
		config["name"] = *r.Name
	}
	if r.Name_prefix != nil {
		config["name_prefix"] = *r.Name_prefix
	}
	config["family"] = r.Family
	if r.Description != nil {
		config["description"] = *r.Description
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_organizations_policy_attachment struct {
	Policy_id string
	Target_id string
}

func Aws_organizations_policy_attachmentMapper(r *Aws_organizations_policy_attachment) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["policy_id"] = r.Policy_id
	config["target_id"] = r.Target_id
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_cloudhsm_v2_cluster struct {
	Source_backup_identifier *string
	Hsm_type                 string
	Cluster_id               string
	Security_group_id        string
	Tags                     *map[string]interface{}
	Vpc_id                   string
	Cluster_state            string
}

func Aws_cloudhsm_v2_clusterMapper(r *Aws_cloudhsm_v2_cluster) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Source_backup_identifier != nil {
		config["source_backup_identifier"] = *r.Source_backup_identifier
	}
	config["hsm_type"] = r.Hsm_type
	config["cluster_id"] = r.Cluster_id
	config["security_group_id"] = r.Security_group_id
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["vpc_id"] = r.Vpc_id
	config["cluster_state"] = r.Cluster_state
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_appmesh_virtual_router struct {
	Mesh_name         string
	Arn               string
	Created_date      string
	Last_updated_date string
	Name              string
}

func Aws_appmesh_virtual_routerMapper(r *Aws_appmesh_virtual_router) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	config["mesh_name"] = r.Mesh_name
	config["arn"] = r.Arn
	config["created_date"] = r.Created_date
	config["last_updated_date"] = r.Last_updated_date
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_elastic_beanstalk_environment struct {
	Name                   string
	Platform_arn           *string
	Version_label          *string
	Wait_for_ready_timeout *string
	Description            *string
	Cname                  string
	Solution_stack_name    *string
	Poll_interval          *string
	Arn                    string
	Application            string
	Cname_prefix           *string
	Tier                   *string
	Template_name          *string
	Tags                   *map[string]interface{}
}

func Aws_elastic_beanstalk_environmentMapper(r *Aws_elastic_beanstalk_environment) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Version_label != nil {
		config["version_label"] = *r.Version_label
	}
	if r.Wait_for_ready_timeout != nil {
		config["wait_for_ready_timeout"] = *r.Wait_for_ready_timeout
	}
	if r.Description != nil {
		config["description"] = *r.Description
	}
	config["cname"] = r.Cname
	if r.Solution_stack_name != nil {
		config["solution_stack_name"] = *r.Solution_stack_name
	}
	if r.Poll_interval != nil {
		config["poll_interval"] = *r.Poll_interval
	}
	if r.Template_name != nil {
		config["template_name"] = *r.Template_name
	}
	config["arn"] = r.Arn
	config["application"] = r.Application
	if r.Cname_prefix != nil {
		config["cname_prefix"] = *r.Cname_prefix
	}
	if r.Tier != nil {
		config["tier"] = *r.Tier
	}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["name"] = r.Name
	if r.Platform_arn != nil {
		config["platform_arn"] = *r.Platform_arn
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_glue_connection struct {
	Name                  string
	Catalog_id            *string
	Connection_properties map[string]interface{}
	Connection_type       *string
	Description           *string
}

func Aws_glue_connectionMapper(r *Aws_glue_connection) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Catalog_id != nil {
		config["catalog_id"] = *r.Catalog_id
	}
	config["connection_properties"] = r.Connection_properties
	if r.Connection_type != nil {
		config["connection_type"] = *r.Connection_type
	}
	if r.Description != nil {
		config["description"] = *r.Description
	}
	config["name"] = r.Name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_s3_bucket_notification struct {
	Bucket string
}

func Aws_s3_bucket_notificationMapper(r *Aws_s3_bucket_notification) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["bucket"] = r.Bucket
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_security_group_rule struct {
	Resource_type            string
	Protocol                 string
	Security_group_id        string
	Self                     *bool
	Source_security_group_id *string
	Description              *string
}

func Aws_security_group_ruleMapper(r *Aws_security_group_rule) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Source_security_group_id != nil {
		config["source_security_group_id"] = *r.Source_security_group_id
	}
	if r.Description != nil {
		config["description"] = *r.Description
	}
	config["resource_type"] = r.Resource_type
	config["protocol"] = r.Protocol
	config["security_group_id"] = r.Security_group_id
	if r.Self != nil {
		config["self"] = *r.Self
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_load_balancer_listener_policy struct {
	Load_balancer_name string
}

func Aws_load_balancer_listener_policyMapper(r *Aws_load_balancer_listener_policy) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["load_balancer_name"] = r.Load_balancer_name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_service_discovery_service struct {
	Arn         string
	Name        string
	Description *string
}

func Aws_service_discovery_serviceMapper(r *Aws_service_discovery_service) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["arn"] = r.Arn
	config["name"] = r.Name
	if r.Description != nil {
		config["description"] = *r.Description
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_guardduty_threatintelset struct {
	Detector_id string
	Name        string
	Format      string
	Location    string
	Activate    bool
}

func Aws_guardduty_threatintelsetMapper(r *Aws_guardduty_threatintelset) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["format"] = r.Format
	config["location"] = r.Location
	config["activate"] = r.Activate
	config["detector_id"] = r.Detector_id
	config["name"] = r.Name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_iam_openid_connect_provider struct {
	Arn string
	Url string
}

func Aws_iam_openid_connect_providerMapper(r *Aws_iam_openid_connect_provider) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["arn"] = r.Arn
	config["url"] = r.Url
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_iam_server_certificate struct {
	Arn               *string
	Certificate_body  string
	Certificate_chain *string
	Path              *string
	Private_key       string
	Name              *string
	Name_prefix       *string
}

func Aws_iam_server_certificateMapper(r *Aws_iam_server_certificate) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Arn != nil {
		config["arn"] = *r.Arn
	}
	config["certificate_body"] = r.Certificate_body
	if r.Certificate_chain != nil {
		config["certificate_chain"] = *r.Certificate_chain
	}
	if r.Path != nil {
		config["path"] = *r.Path
	}
	config["private_key"] = r.Private_key
	if r.Name != nil {
		config["name"] = *r.Name
	}
	if r.Name_prefix != nil {
		config["name_prefix"] = *r.Name_prefix
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_cloudfront_origin_access_identity struct {
	S3_canonical_user_id            string
	Comment                         *string
	Caller_reference                string
	Cloudfront_access_identity_path string
	Etag                            string
	Iam_arn                         string
}

func Aws_cloudfront_origin_access_identityMapper(r *Aws_cloudfront_origin_access_identity) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["cloudfront_access_identity_path"] = r.Cloudfront_access_identity_path
	config["etag"] = r.Etag
	config["iam_arn"] = r.Iam_arn
	config["s3_canonical_user_id"] = r.S3_canonical_user_id
	if r.Comment != nil {
		config["comment"] = *r.Comment
	}
	config["caller_reference"] = r.Caller_reference
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_cloudwatch_event_rule struct {
	Name_prefix         *string
	Schedule_expression *string
	Event_pattern       *string
	Description         *string
	Role_arn            *string
	Is_enabled          *bool
	Arn                 string
	Name                *string
}

func Aws_cloudwatch_event_ruleMapper(r *Aws_cloudwatch_event_rule) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Event_pattern != nil {
		config["event_pattern"] = *r.Event_pattern
	}
	if r.Description != nil {
		config["description"] = *r.Description
	}
	if r.Role_arn != nil {
		config["role_arn"] = *r.Role_arn
	}
	if r.Is_enabled != nil {
		config["is_enabled"] = *r.Is_enabled
	}
	config["arn"] = r.Arn
	if r.Name != nil {
		config["name"] = *r.Name
	}
	if r.Name_prefix != nil {
		config["name_prefix"] = *r.Name_prefix
	}
	if r.Schedule_expression != nil {
		config["schedule_expression"] = *r.Schedule_expression
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_db_option_group struct {
	Engine_name              string
	Major_engine_version     string
	Option_group_description *string
	Tags                     *map[string]interface{}
	Arn                      string
	Name                     *string
	Name_prefix              *string
}

func Aws_db_option_groupMapper(r *Aws_db_option_group) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["major_engine_version"] = r.Major_engine_version
	if r.Option_group_description != nil {
		config["option_group_description"] = *r.Option_group_description
	}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["arn"] = r.Arn
	if r.Name != nil {
		config["name"] = *r.Name
	}
	if r.Name_prefix != nil {
		config["name_prefix"] = *r.Name_prefix
	}
	config["engine_name"] = r.Engine_name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_dms_certificate struct {
	Certificate_arn    string
	Certificate_id     string
	Certificate_pem    *string
	Certificate_wallet *string
}

func Aws_dms_certificateMapper(r *Aws_dms_certificate) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["certificate_arn"] = r.Certificate_arn
	config["certificate_id"] = r.Certificate_id
	if r.Certificate_pem != nil {
		config["certificate_pem"] = *r.Certificate_pem
	}
	if r.Certificate_wallet != nil {
		config["certificate_wallet"] = *r.Certificate_wallet
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_elastic_beanstalk_configuration_template struct {
	Environment_id      *string
	Solution_stack_name *string
	Name                string
	Application         string
	Description         *string
}

func Aws_elastic_beanstalk_configuration_templateMapper(r *Aws_elastic_beanstalk_configuration_template) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	config["application"] = r.Application
	if r.Description != nil {
		config["description"] = *r.Description
	}
	if r.Environment_id != nil {
		config["environment_id"] = *r.Environment_id
	}
	if r.Solution_stack_name != nil {
		config["solution_stack_name"] = *r.Solution_stack_name
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_secretsmanager_secret struct {
	Description         *string
	Kms_key_id          *string
	Policy              *string
	Rotation_lambda_arn *string
	Tags                *map[string]interface{}
	Arn                 string
	Name                *string
	Name_prefix         *string
	Rotation_enabled    bool
}

func Aws_secretsmanager_secretMapper(r *Aws_secretsmanager_secret) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	if r.Description != nil {
		config["description"] = *r.Description
	}
	if r.Kms_key_id != nil {
		config["kms_key_id"] = *r.Kms_key_id
	}
	if r.Policy != nil {
		config["policy"] = *r.Policy
	}
	if r.Rotation_lambda_arn != nil {
		config["rotation_lambda_arn"] = *r.Rotation_lambda_arn
	}
	config["arn"] = r.Arn
	if r.Name != nil {
		config["name"] = *r.Name
	}
	if r.Name_prefix != nil {
		config["name_prefix"] = *r.Name_prefix
	}
	config["rotation_enabled"] = r.Rotation_enabled
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_ssm_parameter struct {
	Key_id          *string
	Description     *string
	Resource_type   string
	Value           string
	Arn             *string
	Overwrite       *bool
	Allowed_pattern *string
	Tags            *map[string]interface{}
	Name            string
}

func Aws_ssm_parameterMapper(r *Aws_ssm_parameter) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Key_id != nil {
		config["key_id"] = *r.Key_id
	}
	if r.Arn != nil {
		config["arn"] = *r.Arn
	}
	if r.Overwrite != nil {
		config["overwrite"] = *r.Overwrite
	}
	if r.Allowed_pattern != nil {
		config["allowed_pattern"] = *r.Allowed_pattern
	}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["name"] = r.Name
	if r.Description != nil {
		config["description"] = *r.Description
	}
	config["resource_type"] = r.Resource_type
	config["value"] = r.Value
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_waf_byte_match_set struct {
	Name string
}

func Aws_waf_byte_match_setMapper(r *Aws_waf_byte_match_set) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_appsync_api_key struct {
	Api_id      string
	Expires     *string
	Key         string
	Description *string
}

func Aws_appsync_api_keyMapper(r *Aws_appsync_api_key) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["key"] = r.Key
	if r.Description != nil {
		config["description"] = *r.Description
	}
	config["api_id"] = r.Api_id
	if r.Expires != nil {
		config["expires"] = *r.Expires
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_datasync_task struct {
	Arn                      string
	Cloudwatch_log_group_arn *string
	Destination_location_arn string
	Name                     *string
	Source_location_arn      string
	Tags                     *map[string]interface{}
}

func Aws_datasync_taskMapper(r *Aws_datasync_task) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["source_location_arn"] = r.Source_location_arn
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["arn"] = r.Arn
	if r.Cloudwatch_log_group_arn != nil {
		config["cloudwatch_log_group_arn"] = *r.Cloudwatch_log_group_arn
	}
	config["destination_location_arn"] = r.Destination_location_arn
	if r.Name != nil {
		config["name"] = *r.Name
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_ecr_repository struct {
	Name           string
	Tags           *map[string]interface{}
	Arn            string
	Registry_id    string
	Repository_url string
}

func Aws_ecr_repositoryMapper(r *Aws_ecr_repository) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["arn"] = r.Arn
	config["registry_id"] = r.Registry_id
	config["repository_url"] = r.Repository_url
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_neptune_cluster_instance struct {
	Storage_encrypted            bool
	Address                      string
	Auto_minor_version_upgrade   *bool
	Cluster_identifier           string
	Instance_class               string
	Neptune_subnet_group_name    *string
	Identifier_prefix            *string
	Kms_key_arn                  string
	Availability_zone            *string
	Dbi_resource_id              string
	Endpoint                     string
	Engine                       *string
	Engine_version               *string
	Preferred_backup_window      *string
	Arn                          string
	Identifier                   *string
	Tags                         *map[string]interface{}
	Writer                       bool
	Apply_immediately            *bool
	Neptune_parameter_group_name *string
	Preferred_maintenance_window *string
	Publicly_accessible          *bool
}

func Aws_neptune_cluster_instanceMapper(r *Aws_neptune_cluster_instance) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["dbi_resource_id"] = r.Dbi_resource_id
	config["endpoint"] = r.Endpoint
	if r.Engine != nil {
		config["engine"] = *r.Engine
	}
	if r.Engine_version != nil {
		config["engine_version"] = *r.Engine_version
	}
	if r.Identifier_prefix != nil {
		config["identifier_prefix"] = *r.Identifier_prefix
	}
	config["kms_key_arn"] = r.Kms_key_arn
	if r.Availability_zone != nil {
		config["availability_zone"] = *r.Availability_zone
	}
	if r.Preferred_backup_window != nil {
		config["preferred_backup_window"] = *r.Preferred_backup_window
	}
	if r.Identifier != nil {
		config["identifier"] = *r.Identifier
	}
	config["arn"] = r.Arn
	if r.Neptune_parameter_group_name != nil {
		config["neptune_parameter_group_name"] = *r.Neptune_parameter_group_name
	}
	if r.Preferred_maintenance_window != nil {
		config["preferred_maintenance_window"] = *r.Preferred_maintenance_window
	}
	if r.Publicly_accessible != nil {
		config["publicly_accessible"] = *r.Publicly_accessible
	}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["writer"] = r.Writer
	if r.Apply_immediately != nil {
		config["apply_immediately"] = *r.Apply_immediately
	}
	if r.Auto_minor_version_upgrade != nil {
		config["auto_minor_version_upgrade"] = *r.Auto_minor_version_upgrade
	}
	config["cluster_identifier"] = r.Cluster_identifier
	config["instance_class"] = r.Instance_class
	if r.Neptune_subnet_group_name != nil {
		config["neptune_subnet_group_name"] = *r.Neptune_subnet_group_name
	}
	config["storage_encrypted"] = r.Storage_encrypted
	config["address"] = r.Address
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_iam_user struct {
	Arn                  string
	Unique_id            string
	Name                 string
	Path                 *string
	Permissions_boundary *string
	Force_destroy        *bool
	Tags                 *map[string]interface{}
}

func Aws_iam_userMapper(r *Aws_iam_user) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["arn"] = r.Arn
	config["unique_id"] = r.Unique_id
	config["name"] = r.Name
	if r.Path != nil {
		config["path"] = *r.Path
	}
	if r.Permissions_boundary != nil {
		config["permissions_boundary"] = *r.Permissions_boundary
	}
	if r.Force_destroy != nil {
		config["force_destroy"] = *r.Force_destroy
	}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_ami_copy struct {
	Encrypted            *bool
	Kms_key_id           *string
	Architecture         string
	Manage_ebs_snapshots bool
	Name                 string
	Source_ami_id        string
	Sriov_net_support    string
	Virtualization_type  string
	Ena_support          bool
	Kernel_id            string
	Ramdisk_id           string
	Root_snapshot_id     string
	Source_ami_region    string
	Description          *string
	Image_location       string
	Root_device_name     string
	Tags                 *map[string]interface{}
}

func Aws_ami_copyMapper(r *Aws_ami_copy) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Encrypted != nil {
		config["encrypted"] = *r.Encrypted
	}
	if r.Kms_key_id != nil {
		config["kms_key_id"] = *r.Kms_key_id
	}
	config["architecture"] = r.Architecture
	config["name"] = r.Name
	config["source_ami_id"] = r.Source_ami_id
	config["sriov_net_support"] = r.Sriov_net_support
	config["virtualization_type"] = r.Virtualization_type
	config["ena_support"] = r.Ena_support
	config["manage_ebs_snapshots"] = r.Manage_ebs_snapshots
	config["ramdisk_id"] = r.Ramdisk_id
	config["root_snapshot_id"] = r.Root_snapshot_id
	config["source_ami_region"] = r.Source_ami_region
	if r.Description != nil {
		config["description"] = *r.Description
	}
	config["kernel_id"] = r.Kernel_id
	config["root_device_name"] = r.Root_device_name
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["image_location"] = r.Image_location
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_appautoscaling_scheduled_action struct {
	Name               string
	Resource_id        string
	Scalable_dimension *string
	Start_time         *string
	End_time           *string
	Arn                string
	Service_namespace  string
	Schedule           *string
}

func Aws_appautoscaling_scheduled_actionMapper(r *Aws_appautoscaling_scheduled_action) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	config["resource_id"] = r.Resource_id
	if r.Scalable_dimension != nil {
		config["scalable_dimension"] = *r.Scalable_dimension
	}
	if r.Start_time != nil {
		config["start_time"] = *r.Start_time
	}
	if r.End_time != nil {
		config["end_time"] = *r.End_time
	}
	config["arn"] = r.Arn
	config["service_namespace"] = r.Service_namespace
	if r.Schedule != nil {
		config["schedule"] = *r.Schedule
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_cognito_identity_provider struct {
	User_pool_id      string
	Attribute_mapping *map[string]interface{}
	Provider_details  map[string]interface{}
	Provider_name     string
	Provider_type     string
}

func Aws_cognito_identity_providerMapper(r *Aws_cognito_identity_provider) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["provider_details"] = r.Provider_details
	config["provider_name"] = r.Provider_name
	config["provider_type"] = r.Provider_type
	config["user_pool_id"] = r.User_pool_id
	if r.Attribute_mapping != nil {
		config["attribute_mapping"] = *r.Attribute_mapping
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_elastictranscoder_pipeline struct {
	Aws_kms_key_arn *string
	Name            *string
	Role            string
	Arn             string
	Input_bucket    string
	Output_bucket   *string
}

func Aws_elastictranscoder_pipelineMapper(r *Aws_elastictranscoder_pipeline) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Output_bucket != nil {
		config["output_bucket"] = *r.Output_bucket
	}
	config["arn"] = r.Arn
	config["input_bucket"] = r.Input_bucket
	config["role"] = r.Role
	if r.Aws_kms_key_arn != nil {
		config["aws_kms_key_arn"] = *r.Aws_kms_key_arn
	}
	if r.Name != nil {
		config["name"] = *r.Name
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_glue_job struct {
	Name                   string
	Role_arn               string
	Default_arguments      *map[string]interface{}
	Description            *string
	Security_configuration *string
}

func Aws_glue_jobMapper(r *Aws_glue_job) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Default_arguments != nil {
		config["default_arguments"] = *r.Default_arguments
	}
	config["name"] = r.Name
	config["role_arn"] = r.Role_arn
	if r.Description != nil {
		config["description"] = *r.Description
	}
	if r.Security_configuration != nil {
		config["security_configuration"] = *r.Security_configuration
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_ec2_capacity_reservation struct {
	Ebs_optimized           *bool
	End_date                *string
	Ephemeral_storage       *bool
	Instance_platform       string
	Availability_zone       string
	End_date_type           *string
	Instance_match_criteria *string
	Instance_type           string
	Tags                    *map[string]interface{}
	Tenancy                 *string
}

func Aws_ec2_capacity_reservationMapper(r *Aws_ec2_capacity_reservation) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Ephemeral_storage != nil {
		config["ephemeral_storage"] = *r.Ephemeral_storage
	}
	config["instance_platform"] = r.Instance_platform
	if r.Ebs_optimized != nil {
		config["ebs_optimized"] = *r.Ebs_optimized
	}
	if r.End_date != nil {
		config["end_date"] = *r.End_date
	}
	if r.Instance_match_criteria != nil {
		config["instance_match_criteria"] = *r.Instance_match_criteria
	}
	config["instance_type"] = r.Instance_type
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	if r.Tenancy != nil {
		config["tenancy"] = *r.Tenancy
	}
	config["availability_zone"] = r.Availability_zone
	if r.End_date_type != nil {
		config["end_date_type"] = *r.End_date_type
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_lambda_permission struct {
	Statement_id        *string
	Action              string
	Event_source_token  *string
	Function_name       string
	Principal           string
	Qualifier           *string
	Source_account      *string
	Source_arn          *string
	Statement_id_prefix *string
}

func Aws_lambda_permissionMapper(r *Aws_lambda_permission) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Statement_id != nil {
		config["statement_id"] = *r.Statement_id
	}
	if r.Source_account != nil {
		config["source_account"] = *r.Source_account
	}
	if r.Source_arn != nil {
		config["source_arn"] = *r.Source_arn
	}
	if r.Statement_id_prefix != nil {
		config["statement_id_prefix"] = *r.Statement_id_prefix
	}
	config["action"] = r.Action
	if r.Event_source_token != nil {
		config["event_source_token"] = *r.Event_source_token
	}
	config["function_name"] = r.Function_name
	config["principal"] = r.Principal
	if r.Qualifier != nil {
		config["qualifier"] = *r.Qualifier
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_sns_sms_preferences struct {
	Delivery_status_iam_role_arn          *string
	Delivery_status_success_sampling_rate *string
	Default_sender_id                     *string
	Default_sms_type                      *string
	Usage_report_s3_bucket                *string
	Monthly_spend_limit                   *string
}

func Aws_sns_sms_preferencesMapper(r *Aws_sns_sms_preferences) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Monthly_spend_limit != nil {
		config["monthly_spend_limit"] = *r.Monthly_spend_limit
	}
	if r.Delivery_status_iam_role_arn != nil {
		config["delivery_status_iam_role_arn"] = *r.Delivery_status_iam_role_arn
	}
	if r.Delivery_status_success_sampling_rate != nil {
		config["delivery_status_success_sampling_rate"] = *r.Delivery_status_success_sampling_rate
	}
	if r.Default_sender_id != nil {
		config["default_sender_id"] = *r.Default_sender_id
	}
	if r.Default_sms_type != nil {
		config["default_sms_type"] = *r.Default_sms_type
	}
	if r.Usage_report_s3_bucket != nil {
		config["usage_report_s3_bucket"] = *r.Usage_report_s3_bucket
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_vpn_gateway_route_propagation struct {
	Vpn_gateway_id string
	Route_table_id string
}

func Aws_vpn_gateway_route_propagationMapper(r *Aws_vpn_gateway_route_propagation) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["vpn_gateway_id"] = r.Vpn_gateway_id
	config["route_table_id"] = r.Route_table_id
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_datasync_agent struct {
	Ip_address     *string
	Name           *string
	Tags           *map[string]interface{}
	Arn            string
	Activation_key *string
}

func Aws_datasync_agentMapper(r *Aws_datasync_agent) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["arn"] = r.Arn
	if r.Activation_key != nil {
		config["activation_key"] = *r.Activation_key
	}
	if r.Ip_address != nil {
		config["ip_address"] = *r.Ip_address
	}
	if r.Name != nil {
		config["name"] = *r.Name
	}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_guardduty_detector struct {
	Enable                       *bool
	Account_id                   string
	Finding_publishing_frequency *string
}

func Aws_guardduty_detectorMapper(r *Aws_guardduty_detector) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Enable != nil {
		config["enable"] = *r.Enable
	}
	config["account_id"] = r.Account_id
	if r.Finding_publishing_frequency != nil {
		config["finding_publishing_frequency"] = *r.Finding_publishing_frequency
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_ses_receipt_rule struct {
	Rule_set_name string
	Scan_enabled  *bool
	After         *string
	Name          string
	Enabled       *bool
	Tls_policy    *string
}

func Aws_ses_receipt_ruleMapper(r *Aws_ses_receipt_rule) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	if r.Enabled != nil {
		config["enabled"] = *r.Enabled
	}
	if r.Tls_policy != nil {
		config["tls_policy"] = *r.Tls_policy
	}
	config["rule_set_name"] = r.Rule_set_name
	if r.Scan_enabled != nil {
		config["scan_enabled"] = *r.Scan_enabled
	}
	if r.After != nil {
		config["after"] = *r.After
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_s3_bucket_policy struct {
	Bucket string
	Policy string
}

func Aws_s3_bucket_policyMapper(r *Aws_s3_bucket_policy) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["bucket"] = r.Bucket
	config["policy"] = r.Policy
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_app_cookie_stickiness_policy struct {
	Cookie_name   string
	Name          string
	Load_balancer string
}

func Aws_app_cookie_stickiness_policyMapper(r *Aws_app_cookie_stickiness_policy) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	config["load_balancer"] = r.Load_balancer
	config["cookie_name"] = r.Cookie_name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_ec2_transit_gateway struct {
	Auto_accept_shared_attachments     *string
	Default_route_table_propagation    *string
	Description                        *string
	Association_default_route_table_id string
	Dns_support                        *string
	Owner_id                           string
	Propagation_default_route_table_id string
	Tags                               *map[string]interface{}
	Vpn_ecmp_support                   *string
	Arn                                string
	Default_route_table_association    *string
}

func Aws_ec2_transit_gatewayMapper(r *Aws_ec2_transit_gateway) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Default_route_table_propagation != nil {
		config["default_route_table_propagation"] = *r.Default_route_table_propagation
	}
	if r.Description != nil {
		config["description"] = *r.Description
	}
	config["association_default_route_table_id"] = r.Association_default_route_table_id
	if r.Auto_accept_shared_attachments != nil {
		config["auto_accept_shared_attachments"] = *r.Auto_accept_shared_attachments
	}
	config["owner_id"] = r.Owner_id
	config["propagation_default_route_table_id"] = r.Propagation_default_route_table_id
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	if r.Vpn_ecmp_support != nil {
		config["vpn_ecmp_support"] = *r.Vpn_ecmp_support
	}
	config["arn"] = r.Arn
	if r.Default_route_table_association != nil {
		config["default_route_table_association"] = *r.Default_route_table_association
	}
	if r.Dns_support != nil {
		config["dns_support"] = *r.Dns_support
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_ram_resource_share struct {
	Name                      string
	Allow_external_principals *bool
	Tags                      *map[string]interface{}
}

func Aws_ram_resource_shareMapper(r *Aws_ram_resource_share) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Allow_external_principals != nil {
		config["allow_external_principals"] = *r.Allow_external_principals
	}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["name"] = r.Name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_wafregional_web_acl_association struct {
	Web_acl_id   string
	Resource_arn string
}

func Aws_wafregional_web_acl_associationMapper(r *Aws_wafregional_web_acl_association) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["web_acl_id"] = r.Web_acl_id
	config["resource_arn"] = r.Resource_arn
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_cloudhsm_v2_hsm struct {
	Hsm_state         string
	Hsm_eni_id        string
	Cluster_id        string
	Subnet_id         *string
	Availability_zone *string
	Ip_address        *string
	Hsm_id            string
}

func Aws_cloudhsm_v2_hsmMapper(r *Aws_cloudhsm_v2_hsm) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Ip_address != nil {
		config["ip_address"] = *r.Ip_address
	}
	config["hsm_id"] = r.Hsm_id
	config["hsm_state"] = r.Hsm_state
	config["hsm_eni_id"] = r.Hsm_eni_id
	config["cluster_id"] = r.Cluster_id
	if r.Subnet_id != nil {
		config["subnet_id"] = *r.Subnet_id
	}
	if r.Availability_zone != nil {
		config["availability_zone"] = *r.Availability_zone
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_launch_configuration struct {
	Image_id                    string
	Iam_instance_profile        *string
	Spot_price                  *string
	Ebs_optimized               *bool
	Name_prefix                 *string
	Instance_type               string
	Key_name                    *string
	Name                        *string
	User_data                   *string
	Placement_tenancy           *string
	Enable_monitoring           *bool
	User_data_base64            *string
	Vpc_classic_link_id         *string
	Associate_public_ip_address *bool
}

func Aws_launch_configurationMapper(r *Aws_launch_configuration) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["image_id"] = r.Image_id
	if r.Iam_instance_profile != nil {
		config["iam_instance_profile"] = *r.Iam_instance_profile
	}
	if r.Spot_price != nil {
		config["spot_price"] = *r.Spot_price
	}
	if r.Ebs_optimized != nil {
		config["ebs_optimized"] = *r.Ebs_optimized
	}
	if r.Name_prefix != nil {
		config["name_prefix"] = *r.Name_prefix
	}
	config["instance_type"] = r.Instance_type
	if r.Key_name != nil {
		config["key_name"] = *r.Key_name
	}
	if r.Name != nil {
		config["name"] = *r.Name
	}
	if r.User_data != nil {
		config["user_data"] = *r.User_data
	}
	if r.Placement_tenancy != nil {
		config["placement_tenancy"] = *r.Placement_tenancy
	}
	if r.Enable_monitoring != nil {
		config["enable_monitoring"] = *r.Enable_monitoring
	}
	if r.User_data_base64 != nil {
		config["user_data_base64"] = *r.User_data_base64
	}
	if r.Vpc_classic_link_id != nil {
		config["vpc_classic_link_id"] = *r.Vpc_classic_link_id
	}
	if r.Associate_public_ip_address != nil {
		config["associate_public_ip_address"] = *r.Associate_public_ip_address
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_redshift_parameter_group struct {
	Name        string
	Family      string
	Description *string
}

func Aws_redshift_parameter_groupMapper(r *Aws_redshift_parameter_group) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	config["family"] = r.Family
	if r.Description != nil {
		config["description"] = *r.Description
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_s3_bucket_metric struct {
	Bucket string
	Name   string
}

func Aws_s3_bucket_metricMapper(r *Aws_s3_bucket_metric) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	config["bucket"] = r.Bucket
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_securityhub_product_subscription struct {
	Product_arn string
	Arn         string
}

func Aws_securityhub_product_subscriptionMapper(r *Aws_securityhub_product_subscription) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["product_arn"] = r.Product_arn
	config["arn"] = r.Arn
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_api_gateway_gateway_response struct {
	Response_parameters *map[string]interface{}
	Rest_api_id         string
	Response_type       string
	Status_code         *string
	Response_templates  *map[string]interface{}
}

func Aws_api_gateway_gateway_responseMapper(r *Aws_api_gateway_gateway_response) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Status_code != nil {
		config["status_code"] = *r.Status_code
	}
	if r.Response_templates != nil {
		config["response_templates"] = *r.Response_templates
	}
	if r.Response_parameters != nil {
		config["response_parameters"] = *r.Response_parameters
	}
	config["rest_api_id"] = r.Rest_api_id
	config["response_type"] = r.Response_type
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_dms_replication_task struct {
	Replication_task_arn      string
	Table_mappings            string
	Target_endpoint_arn       string
	Migration_type            string
	Replication_instance_arn  string
	Replication_task_settings *string
	Source_endpoint_arn       string
	Tags                      *map[string]interface{}
	Cdc_start_time            *string
	Replication_task_id       string
}

func Aws_dms_replication_taskMapper(r *Aws_dms_replication_task) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Cdc_start_time != nil {
		config["cdc_start_time"] = *r.Cdc_start_time
	}
	config["replication_task_id"] = r.Replication_task_id
	if r.Replication_task_settings != nil {
		config["replication_task_settings"] = *r.Replication_task_settings
	}
	config["source_endpoint_arn"] = r.Source_endpoint_arn
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["migration_type"] = r.Migration_type
	config["replication_instance_arn"] = r.Replication_instance_arn
	config["replication_task_arn"] = r.Replication_task_arn
	config["table_mappings"] = r.Table_mappings
	config["target_endpoint_arn"] = r.Target_endpoint_arn
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_glue_trigger struct {
	Enabled       *bool
	Name          string
	Schedule      *string
	Resource_type string
	Description   *string
}

func Aws_glue_triggerMapper(r *Aws_glue_trigger) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	if r.Schedule != nil {
		config["schedule"] = *r.Schedule
	}
	config["resource_type"] = r.Resource_type
	if r.Description != nil {
		config["description"] = *r.Description
	}
	if r.Enabled != nil {
		config["enabled"] = *r.Enabled
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_batch_compute_environment struct {
	Arn                      string
	Ecc_cluster_arn          string
	Ecs_cluster_arn          string
	Status                   string
	Status_reason            string
	State                    *string
	Resource_type            string
	Compute_environment_name string
	Service_role             string
}

func Aws_batch_compute_environmentMapper(r *Aws_batch_compute_environment) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["compute_environment_name"] = r.Compute_environment_name
	config["service_role"] = r.Service_role
	config["resource_type"] = r.Resource_type
	config["arn"] = r.Arn
	config["ecc_cluster_arn"] = r.Ecc_cluster_arn
	config["ecs_cluster_arn"] = r.Ecs_cluster_arn
	config["status"] = r.Status
	config["status_reason"] = r.Status_reason
	if r.State != nil {
		config["state"] = *r.State
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_s3_bucket struct {
	Bucket                      *string
	Policy                      *string
	Hosted_zone_id              *string
	Website_endpoint            *string
	Acceleration_status         *string
	Arn                         *string
	Bucket_regional_domain_name string
	Acl                         *string
	Website_domain              *string
	Force_destroy               *bool
	Request_payer               *string
	Tags                        *map[string]interface{}
	Bucket_prefix               *string
	Bucket_domain_name          string
	Region                      *string
}

func Aws_s3_bucketMapper(r *Aws_s3_bucket) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Acceleration_status != nil {
		config["acceleration_status"] = *r.Acceleration_status
	}
	if r.Bucket != nil {
		config["bucket"] = *r.Bucket
	}
	if r.Policy != nil {
		config["policy"] = *r.Policy
	}
	if r.Hosted_zone_id != nil {
		config["hosted_zone_id"] = *r.Hosted_zone_id
	}
	if r.Website_endpoint != nil {
		config["website_endpoint"] = *r.Website_endpoint
	}
	if r.Arn != nil {
		config["arn"] = *r.Arn
	}
	if r.Force_destroy != nil {
		config["force_destroy"] = *r.Force_destroy
	}
	if r.Request_payer != nil {
		config["request_payer"] = *r.Request_payer
	}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["bucket_regional_domain_name"] = r.Bucket_regional_domain_name
	if r.Acl != nil {
		config["acl"] = *r.Acl
	}
	if r.Website_domain != nil {
		config["website_domain"] = *r.Website_domain
	}
	if r.Region != nil {
		config["region"] = *r.Region
	}
	if r.Bucket_prefix != nil {
		config["bucket_prefix"] = *r.Bucket_prefix
	}
	config["bucket_domain_name"] = r.Bucket_domain_name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_dx_connection struct {
	Jumbo_frame_capable bool
	Tags                *map[string]interface{}
	Arn                 string
	Name                string
	Bandwidth           string
	Location            string
}

func Aws_dx_connectionMapper(r *Aws_dx_connection) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["location"] = r.Location
	config["jumbo_frame_capable"] = r.Jumbo_frame_capable
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["arn"] = r.Arn
	config["name"] = r.Name
	config["bandwidth"] = r.Bandwidth
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_elastictranscoder_preset struct {
	Description         *string
	Video_codec_options *map[string]interface{}
	Arn                 string
	Container           string
	Name                *string
	Resource_type       *string
}

func Aws_elastictranscoder_presetMapper(r *Aws_elastictranscoder_preset) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Video_codec_options != nil {
		config["video_codec_options"] = *r.Video_codec_options
	}
	if r.Description != nil {
		config["description"] = *r.Description
	}
	config["container"] = r.Container
	if r.Name != nil {
		config["name"] = *r.Name
	}
	if r.Resource_type != nil {
		config["resource_type"] = *r.Resource_type
	}
	config["arn"] = r.Arn
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_placement_group struct {
	Name     string
	Strategy string
}

func Aws_placement_groupMapper(r *Aws_placement_group) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	config["strategy"] = r.Strategy
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_waf_sql_injection_match_set struct {
	Name string
}

func Aws_waf_sql_injection_match_setMapper(r *Aws_waf_sql_injection_match_set) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_alb_listener struct {
	Arn               string
	Load_balancer_arn string
	Protocol          *string
	Ssl_policy        *string
	Certificate_arn   *string
}

func Aws_alb_listenerMapper(r *Aws_alb_listener) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Protocol != nil {
		config["protocol"] = *r.Protocol
	}
	if r.Ssl_policy != nil {
		config["ssl_policy"] = *r.Ssl_policy
	}
	if r.Certificate_arn != nil {
		config["certificate_arn"] = *r.Certificate_arn
	}
	config["arn"] = r.Arn
	config["load_balancer_arn"] = r.Load_balancer_arn
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_redshift_event_subscription struct {
	Source_type     *string
	Enabled         *bool
	Severity        *string
	Customer_aws_id string
	Tags            *map[string]interface{}
	Name            string
	Sns_topic_arn   string
	Status          string
}

func Aws_redshift_event_subscriptionMapper(r *Aws_redshift_event_subscription) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["customer_aws_id"] = r.Customer_aws_id
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	if r.Source_type != nil {
		config["source_type"] = *r.Source_type
	}
	if r.Enabled != nil {
		config["enabled"] = *r.Enabled
	}
	if r.Severity != nil {
		config["severity"] = *r.Severity
	}
	config["name"] = r.Name
	config["sns_topic_arn"] = r.Sns_topic_arn
	config["status"] = r.Status
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_route53_delegation_set struct {
	Reference_name *string
}

func Aws_route53_delegation_setMapper(r *Aws_route53_delegation_set) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Reference_name != nil {
		config["reference_name"] = *r.Reference_name
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_dx_private_virtual_interface struct {
	Arn                 string
	Connection_id       string
	Bgp_auth_key        *string
	Address_family      string
	Jumbo_frame_capable bool
	Tags                *map[string]interface{}
	Vpn_gateway_id      *string
	Customer_address    *string
	Name                string
	Dx_gateway_id       *string
	Amazon_address      *string
}

func Aws_dx_private_virtual_interfaceMapper(r *Aws_dx_private_virtual_interface) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Vpn_gateway_id != nil {
		config["vpn_gateway_id"] = *r.Vpn_gateway_id
	}
	if r.Customer_address != nil {
		config["customer_address"] = *r.Customer_address
	}
	config["name"] = r.Name
	if r.Dx_gateway_id != nil {
		config["dx_gateway_id"] = *r.Dx_gateway_id
	}
	if r.Amazon_address != nil {
		config["amazon_address"] = *r.Amazon_address
	}
	config["arn"] = r.Arn
	config["connection_id"] = r.Connection_id
	if r.Bgp_auth_key != nil {
		config["bgp_auth_key"] = *r.Bgp_auth_key
	}
	config["address_family"] = r.Address_family
	config["jumbo_frame_capable"] = r.Jumbo_frame_capable
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_gamelift_build struct {
	Version          *string
	Name             string
	Operating_system string
}

func Aws_gamelift_buildMapper(r *Aws_gamelift_build) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	config["operating_system"] = r.Operating_system
	if r.Version != nil {
		config["version"] = *r.Version
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_storagegateway_cached_iscsi_volume struct {
	Target_name          string
	Volume_arn           string
	Arn                  string
	Gateway_arn          string
	Snapshot_id          *string
	Target_arn           string
	Volume_id            string
	Chap_enabled         bool
	Network_interface_id string
	Source_volume_arn    *string
}

func Aws_storagegateway_cached_iscsi_volumeMapper(r *Aws_storagegateway_cached_iscsi_volume) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["volume_id"] = r.Volume_id
	config["chap_enabled"] = r.Chap_enabled
	config["network_interface_id"] = r.Network_interface_id
	if r.Source_volume_arn != nil {
		config["source_volume_arn"] = *r.Source_volume_arn
	}
	config["target_arn"] = r.Target_arn
	config["volume_arn"] = r.Volume_arn
	config["arn"] = r.Arn
	config["gateway_arn"] = r.Gateway_arn
	if r.Snapshot_id != nil {
		config["snapshot_id"] = *r.Snapshot_id
	}
	config["target_name"] = r.Target_name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_gamelift_alias struct {
	Name        string
	Description *string
	Arn         string
}

func Aws_gamelift_aliasMapper(r *Aws_gamelift_alias) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["arn"] = r.Arn
	config["name"] = r.Name
	if r.Description != nil {
		config["description"] = *r.Description
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_network_acl struct {
	Tags      *map[string]interface{}
	Owner_id  string
	Vpc_id    string
	Subnet_id *string
}

func Aws_network_aclMapper(r *Aws_network_acl) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["owner_id"] = r.Owner_id
	config["vpc_id"] = r.Vpc_id
	if r.Subnet_id != nil {
		config["subnet_id"] = *r.Subnet_id
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_alb struct {
	Zone_id                          string
	Name                             *string
	Name_prefix                      *string
	Enable_cross_zone_load_balancing *bool
	Internal                         *bool
	Load_balancer_type               *string
	Vpc_id                           string
	Tags                             *map[string]interface{}
	Arn                              string
	Arn_suffix                       string
	Enable_deletion_protection       *bool
	Dns_name                         string
	Enable_http2                     *bool
	Ip_address_type                  *string
}

func Aws_albMapper(r *Aws_alb) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Enable_http2 != nil {
		config["enable_http2"] = *r.Enable_http2
	}
	if r.Ip_address_type != nil {
		config["ip_address_type"] = *r.Ip_address_type
	}
	if r.Name != nil {
		config["name"] = *r.Name
	}
	if r.Name_prefix != nil {
		config["name_prefix"] = *r.Name_prefix
	}
	if r.Enable_cross_zone_load_balancing != nil {
		config["enable_cross_zone_load_balancing"] = *r.Enable_cross_zone_load_balancing
	}
	config["zone_id"] = r.Zone_id
	if r.Internal != nil {
		config["internal"] = *r.Internal
	}
	if r.Load_balancer_type != nil {
		config["load_balancer_type"] = *r.Load_balancer_type
	}
	config["vpc_id"] = r.Vpc_id
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["arn"] = r.Arn
	config["arn_suffix"] = r.Arn_suffix
	if r.Enable_deletion_protection != nil {
		config["enable_deletion_protection"] = *r.Enable_deletion_protection
	}
	config["dns_name"] = r.Dns_name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_acm_certificate struct {
	Certificate_chain *string
	Private_key       *string
	Validation_method *string
	Tags              *map[string]interface{}
	Certificate_body  *string
	Domain_name       *string
	Arn               string
}

func Aws_acm_certificateMapper(r *Aws_acm_certificate) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Certificate_body != nil {
		config["certificate_body"] = *r.Certificate_body
	}
	if r.Domain_name != nil {
		config["domain_name"] = *r.Domain_name
	}
	config["arn"] = r.Arn
	if r.Certificate_chain != nil {
		config["certificate_chain"] = *r.Certificate_chain
	}
	if r.Private_key != nil {
		config["private_key"] = *r.Private_key
	}
	if r.Validation_method != nil {
		config["validation_method"] = *r.Validation_method
	}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_efs_file_system struct {
	Tags             *map[string]interface{}
	Throughput_mode  *string
	Arn              string
	Reference_name   *string
	Performance_mode *string
	Kms_key_id       *string
	Creation_token   *string
	Encrypted        *bool
	Dns_name         string
}

func Aws_efs_file_systemMapper(r *Aws_efs_file_system) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Kms_key_id != nil {
		config["kms_key_id"] = *r.Kms_key_id
	}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	if r.Throughput_mode != nil {
		config["throughput_mode"] = *r.Throughput_mode
	}
	config["arn"] = r.Arn
	if r.Reference_name != nil {
		config["reference_name"] = *r.Reference_name
	}
	if r.Performance_mode != nil {
		config["performance_mode"] = *r.Performance_mode
	}
	if r.Creation_token != nil {
		config["creation_token"] = *r.Creation_token
	}
	if r.Encrypted != nil {
		config["encrypted"] = *r.Encrypted
	}
	config["dns_name"] = r.Dns_name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_iam_instance_profile struct {
	Create_date string
	Unique_id   string
	Name        *string
	Name_prefix *string
	Path        *string
	Role        *string
	Arn         string
}

func Aws_iam_instance_profileMapper(r *Aws_iam_instance_profile) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Name != nil {
		config["name"] = *r.Name
	}
	if r.Name_prefix != nil {
		config["name_prefix"] = *r.Name_prefix
	}
	if r.Path != nil {
		config["path"] = *r.Path
	}
	if r.Role != nil {
		config["role"] = *r.Role
	}
	config["arn"] = r.Arn
	config["create_date"] = r.Create_date
	config["unique_id"] = r.Unique_id
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_lb_listener_rule struct {
	Listener_arn string
	Arn          string
}

func Aws_lb_listener_ruleMapper(r *Aws_lb_listener_rule) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["arn"] = r.Arn
	config["listener_arn"] = r.Listener_arn
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_transfer_server struct {
	Invocation_role        *string
	Url                    *string
	Identity_provider_type *string
	Logging_role           *string
	Force_destroy          *bool
	Tags                   *map[string]interface{}
	Arn                    string
	Endpoint               string
}

func Aws_transfer_serverMapper(r *Aws_transfer_server) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["endpoint"] = r.Endpoint
	if r.Invocation_role != nil {
		config["invocation_role"] = *r.Invocation_role
	}
	if r.Url != nil {
		config["url"] = *r.Url
	}
	if r.Identity_provider_type != nil {
		config["identity_provider_type"] = *r.Identity_provider_type
	}
	if r.Logging_role != nil {
		config["logging_role"] = *r.Logging_role
	}
	if r.Force_destroy != nil {
		config["force_destroy"] = *r.Force_destroy
	}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["arn"] = r.Arn
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_config_configuration_aggregator struct {
	Arn  string
	Name string
}

func Aws_config_configuration_aggregatorMapper(r *Aws_config_configuration_aggregator) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["arn"] = r.Arn
	config["name"] = r.Name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_config_configuration_recorder struct {
	Name     *string
	Role_arn string
}

func Aws_config_configuration_recorderMapper(r *Aws_config_configuration_recorder) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Name != nil {
		config["name"] = *r.Name
	}
	config["role_arn"] = r.Role_arn
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_dynamodb_global_table struct {
	Arn  string
	Name string
}

func Aws_dynamodb_global_tableMapper(r *Aws_dynamodb_global_table) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["arn"] = r.Arn
	config["name"] = r.Name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_iam_policy struct {
	Description *string
	Path        *string
	Policy      string
	Name        *string
	Name_prefix *string
	Arn         string
}

func Aws_iam_policyMapper(r *Aws_iam_policy) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Description != nil {
		config["description"] = *r.Description
	}
	if r.Path != nil {
		config["path"] = *r.Path
	}
	config["policy"] = r.Policy
	if r.Name != nil {
		config["name"] = *r.Name
	}
	if r.Name_prefix != nil {
		config["name_prefix"] = *r.Name_prefix
	}
	config["arn"] = r.Arn
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_iam_saml_provider struct {
	Arn                    string
	Valid_until            string
	Name                   string
	Saml_metadata_document string
}

func Aws_iam_saml_providerMapper(r *Aws_iam_saml_provider) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	config["saml_metadata_document"] = r.Saml_metadata_document
	config["arn"] = r.Arn
	config["valid_until"] = r.Valid_until
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_cloud9_environment_ec2 struct {
	Description   *string
	Owner_arn     *string
	Subnet_id     *string
	Arn           string
	Resource_type string
	Name          string
	Instance_type string
}

func Aws_cloud9_environment_ec2Mapper(r *Aws_cloud9_environment_ec2) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Subnet_id != nil {
		config["subnet_id"] = *r.Subnet_id
	}
	config["arn"] = r.Arn
	config["resource_type"] = r.Resource_type
	config["name"] = r.Name
	config["instance_type"] = r.Instance_type
	if r.Description != nil {
		config["description"] = *r.Description
	}
	if r.Owner_arn != nil {
		config["owner_arn"] = *r.Owner_arn
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_iam_account_alias struct {
	Account_alias string
}

func Aws_iam_account_aliasMapper(r *Aws_iam_account_alias) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["account_alias"] = r.Account_alias
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_swf_domain struct {
	Name_prefix                                 *string
	Description                                 *string
	Workflow_execution_retention_period_in_days string
	Name                                        *string
}

func Aws_swf_domainMapper(r *Aws_swf_domain) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Description != nil {
		config["description"] = *r.Description
	}
	config["workflow_execution_retention_period_in_days"] = r.Workflow_execution_retention_period_in_days
	if r.Name != nil {
		config["name"] = *r.Name
	}
	if r.Name_prefix != nil {
		config["name_prefix"] = *r.Name_prefix
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_wafregional_web_acl struct {
	Name        string
	Metric_name string
}

func Aws_wafregional_web_aclMapper(r *Aws_wafregional_web_acl) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	config["metric_name"] = r.Metric_name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_appmesh_mesh struct {
	Name              string
	Arn               string
	Created_date      string
	Last_updated_date string
}

func Aws_appmesh_meshMapper(r *Aws_appmesh_mesh) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	config["arn"] = r.Arn
	config["created_date"] = r.Created_date
	config["last_updated_date"] = r.Last_updated_date
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_ecr_lifecycle_policy struct {
	Repository  string
	Policy      string
	Registry_id string
}

func Aws_ecr_lifecycle_policyMapper(r *Aws_ecr_lifecycle_policy) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["repository"] = r.Repository
	config["policy"] = r.Policy
	config["registry_id"] = r.Registry_id
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_proxy_protocol_policy struct {
	Load_balancer string
}

func Aws_proxy_protocol_policyMapper(r *Aws_proxy_protocol_policy) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["load_balancer"] = r.Load_balancer
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_cognito_identity_pool_roles_attachment struct {
	Identity_pool_id string
	Roles            map[string]interface{}
}

func Aws_cognito_identity_pool_roles_attachmentMapper(r *Aws_cognito_identity_pool_roles_attachment) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["identity_pool_id"] = r.Identity_pool_id
	config["roles"] = r.Roles
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_eip struct {
	Vpc                       *bool
	Network_interface         *string
	Allocation_id             string
	Private_ip                string
	Associate_with_private_ip *string
	Public_ipv4_pool          *string
	Tags                      *map[string]interface{}
	Instance                  *string
	Association_id            string
	Domain                    string
	Public_ip                 string
}

func Aws_eipMapper(r *Aws_eip) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Instance != nil {
		config["instance"] = *r.Instance
	}
	config["association_id"] = r.Association_id
	config["domain"] = r.Domain
	config["public_ip"] = r.Public_ip
	if r.Vpc != nil {
		config["vpc"] = *r.Vpc
	}
	if r.Network_interface != nil {
		config["network_interface"] = *r.Network_interface
	}
	config["allocation_id"] = r.Allocation_id
	config["private_ip"] = r.Private_ip
	if r.Associate_with_private_ip != nil {
		config["associate_with_private_ip"] = *r.Associate_with_private_ip
	}
	if r.Public_ipv4_pool != nil {
		config["public_ipv4_pool"] = *r.Public_ipv4_pool
	}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_kinesis_firehose_delivery_stream struct {
	Destination    string
	Arn            *string
	Destination_id *string
	Tags           *map[string]interface{}
	Version_id     *string
	Name           string
}

func Aws_kinesis_firehose_delivery_streamMapper(r *Aws_kinesis_firehose_delivery_stream) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	if r.Version_id != nil {
		config["version_id"] = *r.Version_id
	}
	config["name"] = r.Name
	config["destination"] = r.Destination
	if r.Arn != nil {
		config["arn"] = *r.Arn
	}
	if r.Destination_id != nil {
		config["destination_id"] = *r.Destination_id
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_sfn_activity struct {
	Tags          *map[string]interface{}
	Name          string
	Creation_date string
}

func Aws_sfn_activityMapper(r *Aws_sfn_activity) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	config["creation_date"] = r.Creation_date
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_db_subnet_group struct {
	Name_prefix *string
	Description *string
	Tags        *map[string]interface{}
	Arn         string
	Name        *string
}

func Aws_db_subnet_groupMapper(r *Aws_db_subnet_group) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["arn"] = r.Arn
	if r.Name != nil {
		config["name"] = *r.Name
	}
	if r.Name_prefix != nil {
		config["name_prefix"] = *r.Name_prefix
	}
	if r.Description != nil {
		config["description"] = *r.Description
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_mq_broker struct {
	Engine_type                string
	Host_instance_type         string
	Publicly_accessible        *bool
	Apply_immediately          *bool
	Auto_minor_version_upgrade *bool
	Deployment_mode            *string
	Arn                        string
	Broker_name                string
	Engine_version             string
}

func Aws_mq_brokerMapper(r *Aws_mq_broker) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["engine_type"] = r.Engine_type
	config["host_instance_type"] = r.Host_instance_type
	if r.Publicly_accessible != nil {
		config["publicly_accessible"] = *r.Publicly_accessible
	}
	if r.Apply_immediately != nil {
		config["apply_immediately"] = *r.Apply_immediately
	}
	if r.Auto_minor_version_upgrade != nil {
		config["auto_minor_version_upgrade"] = *r.Auto_minor_version_upgrade
	}
	if r.Deployment_mode != nil {
		config["deployment_mode"] = *r.Deployment_mode
	}
	config["arn"] = r.Arn
	config["broker_name"] = r.Broker_name
	config["engine_version"] = r.Engine_version
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_cloudwatch_log_stream struct {
	Arn            string
	Name           string
	Log_group_name string
}

func Aws_cloudwatch_log_streamMapper(r *Aws_cloudwatch_log_stream) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["arn"] = r.Arn
	config["name"] = r.Name
	config["log_group_name"] = r.Log_group_name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_ami_from_instance struct {
	Image_location          string
	Manage_ebs_snapshots    bool
	Root_device_name        string
	Root_snapshot_id        string
	Kernel_id               string
	Ramdisk_id              string
	Tags                    *map[string]interface{}
	Virtualization_type     string
	Architecture            string
	Name                    string
	Source_instance_id      string
	Sriov_net_support       string
	Description             *string
	Ena_support             bool
	Snapshot_without_reboot *bool
}

func Aws_ami_from_instanceMapper(r *Aws_ami_from_instance) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["architecture"] = r.Architecture
	config["name"] = r.Name
	config["source_instance_id"] = r.Source_instance_id
	config["sriov_net_support"] = r.Sriov_net_support
	if r.Description != nil {
		config["description"] = *r.Description
	}
	config["ena_support"] = r.Ena_support
	if r.Snapshot_without_reboot != nil {
		config["snapshot_without_reboot"] = *r.Snapshot_without_reboot
	}
	config["image_location"] = r.Image_location
	config["manage_ebs_snapshots"] = r.Manage_ebs_snapshots
	config["root_device_name"] = r.Root_device_name
	config["root_snapshot_id"] = r.Root_snapshot_id
	config["kernel_id"] = r.Kernel_id
	config["ramdisk_id"] = r.Ramdisk_id
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["virtualization_type"] = r.Virtualization_type
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_dx_public_virtual_interface struct {
	Name             string
	Address_family   string
	Customer_address *string
	Tags             *map[string]interface{}
	Arn              string
	Bgp_auth_key     *string
	Amazon_address   *string
	Connection_id    string
}

func Aws_dx_public_virtual_interfaceMapper(r *Aws_dx_public_virtual_interface) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["address_family"] = r.Address_family
	if r.Customer_address != nil {
		config["customer_address"] = *r.Customer_address
	}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["arn"] = r.Arn
	config["name"] = r.Name
	if r.Bgp_auth_key != nil {
		config["bgp_auth_key"] = *r.Bgp_auth_key
	}
	if r.Amazon_address != nil {
		config["amazon_address"] = *r.Amazon_address
	}
	config["connection_id"] = r.Connection_id
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_lightsail_instance struct {
	Bundle_id          string
	User_data          *string
	Name               string
	Availability_zone  string
	Is_static_ip       bool
	Public_ip_address  string
	Created_at         string
	Ipv6_address       string
	Private_ip_address string
	Blueprint_id       string
	Key_pair_name      *string
	Username           string
	Arn                string
}

func Aws_lightsail_instanceMapper(r *Aws_lightsail_instance) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["is_static_ip"] = r.Is_static_ip
	config["public_ip_address"] = r.Public_ip_address
	config["private_ip_address"] = r.Private_ip_address
	config["blueprint_id"] = r.Blueprint_id
	if r.Key_pair_name != nil {
		config["key_pair_name"] = *r.Key_pair_name
	}
	config["created_at"] = r.Created_at
	config["ipv6_address"] = r.Ipv6_address
	config["arn"] = r.Arn
	config["username"] = r.Username
	config["name"] = r.Name
	config["availability_zone"] = r.Availability_zone
	config["bundle_id"] = r.Bundle_id
	if r.User_data != nil {
		config["user_data"] = *r.User_data
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_alb_listener_rule struct {
	Arn          string
	Listener_arn string
}

func Aws_alb_listener_ruleMapper(r *Aws_alb_listener_rule) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["arn"] = r.Arn
	config["listener_arn"] = r.Listener_arn
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_appsync_graphql_api struct {
	Arn                 string
	Uris                map[string]interface{}
	Authentication_type string
	Name                string
}

func Aws_appsync_graphql_apiMapper(r *Aws_appsync_graphql_api) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["uris"] = r.Uris
	config["authentication_type"] = r.Authentication_type
	config["name"] = r.Name
	config["arn"] = r.Arn
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_lambda_event_source_mapping struct {
	Enabled                     *bool
	Function_arn                string
	Last_modified               string
	Last_processing_result      string
	State                       string
	Event_source_arn            string
	Function_name               string
	Starting_position_timestamp *string
	Uuid                        string
	Starting_position           *string
	State_transition_reason     string
}

func Aws_lambda_event_source_mappingMapper(r *Aws_lambda_event_source_mapping) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["last_processing_result"] = r.Last_processing_result
	config["state"] = r.State
	config["event_source_arn"] = r.Event_source_arn
	config["function_name"] = r.Function_name
	if r.Starting_position_timestamp != nil {
		config["starting_position_timestamp"] = *r.Starting_position_timestamp
	}
	if r.Enabled != nil {
		config["enabled"] = *r.Enabled
	}
	config["function_arn"] = r.Function_arn
	config["last_modified"] = r.Last_modified
	config["uuid"] = r.Uuid
	if r.Starting_position != nil {
		config["starting_position"] = *r.Starting_position
	}
	config["state_transition_reason"] = r.State_transition_reason
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_opsworks_memcached_layer struct {
	Custom_json                 *string
	Auto_assign_elastic_ips     *bool
	Stack_id                    string
	Use_ebs_optimized_instances *bool
	Auto_assign_public_ips      *bool
	Auto_healing                *bool
	Name                        *string
	Elastic_load_balancer       *string
	Install_updates_on_boot     *bool
	Drain_elb_on_shutdown       *bool
	Custom_instance_profile_arn *string
}

func Aws_opsworks_memcached_layerMapper(r *Aws_opsworks_memcached_layer) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Custom_json != nil {
		config["custom_json"] = *r.Custom_json
	}
	if r.Auto_assign_elastic_ips != nil {
		config["auto_assign_elastic_ips"] = *r.Auto_assign_elastic_ips
	}
	config["stack_id"] = r.Stack_id
	if r.Use_ebs_optimized_instances != nil {
		config["use_ebs_optimized_instances"] = *r.Use_ebs_optimized_instances
	}
	if r.Auto_assign_public_ips != nil {
		config["auto_assign_public_ips"] = *r.Auto_assign_public_ips
	}
	if r.Auto_healing != nil {
		config["auto_healing"] = *r.Auto_healing
	}
	if r.Name != nil {
		config["name"] = *r.Name
	}
	if r.Elastic_load_balancer != nil {
		config["elastic_load_balancer"] = *r.Elastic_load_balancer
	}
	if r.Install_updates_on_boot != nil {
		config["install_updates_on_boot"] = *r.Install_updates_on_boot
	}
	if r.Drain_elb_on_shutdown != nil {
		config["drain_elb_on_shutdown"] = *r.Drain_elb_on_shutdown
	}
	if r.Custom_instance_profile_arn != nil {
		config["custom_instance_profile_arn"] = *r.Custom_instance_profile_arn
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_s3_bucket_object struct {
	Kms_key_id             *string
	Bucket                 string
	Key                    string
	Content_type           *string
	Source                 *string
	Tags                   *map[string]interface{}
	Acl                    *string
	Content_disposition    *string
	Storage_class          *string
	Etag                   *string
	Cache_control          *string
	Content_encoding       *string
	Content                *string
	Website_redirect       *string
	Content_language       *string
	Content_base64         *string
	Server_side_encryption *string
	Version_id             string
}

func Aws_s3_bucket_objectMapper(r *Aws_s3_bucket_object) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Cache_control != nil {
		config["cache_control"] = *r.Cache_control
	}
	if r.Content_encoding != nil {
		config["content_encoding"] = *r.Content_encoding
	}
	if r.Content != nil {
		config["content"] = *r.Content
	}
	config["version_id"] = r.Version_id
	if r.Website_redirect != nil {
		config["website_redirect"] = *r.Website_redirect
	}
	if r.Content_language != nil {
		config["content_language"] = *r.Content_language
	}
	if r.Content_base64 != nil {
		config["content_base64"] = *r.Content_base64
	}
	if r.Server_side_encryption != nil {
		config["server_side_encryption"] = *r.Server_side_encryption
	}
	if r.Source != nil {
		config["source"] = *r.Source
	}
	if r.Kms_key_id != nil {
		config["kms_key_id"] = *r.Kms_key_id
	}
	config["bucket"] = r.Bucket
	config["key"] = r.Key
	if r.Content_type != nil {
		config["content_type"] = *r.Content_type
	}
	if r.Etag != nil {
		config["etag"] = *r.Etag
	}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	if r.Acl != nil {
		config["acl"] = *r.Acl
	}
	if r.Content_disposition != nil {
		config["content_disposition"] = *r.Content_disposition
	}
	if r.Storage_class != nil {
		config["storage_class"] = *r.Storage_class
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_spot_instance_request struct {
	Instance_initiated_shutdown_behavior *string
	Monitoring                           *bool
	Tenancy                              *string
	Volume_tags                          *map[string]interface{}
	User_data                            *string
	Public_dns                           string
	Instance_state                       string
	Private_dns                          string
	Host_id                              *string
	Wait_for_fulfillment                 *bool
	Get_password_data                    *bool
	Private_ip                           *string
	Launch_group                         *string
	Valid_until                          *string
	Spot_bid_status                      string
	Instance_type                        string
	Ebs_optimized                        *bool
	Iam_instance_profile                 *string
	Primary_network_interface_id         string
	Instance_interruption_behaviour      *string
	Disable_api_termination              *bool
	Tags                                 *map[string]interface{}
	Ami                                  string
	Availability_zone                    *string
	Password_data                        string
	Subnet_id                            *string
	Spot_type                            *string
	Valid_from                           *string
	Spot_instance_id                     string
	Arn                                  string
	Placement_group                      *string
	Spot_price                           *string
	User_data_base64                     *string
	Network_interface_id                 string
	Public_ip                            string
	Associate_public_ip_address          *bool
	Key_name                             *string
	Source_dest_check                    *bool
	Spot_request_state                   string
	Block_device                         *map[string]interface{}
}

func Aws_spot_instance_requestMapper(r *Aws_spot_instance_request) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["public_ip"] = r.Public_ip
	if r.Associate_public_ip_address != nil {
		config["associate_public_ip_address"] = *r.Associate_public_ip_address
	}
	if r.Key_name != nil {
		config["key_name"] = *r.Key_name
	}
	if r.Source_dest_check != nil {
		config["source_dest_check"] = *r.Source_dest_check
	}
	if r.User_data_base64 != nil {
		config["user_data_base64"] = *r.User_data_base64
	}
	config["network_interface_id"] = r.Network_interface_id
	if r.Block_device != nil {
		config["block_device"] = *r.Block_device
	}
	config["spot_request_state"] = r.Spot_request_state
	if r.Volume_tags != nil {
		config["volume_tags"] = *r.Volume_tags
	}
	if r.User_data != nil {
		config["user_data"] = *r.User_data
	}
	config["public_dns"] = r.Public_dns
	config["instance_state"] = r.Instance_state
	if r.Instance_initiated_shutdown_behavior != nil {
		config["instance_initiated_shutdown_behavior"] = *r.Instance_initiated_shutdown_behavior
	}
	if r.Monitoring != nil {
		config["monitoring"] = *r.Monitoring
	}
	if r.Tenancy != nil {
		config["tenancy"] = *r.Tenancy
	}
	if r.Wait_for_fulfillment != nil {
		config["wait_for_fulfillment"] = *r.Wait_for_fulfillment
	}
	if r.Get_password_data != nil {
		config["get_password_data"] = *r.Get_password_data
	}
	if r.Private_ip != nil {
		config["private_ip"] = *r.Private_ip
	}
	config["private_dns"] = r.Private_dns
	if r.Host_id != nil {
		config["host_id"] = *r.Host_id
	}
	if r.Launch_group != nil {
		config["launch_group"] = *r.Launch_group
	}
	if r.Valid_until != nil {
		config["valid_until"] = *r.Valid_until
	}
	config["instance_type"] = r.Instance_type
	if r.Ebs_optimized != nil {
		config["ebs_optimized"] = *r.Ebs_optimized
	}
	if r.Iam_instance_profile != nil {
		config["iam_instance_profile"] = *r.Iam_instance_profile
	}
	config["spot_bid_status"] = r.Spot_bid_status
	config["primary_network_interface_id"] = r.Primary_network_interface_id
	if r.Instance_interruption_behaviour != nil {
		config["instance_interruption_behaviour"] = *r.Instance_interruption_behaviour
	}
	config["ami"] = r.Ami
	if r.Availability_zone != nil {
		config["availability_zone"] = *r.Availability_zone
	}
	config["password_data"] = r.Password_data
	if r.Disable_api_termination != nil {
		config["disable_api_termination"] = *r.Disable_api_termination
	}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	if r.Subnet_id != nil {
		config["subnet_id"] = *r.Subnet_id
	}
	if r.Spot_type != nil {
		config["spot_type"] = *r.Spot_type
	}
	if r.Valid_from != nil {
		config["valid_from"] = *r.Valid_from
	}
	config["arn"] = r.Arn
	if r.Placement_group != nil {
		config["placement_group"] = *r.Placement_group
	}
	if r.Spot_price != nil {
		config["spot_price"] = *r.Spot_price
	}
	config["spot_instance_id"] = r.Spot_instance_id
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_codepipeline_webhook struct {
	Name            string
	Url             string
	Target_action   string
	Target_pipeline string
	Authentication  string
}

func Aws_codepipeline_webhookMapper(r *Aws_codepipeline_webhook) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["authentication"] = r.Authentication
	config["name"] = r.Name
	config["url"] = r.Url
	config["target_action"] = r.Target_action
	config["target_pipeline"] = r.Target_pipeline
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_lambda_alias struct {
	Invoke_arn       string
	Description      *string
	Function_name    string
	Function_version string
	Name             string
	Arn              string
}

func Aws_lambda_aliasMapper(r *Aws_lambda_alias) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Description != nil {
		config["description"] = *r.Description
	}
	config["function_name"] = r.Function_name
	config["function_version"] = r.Function_version
	config["name"] = r.Name
	config["arn"] = r.Arn
	config["invoke_arn"] = r.Invoke_arn
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_rds_global_cluster struct {
	Database_name              *string
	Deletion_protection        *bool
	Engine                     *string
	Engine_version             *string
	Global_cluster_identifier  string
	Global_cluster_resource_id string
	Storage_encrypted          *bool
	Arn                        string
}

func Aws_rds_global_clusterMapper(r *Aws_rds_global_cluster) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Engine_version != nil {
		config["engine_version"] = *r.Engine_version
	}
	config["global_cluster_identifier"] = r.Global_cluster_identifier
	config["global_cluster_resource_id"] = r.Global_cluster_resource_id
	if r.Storage_encrypted != nil {
		config["storage_encrypted"] = *r.Storage_encrypted
	}
	config["arn"] = r.Arn
	if r.Database_name != nil {
		config["database_name"] = *r.Database_name
	}
	if r.Deletion_protection != nil {
		config["deletion_protection"] = *r.Deletion_protection
	}
	if r.Engine != nil {
		config["engine"] = *r.Engine
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_sns_platform_application struct {
	Platform_principal               *string
	Platform_credential              string
	Failure_feedback_role_arn        *string
	Arn                              string
	Event_delivery_failure_topic_arn *string
	Event_endpoint_created_topic_arn *string
	Event_endpoint_deleted_topic_arn *string
	Event_endpoint_updated_topic_arn *string
	Success_feedback_role_arn        *string
	Name                             string
	Platform                         string
	Success_feedback_sample_rate     *string
}

func Aws_sns_platform_applicationMapper(r *Aws_sns_platform_application) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["platform_credential"] = r.Platform_credential
	if r.Failure_feedback_role_arn != nil {
		config["failure_feedback_role_arn"] = *r.Failure_feedback_role_arn
	}
	if r.Platform_principal != nil {
		config["platform_principal"] = *r.Platform_principal
	}
	if r.Event_endpoint_created_topic_arn != nil {
		config["event_endpoint_created_topic_arn"] = *r.Event_endpoint_created_topic_arn
	}
	if r.Event_endpoint_deleted_topic_arn != nil {
		config["event_endpoint_deleted_topic_arn"] = *r.Event_endpoint_deleted_topic_arn
	}
	if r.Event_endpoint_updated_topic_arn != nil {
		config["event_endpoint_updated_topic_arn"] = *r.Event_endpoint_updated_topic_arn
	}
	if r.Success_feedback_role_arn != nil {
		config["success_feedback_role_arn"] = *r.Success_feedback_role_arn
	}
	config["name"] = r.Name
	config["platform"] = r.Platform
	config["arn"] = r.Arn
	if r.Event_delivery_failure_topic_arn != nil {
		config["event_delivery_failure_topic_arn"] = *r.Event_delivery_failure_topic_arn
	}
	if r.Success_feedback_sample_rate != nil {
		config["success_feedback_sample_rate"] = *r.Success_feedback_sample_rate
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_vpc_endpoint struct {
	Vpc_endpoint_type   *string
	Private_dns_enabled *bool
	State               string
	Service_name        string
	Prefix_list_id      string
	Auto_accept         *bool
	Vpc_id              string
	Policy              *string
}

func Aws_vpc_endpointMapper(r *Aws_vpc_endpoint) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["service_name"] = r.Service_name
	config["vpc_id"] = r.Vpc_id
	if r.Policy != nil {
		config["policy"] = *r.Policy
	}
	config["prefix_list_id"] = r.Prefix_list_id
	if r.Auto_accept != nil {
		config["auto_accept"] = *r.Auto_accept
	}
	if r.Vpc_endpoint_type != nil {
		config["vpc_endpoint_type"] = *r.Vpc_endpoint_type
	}
	if r.Private_dns_enabled != nil {
		config["private_dns_enabled"] = *r.Private_dns_enabled
	}
	config["state"] = r.State
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_config_delivery_channel struct {
	Name           *string
	S3_bucket_name string
	S3_key_prefix  *string
	Sns_topic_arn  *string
}

func Aws_config_delivery_channelMapper(r *Aws_config_delivery_channel) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Name != nil {
		config["name"] = *r.Name
	}
	config["s3_bucket_name"] = r.S3_bucket_name
	if r.S3_key_prefix != nil {
		config["s3_key_prefix"] = *r.S3_key_prefix
	}
	if r.Sns_topic_arn != nil {
		config["sns_topic_arn"] = *r.Sns_topic_arn
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_db_security_group struct {
	Description *string
	Tags        *map[string]interface{}
	Arn         string
	Name        string
}

func Aws_db_security_groupMapper(r *Aws_db_security_group) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	if r.Description != nil {
		config["description"] = *r.Description
	}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["arn"] = r.Arn
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_lb_ssl_negotiation_policy struct {
	Name          string
	Load_balancer string
}

func Aws_lb_ssl_negotiation_policyMapper(r *Aws_lb_ssl_negotiation_policy) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["load_balancer"] = r.Load_balancer
	config["name"] = r.Name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_ses_receipt_filter struct {
	Policy string
	Name   string
	Cidr   string
}

func Aws_ses_receipt_filterMapper(r *Aws_ses_receipt_filter) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	config["cidr"] = r.Cidr
	config["policy"] = r.Policy
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_sns_topic struct {
	Http_success_feedback_role_arn        *string
	Lambda_success_feedback_role_arn      *string
	Arn                                   string
	Name_prefix                           *string
	Application_success_feedback_role_arn *string
	Kms_master_key_id                     *string
	Sqs_success_feedback_role_arn         *string
	Application_failure_feedback_role_arn *string
	Lambda_failure_feedback_role_arn      *string
	Name                                  *string
	Delivery_policy                       *string
	Http_failure_feedback_role_arn        *string
	Sqs_failure_feedback_role_arn         *string
	Display_name                          *string
	Policy                                *string
}

func Aws_sns_topicMapper(r *Aws_sns_topic) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Sqs_success_feedback_role_arn != nil {
		config["sqs_success_feedback_role_arn"] = *r.Sqs_success_feedback_role_arn
	}
	if r.Kms_master_key_id != nil {
		config["kms_master_key_id"] = *r.Kms_master_key_id
	}
	if r.Lambda_failure_feedback_role_arn != nil {
		config["lambda_failure_feedback_role_arn"] = *r.Lambda_failure_feedback_role_arn
	}
	if r.Name != nil {
		config["name"] = *r.Name
	}
	if r.Delivery_policy != nil {
		config["delivery_policy"] = *r.Delivery_policy
	}
	if r.Application_failure_feedback_role_arn != nil {
		config["application_failure_feedback_role_arn"] = *r.Application_failure_feedback_role_arn
	}
	if r.Sqs_failure_feedback_role_arn != nil {
		config["sqs_failure_feedback_role_arn"] = *r.Sqs_failure_feedback_role_arn
	}
	if r.Display_name != nil {
		config["display_name"] = *r.Display_name
	}
	if r.Policy != nil {
		config["policy"] = *r.Policy
	}
	if r.Http_failure_feedback_role_arn != nil {
		config["http_failure_feedback_role_arn"] = *r.Http_failure_feedback_role_arn
	}
	if r.Lambda_success_feedback_role_arn != nil {
		config["lambda_success_feedback_role_arn"] = *r.Lambda_success_feedback_role_arn
	}
	config["arn"] = r.Arn
	if r.Name_prefix != nil {
		config["name_prefix"] = *r.Name_prefix
	}
	if r.Application_success_feedback_role_arn != nil {
		config["application_success_feedback_role_arn"] = *r.Application_success_feedback_role_arn
	}
	if r.Http_success_feedback_role_arn != nil {
		config["http_success_feedback_role_arn"] = *r.Http_success_feedback_role_arn
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_acmpca_certificate_authority struct {
	Status                      string
	Tags                        *map[string]interface{}
	Arn                         string
	Not_before                  string
	Serial                      string
	Not_after                   string
	Resource_type               *string
	Certificate                 string
	Certificate_chain           string
	Certificate_signing_request string
	Enabled                     *bool
}

func Aws_acmpca_certificate_authorityMapper(r *Aws_acmpca_certificate_authority) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["serial"] = r.Serial
	config["status"] = r.Status
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["arn"] = r.Arn
	config["not_before"] = r.Not_before
	if r.Enabled != nil {
		config["enabled"] = *r.Enabled
	}
	config["not_after"] = r.Not_after
	if r.Resource_type != nil {
		config["resource_type"] = *r.Resource_type
	}
	config["certificate"] = r.Certificate
	config["certificate_chain"] = r.Certificate_chain
	config["certificate_signing_request"] = r.Certificate_signing_request
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_config_aggregate_authorization struct {
	Arn        string
	Account_id string
	Region     string
}

func Aws_config_aggregate_authorizationMapper(r *Aws_config_aggregate_authorization) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["arn"] = r.Arn
	config["account_id"] = r.Account_id
	config["region"] = r.Region
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_default_route_table struct {
	Tags                   *map[string]interface{}
	Owner_id               string
	Default_route_table_id string
	Vpc_id                 string
}

func Aws_default_route_tableMapper(r *Aws_default_route_table) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["default_route_table_id"] = r.Default_route_table_id
	config["vpc_id"] = r.Vpc_id
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["owner_id"] = r.Owner_id
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_wafregional_xss_match_set struct {
	Name string
}

func Aws_wafregional_xss_match_setMapper(r *Aws_wafregional_xss_match_set) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_route53_query_log struct {
	Cloudwatch_log_group_arn string
	Zone_id                  string
}

func Aws_route53_query_logMapper(r *Aws_route53_query_log) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["cloudwatch_log_group_arn"] = r.Cloudwatch_log_group_arn
	config["zone_id"] = r.Zone_id
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_s3_bucket_inventory struct {
	Name                     string
	Enabled                  *bool
	Included_object_versions string
	Bucket                   string
}

func Aws_s3_bucket_inventoryMapper(r *Aws_s3_bucket_inventory) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["included_object_versions"] = r.Included_object_versions
	config["bucket"] = r.Bucket
	config["name"] = r.Name
	if r.Enabled != nil {
		config["enabled"] = *r.Enabled
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_sagemaker_notebook_instance struct {
	Kms_key_id    *string
	Tags          *map[string]interface{}
	Arn           string
	Name          string
	Role_arn      string
	Instance_type string
	Subnet_id     *string
}

func Aws_sagemaker_notebook_instanceMapper(r *Aws_sagemaker_notebook_instance) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["arn"] = r.Arn
	config["name"] = r.Name
	config["role_arn"] = r.Role_arn
	config["instance_type"] = r.Instance_type
	if r.Subnet_id != nil {
		config["subnet_id"] = *r.Subnet_id
	}
	if r.Kms_key_id != nil {
		config["kms_key_id"] = *r.Kms_key_id
	}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_waf_ipset struct {
	Name string
	Arn  string
}

func Aws_waf_ipsetMapper(r *Aws_waf_ipset) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	config["arn"] = r.Arn
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_vpc_peering_connection struct {
	Peer_owner_id *string
	Vpc_id        string
	Tags          *map[string]interface{}
	Peer_vpc_id   string
	Auto_accept   *bool
	Accept_status string
	Peer_region   *string
}

func Aws_vpc_peering_connectionMapper(r *Aws_vpc_peering_connection) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["peer_vpc_id"] = r.Peer_vpc_id
	if r.Auto_accept != nil {
		config["auto_accept"] = *r.Auto_accept
	}
	config["accept_status"] = r.Accept_status
	if r.Peer_region != nil {
		config["peer_region"] = *r.Peer_region
	}
	if r.Peer_owner_id != nil {
		config["peer_owner_id"] = *r.Peer_owner_id
	}
	config["vpc_id"] = r.Vpc_id
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_api_gateway_account struct {
	Cloudwatch_role_arn *string
}

func Aws_api_gateway_accountMapper(r *Aws_api_gateway_account) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Cloudwatch_role_arn != nil {
		config["cloudwatch_role_arn"] = *r.Cloudwatch_role_arn
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_codecommit_repository struct {
	Repository_name string
	Description     *string
	Arn             string
	Repository_id   string
	Clone_url_http  string
	Clone_url_ssh   string
	Default_branch  *string
}

func Aws_codecommit_repositoryMapper(r *Aws_codecommit_repository) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["clone_url_ssh"] = r.Clone_url_ssh
	if r.Default_branch != nil {
		config["default_branch"] = *r.Default_branch
	}
	config["repository_name"] = r.Repository_name
	if r.Description != nil {
		config["description"] = *r.Description
	}
	config["arn"] = r.Arn
	config["repository_id"] = r.Repository_id
	config["clone_url_http"] = r.Clone_url_http
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_dx_gateway_association struct {
	Dx_gateway_id  string
	Vpn_gateway_id string
}

func Aws_dx_gateway_associationMapper(r *Aws_dx_gateway_association) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["dx_gateway_id"] = r.Dx_gateway_id
	config["vpn_gateway_id"] = r.Vpn_gateway_id
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_launch_template struct {
	Image_id                             *string
	Key_name                             *string
	Ram_disk_id                          *string
	User_data                            *string
	Name                                 *string
	Disable_api_termination              *bool
	Tags                                 *map[string]interface{}
	Kernel_id                            *string
	Arn                                  string
	Instance_initiated_shutdown_behavior *string
	Name_prefix                          *string
	Description                          *string
	Instance_type                        *string
	Ebs_optimized                        *string
}

func Aws_launch_templateMapper(r *Aws_launch_template) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.User_data != nil {
		config["user_data"] = *r.User_data
	}
	if r.Name != nil {
		config["name"] = *r.Name
	}
	if r.Disable_api_termination != nil {
		config["disable_api_termination"] = *r.Disable_api_termination
	}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	if r.Kernel_id != nil {
		config["kernel_id"] = *r.Kernel_id
	}
	config["arn"] = r.Arn
	if r.Instance_initiated_shutdown_behavior != nil {
		config["instance_initiated_shutdown_behavior"] = *r.Instance_initiated_shutdown_behavior
	}
	if r.Name_prefix != nil {
		config["name_prefix"] = *r.Name_prefix
	}
	if r.Description != nil {
		config["description"] = *r.Description
	}
	if r.Instance_type != nil {
		config["instance_type"] = *r.Instance_type
	}
	if r.Ebs_optimized != nil {
		config["ebs_optimized"] = *r.Ebs_optimized
	}
	if r.Image_id != nil {
		config["image_id"] = *r.Image_id
	}
	if r.Key_name != nil {
		config["key_name"] = *r.Key_name
	}
	if r.Ram_disk_id != nil {
		config["ram_disk_id"] = *r.Ram_disk_id
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_opsworks_mysql_layer struct {
	Stack_id                       string
	Auto_assign_elastic_ips        *bool
	Custom_json                    *string
	Auto_healing                   *bool
	Drain_elb_on_shutdown          *bool
	Auto_assign_public_ips         *bool
	Elastic_load_balancer          *string
	Root_password_on_all_instances *bool
	Install_updates_on_boot        *bool
	Use_ebs_optimized_instances    *bool
	Name                           *string
	Root_password                  *string
	Custom_instance_profile_arn    *string
}

func Aws_opsworks_mysql_layerMapper(r *Aws_opsworks_mysql_layer) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Auto_assign_elastic_ips != nil {
		config["auto_assign_elastic_ips"] = *r.Auto_assign_elastic_ips
	}
	if r.Custom_json != nil {
		config["custom_json"] = *r.Custom_json
	}
	config["stack_id"] = r.Stack_id
	if r.Auto_assign_public_ips != nil {
		config["auto_assign_public_ips"] = *r.Auto_assign_public_ips
	}
	if r.Elastic_load_balancer != nil {
		config["elastic_load_balancer"] = *r.Elastic_load_balancer
	}
	if r.Auto_healing != nil {
		config["auto_healing"] = *r.Auto_healing
	}
	if r.Drain_elb_on_shutdown != nil {
		config["drain_elb_on_shutdown"] = *r.Drain_elb_on_shutdown
	}
	if r.Install_updates_on_boot != nil {
		config["install_updates_on_boot"] = *r.Install_updates_on_boot
	}
	if r.Root_password_on_all_instances != nil {
		config["root_password_on_all_instances"] = *r.Root_password_on_all_instances
	}
	if r.Name != nil {
		config["name"] = *r.Name
	}
	if r.Root_password != nil {
		config["root_password"] = *r.Root_password
	}
	if r.Custom_instance_profile_arn != nil {
		config["custom_instance_profile_arn"] = *r.Custom_instance_profile_arn
	}
	if r.Use_ebs_optimized_instances != nil {
		config["use_ebs_optimized_instances"] = *r.Use_ebs_optimized_instances
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_wafregional_rate_based_rule struct {
	Name        string
	Metric_name string
	Rate_key    string
}

func Aws_wafregional_rate_based_ruleMapper(r *Aws_wafregional_rate_based_rule) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["rate_key"] = r.Rate_key
	config["name"] = r.Name
	config["metric_name"] = r.Metric_name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_dax_subnet_group struct {
	Name        string
	Description *string
	Vpc_id      string
}

func Aws_dax_subnet_groupMapper(r *Aws_dax_subnet_group) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	if r.Description != nil {
		config["description"] = *r.Description
	}
	config["vpc_id"] = r.Vpc_id
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_ecs_cluster struct {
	Arn  string
	Name string
	Tags *map[string]interface{}
}

func Aws_ecs_clusterMapper(r *Aws_ecs_cluster) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["arn"] = r.Arn
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_iam_group struct {
	Arn       string
	Unique_id string
	Name      string
	Path      *string
}

func Aws_iam_groupMapper(r *Aws_iam_group) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Path != nil {
		config["path"] = *r.Path
	}
	config["arn"] = r.Arn
	config["unique_id"] = r.Unique_id
	config["name"] = r.Name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_iot_certificate struct {
	Csr    string
	Active bool
	Arn    string
}

func Aws_iot_certificateMapper(r *Aws_iot_certificate) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["csr"] = r.Csr
	config["active"] = r.Active
	config["arn"] = r.Arn
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_storagegateway_upload_buffer struct {
	Disk_id     string
	Gateway_arn string
}

func Aws_storagegateway_upload_bufferMapper(r *Aws_storagegateway_upload_buffer) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["disk_id"] = r.Disk_id
	config["gateway_arn"] = r.Gateway_arn
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_api_gateway_model struct {
	Rest_api_id  string
	Name         string
	Description  *string
	Schema       *string
	Content_type string
}

func Aws_api_gateway_modelMapper(r *Aws_api_gateway_model) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["content_type"] = r.Content_type
	config["rest_api_id"] = r.Rest_api_id
	config["name"] = r.Name
	if r.Description != nil {
		config["description"] = *r.Description
	}
	if r.Schema != nil {
		config["schema"] = *r.Schema
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_cloudwatch_log_subscription_filter struct {
	Role_arn        *string
	Distribution    *string
	Name            string
	Destination_arn string
	Filter_pattern  string
	Log_group_name  string
}

func Aws_cloudwatch_log_subscription_filterMapper(r *Aws_cloudwatch_log_subscription_filter) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["destination_arn"] = r.Destination_arn
	config["filter_pattern"] = r.Filter_pattern
	config["log_group_name"] = r.Log_group_name
	if r.Role_arn != nil {
		config["role_arn"] = *r.Role_arn
	}
	if r.Distribution != nil {
		config["distribution"] = *r.Distribution
	}
	config["name"] = r.Name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_efs_mount_target struct {
	Network_interface_id string
	Dns_name             string
	File_system_arn      string
	File_system_id       string
	Ip_address           *string
	Subnet_id            string
}

func Aws_efs_mount_targetMapper(r *Aws_efs_mount_target) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["file_system_id"] = r.File_system_id
	if r.Ip_address != nil {
		config["ip_address"] = *r.Ip_address
	}
	config["subnet_id"] = r.Subnet_id
	config["network_interface_id"] = r.Network_interface_id
	config["dns_name"] = r.Dns_name
	config["file_system_arn"] = r.File_system_arn
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_macie_member_account_association struct {
	Member_account_id string
}

func Aws_macie_member_account_associationMapper(r *Aws_macie_member_account_association) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["member_account_id"] = r.Member_account_id
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_neptune_parameter_group struct {
	Tags        *map[string]interface{}
	Arn         string
	Name        string
	Family      string
	Description *string
}

func Aws_neptune_parameter_groupMapper(r *Aws_neptune_parameter_group) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["arn"] = r.Arn
	config["name"] = r.Name
	config["family"] = r.Family
	if r.Description != nil {
		config["description"] = *r.Description
	}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_dms_endpoint struct {
	Database_name               *string
	Password                    *string
	Tags                        *map[string]interface{}
	Certificate_arn             *string
	Endpoint_arn                string
	Endpoint_id                 string
	Engine_name                 string
	Ssl_mode                    *string
	Service_access_role         *string
	Endpoint_type               string
	Server_name                 *string
	Extra_connection_attributes *string
	Kms_key_arn                 *string
	Username                    *string
}

func Aws_dms_endpointMapper(r *Aws_dms_endpoint) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Service_access_role != nil {
		config["service_access_role"] = *r.Service_access_role
	}
	config["endpoint_type"] = r.Endpoint_type
	if r.Server_name != nil {
		config["server_name"] = *r.Server_name
	}
	if r.Extra_connection_attributes != nil {
		config["extra_connection_attributes"] = *r.Extra_connection_attributes
	}
	if r.Kms_key_arn != nil {
		config["kms_key_arn"] = *r.Kms_key_arn
	}
	if r.Username != nil {
		config["username"] = *r.Username
	}
	if r.Database_name != nil {
		config["database_name"] = *r.Database_name
	}
	if r.Password != nil {
		config["password"] = *r.Password
	}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	if r.Certificate_arn != nil {
		config["certificate_arn"] = *r.Certificate_arn
	}
	config["endpoint_arn"] = r.Endpoint_arn
	config["endpoint_id"] = r.Endpoint_id
	config["engine_name"] = r.Engine_name
	if r.Ssl_mode != nil {
		config["ssl_mode"] = *r.Ssl_mode
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_ecs_service struct {
	Propagate_tags          *string
	Cluster                 *string
	Task_definition         string
	Enable_ecs_managed_tags *bool
	Scheduling_strategy     *string
	Launch_type             *string
	Name                    string
	Iam_role                *string
	Tags                    *map[string]interface{}
	Platform_version        *string
}

func Aws_ecs_serviceMapper(r *Aws_ecs_service) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	if r.Iam_role != nil {
		config["iam_role"] = *r.Iam_role
	}
	if r.Platform_version != nil {
		config["platform_version"] = *r.Platform_version
	}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	if r.Propagate_tags != nil {
		config["propagate_tags"] = *r.Propagate_tags
	}
	if r.Cluster != nil {
		config["cluster"] = *r.Cluster
	}
	config["task_definition"] = r.Task_definition
	if r.Enable_ecs_managed_tags != nil {
		config["enable_ecs_managed_tags"] = *r.Enable_ecs_managed_tags
	}
	if r.Scheduling_strategy != nil {
		config["scheduling_strategy"] = *r.Scheduling_strategy
	}
	if r.Launch_type != nil {
		config["launch_type"] = *r.Launch_type
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_main_route_table_association struct {
	Vpc_id                  string
	Route_table_id          string
	Original_route_table_id string
}

func Aws_main_route_table_associationMapper(r *Aws_main_route_table_association) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["vpc_id"] = r.Vpc_id
	config["route_table_id"] = r.Route_table_id
	config["original_route_table_id"] = r.Original_route_table_id
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_subnet struct {
	Cidr_block                      string
	Ipv6_cidr_block                 *string
	Availability_zone_id            *string
	Ipv6_cidr_block_association_id  string
	Arn                             string
	Tags                            *map[string]interface{}
	Vpc_id                          string
	Availability_zone               *string
	Map_public_ip_on_launch         *bool
	Assign_ipv6_address_on_creation *bool
	Owner_id                        string
}

func Aws_subnetMapper(r *Aws_subnet) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["cidr_block"] = r.Cidr_block
	if r.Ipv6_cidr_block != nil {
		config["ipv6_cidr_block"] = *r.Ipv6_cidr_block
	}
	if r.Availability_zone_id != nil {
		config["availability_zone_id"] = *r.Availability_zone_id
	}
	config["ipv6_cidr_block_association_id"] = r.Ipv6_cidr_block_association_id
	config["arn"] = r.Arn
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["vpc_id"] = r.Vpc_id
	if r.Availability_zone != nil {
		config["availability_zone"] = *r.Availability_zone
	}
	if r.Map_public_ip_on_launch != nil {
		config["map_public_ip_on_launch"] = *r.Map_public_ip_on_launch
	}
	if r.Assign_ipv6_address_on_creation != nil {
		config["assign_ipv6_address_on_creation"] = *r.Assign_ipv6_address_on_creation
	}
	config["owner_id"] = r.Owner_id
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_vpn_gateway_attachment struct {
	Vpc_id         string
	Vpn_gateway_id string
}

func Aws_vpn_gateway_attachmentMapper(r *Aws_vpn_gateway_attachment) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["vpc_id"] = r.Vpc_id
	config["vpn_gateway_id"] = r.Vpn_gateway_id
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_codepipeline struct {
	Arn      string
	Name     string
	Role_arn string
}

func Aws_codepipelineMapper(r *Aws_codepipeline) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["arn"] = r.Arn
	config["name"] = r.Name
	config["role_arn"] = r.Role_arn
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_wafregional_ipset struct {
	Arn  string
	Name string
}

func Aws_wafregional_ipsetMapper(r *Aws_wafregional_ipset) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["arn"] = r.Arn
	config["name"] = r.Name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_wafregional_regex_pattern_set struct {
	Name string
}

func Aws_wafregional_regex_pattern_setMapper(r *Aws_wafregional_regex_pattern_set) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_opsworks_static_web_layer struct {
	Name                        *string
	Auto_assign_public_ips      *bool
	Drain_elb_on_shutdown       *bool
	Custom_instance_profile_arn *string
	Custom_json                 *string
	Install_updates_on_boot     *bool
	Use_ebs_optimized_instances *bool
	Auto_assign_elastic_ips     *bool
	Auto_healing                *bool
	Stack_id                    string
	Elastic_load_balancer       *string
}

func Aws_opsworks_static_web_layerMapper(r *Aws_opsworks_static_web_layer) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Drain_elb_on_shutdown != nil {
		config["drain_elb_on_shutdown"] = *r.Drain_elb_on_shutdown
	}
	if r.Auto_assign_elastic_ips != nil {
		config["auto_assign_elastic_ips"] = *r.Auto_assign_elastic_ips
	}
	if r.Custom_instance_profile_arn != nil {
		config["custom_instance_profile_arn"] = *r.Custom_instance_profile_arn
	}
	if r.Custom_json != nil {
		config["custom_json"] = *r.Custom_json
	}
	if r.Install_updates_on_boot != nil {
		config["install_updates_on_boot"] = *r.Install_updates_on_boot
	}
	if r.Use_ebs_optimized_instances != nil {
		config["use_ebs_optimized_instances"] = *r.Use_ebs_optimized_instances
	}
	if r.Elastic_load_balancer != nil {
		config["elastic_load_balancer"] = *r.Elastic_load_balancer
	}
	if r.Auto_healing != nil {
		config["auto_healing"] = *r.Auto_healing
	}
	config["stack_id"] = r.Stack_id
	if r.Auto_assign_public_ips != nil {
		config["auto_assign_public_ips"] = *r.Auto_assign_public_ips
	}
	if r.Name != nil {
		config["name"] = *r.Name
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_security_group struct {
	Vpc_id                 *string
	Arn                    string
	Owner_id               string
	Tags                   *map[string]interface{}
	Name                   *string
	Description            *string
	Revoke_rules_on_delete *bool
	Name_prefix            *string
}

func Aws_security_groupMapper(r *Aws_security_group) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Name != nil {
		config["name"] = *r.Name
	}
	if r.Vpc_id != nil {
		config["vpc_id"] = *r.Vpc_id
	}
	config["arn"] = r.Arn
	config["owner_id"] = r.Owner_id
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	if r.Name_prefix != nil {
		config["name_prefix"] = *r.Name_prefix
	}
	if r.Description != nil {
		config["description"] = *r.Description
	}
	if r.Revoke_rules_on_delete != nil {
		config["revoke_rules_on_delete"] = *r.Revoke_rules_on_delete
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_sqs_queue_policy struct {
	Queue_url string
	Policy    string
}

func Aws_sqs_queue_policyMapper(r *Aws_sqs_queue_policy) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["queue_url"] = r.Queue_url
	config["policy"] = r.Policy
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_pinpoint_apns_voip_sandbox_channel struct {
	Team_id                       *string
	Token_key_id                  *string
	Application_id                string
	Certificate                   *string
	Default_authentication_method *string
	Enabled                       *bool
	Private_key                   *string
	Bundle_id                     *string
	Token_key                     *string
}

func Aws_pinpoint_apns_voip_sandbox_channelMapper(r *Aws_pinpoint_apns_voip_sandbox_channel) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Bundle_id != nil {
		config["bundle_id"] = *r.Bundle_id
	}
	if r.Token_key != nil {
		config["token_key"] = *r.Token_key
	}
	if r.Enabled != nil {
		config["enabled"] = *r.Enabled
	}
	if r.Private_key != nil {
		config["private_key"] = *r.Private_key
	}
	if r.Team_id != nil {
		config["team_id"] = *r.Team_id
	}
	if r.Token_key_id != nil {
		config["token_key_id"] = *r.Token_key_id
	}
	config["application_id"] = r.Application_id
	if r.Certificate != nil {
		config["certificate"] = *r.Certificate
	}
	if r.Default_authentication_method != nil {
		config["default_authentication_method"] = *r.Default_authentication_method
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_route struct {
	Destination_prefix_list_id  string
	Transit_gateway_id          *string
	Destination_cidr_block      *string
	Destination_ipv6_cidr_block *string
	Origin                      string
	State                       string
	Route_table_id              string
	Vpc_peering_connection_id   *string
	Nat_gateway_id              *string
	Instance_owner_id           string
	Network_interface_id        *string
	Instance_id                 *string
	Gateway_id                  *string
	Egress_only_gateway_id      *string
}

func Aws_routeMapper(r *Aws_route) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Egress_only_gateway_id != nil {
		config["egress_only_gateway_id"] = *r.Egress_only_gateway_id
	}
	if r.Instance_id != nil {
		config["instance_id"] = *r.Instance_id
	}
	if r.Gateway_id != nil {
		config["gateway_id"] = *r.Gateway_id
	}
	if r.Destination_ipv6_cidr_block != nil {
		config["destination_ipv6_cidr_block"] = *r.Destination_ipv6_cidr_block
	}
	config["destination_prefix_list_id"] = r.Destination_prefix_list_id
	if r.Transit_gateway_id != nil {
		config["transit_gateway_id"] = *r.Transit_gateway_id
	}
	if r.Destination_cidr_block != nil {
		config["destination_cidr_block"] = *r.Destination_cidr_block
	}
	config["instance_owner_id"] = r.Instance_owner_id
	config["origin"] = r.Origin
	config["state"] = r.State
	config["route_table_id"] = r.Route_table_id
	if r.Vpc_peering_connection_id != nil {
		config["vpc_peering_connection_id"] = *r.Vpc_peering_connection_id
	}
	if r.Nat_gateway_id != nil {
		config["nat_gateway_id"] = *r.Nat_gateway_id
	}
	if r.Network_interface_id != nil {
		config["network_interface_id"] = *r.Network_interface_id
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_waf_regex_pattern_set struct {
	Name string
}

func Aws_waf_regex_pattern_setMapper(r *Aws_waf_regex_pattern_set) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_pinpoint_email_channel struct {
	Application_id string
	Enabled        *bool
	From_address   string
	Identity       string
	Role_arn       string
}

func Aws_pinpoint_email_channelMapper(r *Aws_pinpoint_email_channel) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Enabled != nil {
		config["enabled"] = *r.Enabled
	}
	config["from_address"] = r.From_address
	config["identity"] = r.Identity
	config["role_arn"] = r.Role_arn
	config["application_id"] = r.Application_id
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_api_gateway_method struct {
	Api_key_required           *bool
	Request_models             *map[string]interface{}
	Request_parameters         *map[string]interface{}
	Request_parameters_in_json *string
	Request_validator_id       *string
	Http_method                string
	Authorization              string
	Authorizer_id              *string
	Rest_api_id                string
	Resource_id                string
}

func Aws_api_gateway_methodMapper(r *Aws_api_gateway_method) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["resource_id"] = r.Resource_id
	config["authorization"] = r.Authorization
	if r.Authorizer_id != nil {
		config["authorizer_id"] = *r.Authorizer_id
	}
	config["rest_api_id"] = r.Rest_api_id
	if r.Api_key_required != nil {
		config["api_key_required"] = *r.Api_key_required
	}
	if r.Request_models != nil {
		config["request_models"] = *r.Request_models
	}
	if r.Request_parameters != nil {
		config["request_parameters"] = *r.Request_parameters
	}
	if r.Request_parameters_in_json != nil {
		config["request_parameters_in_json"] = *r.Request_parameters_in_json
	}
	if r.Request_validator_id != nil {
		config["request_validator_id"] = *r.Request_validator_id
	}
	config["http_method"] = r.Http_method
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_cognito_identity_pool struct {
	Supported_login_providers        *map[string]interface{}
	Identity_pool_name               string
	Arn                              string
	Developer_provider_name          *string
	Allow_unauthenticated_identities *bool
}

func Aws_cognito_identity_poolMapper(r *Aws_cognito_identity_pool) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["arn"] = r.Arn
	if r.Developer_provider_name != nil {
		config["developer_provider_name"] = *r.Developer_provider_name
	}
	if r.Allow_unauthenticated_identities != nil {
		config["allow_unauthenticated_identities"] = *r.Allow_unauthenticated_identities
	}
	if r.Supported_login_providers != nil {
		config["supported_login_providers"] = *r.Supported_login_providers
	}
	config["identity_pool_name"] = r.Identity_pool_name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_dx_lag struct {
	Tags                  *map[string]interface{}
	Arn                   string
	Name                  string
	Connections_bandwidth string
	Location              string
	Force_destroy         *bool
}

func Aws_dx_lagMapper(r *Aws_dx_lag) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Force_destroy != nil {
		config["force_destroy"] = *r.Force_destroy
	}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["arn"] = r.Arn
	config["name"] = r.Name
	config["connections_bandwidth"] = r.Connections_bandwidth
	config["location"] = r.Location
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_dynamodb_table struct {
	Range_key        *string
	Stream_label     string
	Stream_enabled   *bool
	Tags             *map[string]interface{}
	Arn              string
	Name             string
	Hash_key         string
	Stream_view_type *string
	Billing_mode     *string
	Stream_arn       string
}

func Aws_dynamodb_tableMapper(r *Aws_dynamodb_table) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Range_key != nil {
		config["range_key"] = *r.Range_key
	}
	config["stream_label"] = r.Stream_label
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["arn"] = r.Arn
	config["name"] = r.Name
	if r.Stream_enabled != nil {
		config["stream_enabled"] = *r.Stream_enabled
	}
	config["hash_key"] = r.Hash_key
	if r.Stream_view_type != nil {
		config["stream_view_type"] = *r.Stream_view_type
	}
	if r.Billing_mode != nil {
		config["billing_mode"] = *r.Billing_mode
	}
	config["stream_arn"] = r.Stream_arn
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_iot_thing struct {
	Name              string
	Attributes        *map[string]interface{}
	Thing_type_name   *string
	Default_client_id string
	Arn               string
}

func Aws_iot_thingMapper(r *Aws_iot_thing) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	if r.Attributes != nil {
		config["attributes"] = *r.Attributes
	}
	if r.Thing_type_name != nil {
		config["thing_type_name"] = *r.Thing_type_name
	}
	config["default_client_id"] = r.Default_client_id
	config["arn"] = r.Arn
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_datasync_location_efs struct {
	Tags                *map[string]interface{}
	Uri                 string
	Arn                 string
	Efs_file_system_arn string
	Subdirectory        *string
}

func Aws_datasync_location_efsMapper(r *Aws_datasync_location_efs) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["arn"] = r.Arn
	config["efs_file_system_arn"] = r.Efs_file_system_arn
	if r.Subdirectory != nil {
		config["subdirectory"] = *r.Subdirectory
	}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["uri"] = r.Uri
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_redshift_cluster struct {
	Cluster_type                 *string
	Preferred_maintenance_window *string
	Allow_version_upgrade        *bool
	Owner_account                *string
	Database_name                *string
	Cluster_identifier           string
	Cluster_subnet_group_name    *string
	Cluster_version              *string
	Dns_name                     string
	Bucket_name                  *string
	Cluster_parameter_group_name *string
	Endpoint                     *string
	S3_key_prefix                *string
	Cluster_revision_number      *string
	Master_password              *string
	Snapshot_identifier          *string
	Availability_zone            *string
	Encrypted                    *bool
	Snapshot_cluster_identifier  *string
	Master_username              *string
	Publicly_accessible          *bool
	Final_snapshot_identifier    *string
	Enable_logging               *bool
	Node_type                    string
	Enhanced_vpc_routing         *bool
	Kms_key_id                   *string
	Elastic_ip                   *string
	Skip_final_snapshot          *bool
	Cluster_public_key           *string
	Tags                         *map[string]interface{}
}

func Aws_redshift_clusterMapper(r *Aws_redshift_cluster) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Cluster_subnet_group_name != nil {
		config["cluster_subnet_group_name"] = *r.Cluster_subnet_group_name
	}
	if r.Cluster_version != nil {
		config["cluster_version"] = *r.Cluster_version
	}
	config["dns_name"] = r.Dns_name
	if r.Bucket_name != nil {
		config["bucket_name"] = *r.Bucket_name
	}
	if r.Owner_account != nil {
		config["owner_account"] = *r.Owner_account
	}
	if r.Database_name != nil {
		config["database_name"] = *r.Database_name
	}
	config["cluster_identifier"] = r.Cluster_identifier
	if r.S3_key_prefix != nil {
		config["s3_key_prefix"] = *r.S3_key_prefix
	}
	if r.Cluster_parameter_group_name != nil {
		config["cluster_parameter_group_name"] = *r.Cluster_parameter_group_name
	}
	if r.Endpoint != nil {
		config["endpoint"] = *r.Endpoint
	}
	if r.Cluster_revision_number != nil {
		config["cluster_revision_number"] = *r.Cluster_revision_number
	}
	if r.Snapshot_identifier != nil {
		config["snapshot_identifier"] = *r.Snapshot_identifier
	}
	if r.Master_password != nil {
		config["master_password"] = *r.Master_password
	}
	if r.Availability_zone != nil {
		config["availability_zone"] = *r.Availability_zone
	}
	if r.Encrypted != nil {
		config["encrypted"] = *r.Encrypted
	}
	if r.Snapshot_cluster_identifier != nil {
		config["snapshot_cluster_identifier"] = *r.Snapshot_cluster_identifier
	}
	if r.Final_snapshot_identifier != nil {
		config["final_snapshot_identifier"] = *r.Final_snapshot_identifier
	}
	if r.Master_username != nil {
		config["master_username"] = *r.Master_username
	}
	if r.Publicly_accessible != nil {
		config["publicly_accessible"] = *r.Publicly_accessible
	}
	if r.Elastic_ip != nil {
		config["elastic_ip"] = *r.Elastic_ip
	}
	if r.Skip_final_snapshot != nil {
		config["skip_final_snapshot"] = *r.Skip_final_snapshot
	}
	if r.Cluster_public_key != nil {
		config["cluster_public_key"] = *r.Cluster_public_key
	}
	if r.Enable_logging != nil {
		config["enable_logging"] = *r.Enable_logging
	}
	config["node_type"] = r.Node_type
	if r.Enhanced_vpc_routing != nil {
		config["enhanced_vpc_routing"] = *r.Enhanced_vpc_routing
	}
	if r.Kms_key_id != nil {
		config["kms_key_id"] = *r.Kms_key_id
	}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	if r.Cluster_type != nil {
		config["cluster_type"] = *r.Cluster_type
	}
	if r.Preferred_maintenance_window != nil {
		config["preferred_maintenance_window"] = *r.Preferred_maintenance_window
	}
	if r.Allow_version_upgrade != nil {
		config["allow_version_upgrade"] = *r.Allow_version_upgrade
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_api_gateway_resource struct {
	Rest_api_id string
	Parent_id   string
	Path_part   string
	Path        string
}

func Aws_api_gateway_resourceMapper(r *Aws_api_gateway_resource) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["rest_api_id"] = r.Rest_api_id
	config["parent_id"] = r.Parent_id
	config["path_part"] = r.Path_part
	config["path"] = r.Path
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_load_balancer_policy struct {
	Load_balancer_name string
	Policy_name        string
	Policy_type_name   string
}

func Aws_load_balancer_policyMapper(r *Aws_load_balancer_policy) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["load_balancer_name"] = r.Load_balancer_name
	config["policy_name"] = r.Policy_name
	config["policy_type_name"] = r.Policy_type_name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_ses_active_receipt_rule_set struct {
	Rule_set_name string
}

func Aws_ses_active_receipt_rule_setMapper(r *Aws_ses_active_receipt_rule_set) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["rule_set_name"] = r.Rule_set_name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_pinpoint_adm_channel struct {
	Client_id      string
	Client_secret  string
	Enabled        *bool
	Application_id string
}

func Aws_pinpoint_adm_channelMapper(r *Aws_pinpoint_adm_channel) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["application_id"] = r.Application_id
	config["client_id"] = r.Client_id
	config["client_secret"] = r.Client_secret
	if r.Enabled != nil {
		config["enabled"] = *r.Enabled
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_ssm_maintenance_window_task struct {
	Max_concurrency  string
	Task_type        string
	Task_arn         string
	Description      *string
	Window_id        string
	Max_errors       string
	Service_role_arn string
	Name             *string
}

func Aws_ssm_maintenance_window_taskMapper(r *Aws_ssm_maintenance_window_task) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["window_id"] = r.Window_id
	config["max_errors"] = r.Max_errors
	config["service_role_arn"] = r.Service_role_arn
	if r.Name != nil {
		config["name"] = *r.Name
	}
	config["max_concurrency"] = r.Max_concurrency
	config["task_type"] = r.Task_type
	config["task_arn"] = r.Task_arn
	if r.Description != nil {
		config["description"] = *r.Description
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_storagegateway_smb_file_share struct {
	Fileshare_id            string
	Gateway_arn             string
	Location_arn            string
	Requester_pays          *bool
	Arn                     string
	Read_only               *bool
	Default_storage_class   *string
	Kms_key_arn             *string
	Object_acl              *string
	Role_arn                string
	Authentication          *string
	Kms_encrypted           *bool
	Guess_mime_type_enabled *bool
}

func Aws_storagegateway_smb_file_shareMapper(r *Aws_storagegateway_smb_file_share) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Authentication != nil {
		config["authentication"] = *r.Authentication
	}
	if r.Default_storage_class != nil {
		config["default_storage_class"] = *r.Default_storage_class
	}
	if r.Kms_key_arn != nil {
		config["kms_key_arn"] = *r.Kms_key_arn
	}
	if r.Object_acl != nil {
		config["object_acl"] = *r.Object_acl
	}
	config["role_arn"] = r.Role_arn
	if r.Guess_mime_type_enabled != nil {
		config["guess_mime_type_enabled"] = *r.Guess_mime_type_enabled
	}
	if r.Kms_encrypted != nil {
		config["kms_encrypted"] = *r.Kms_encrypted
	}
	config["arn"] = r.Arn
	config["fileshare_id"] = r.Fileshare_id
	config["gateway_arn"] = r.Gateway_arn
	config["location_arn"] = r.Location_arn
	if r.Requester_pays != nil {
		config["requester_pays"] = *r.Requester_pays
	}
	if r.Read_only != nil {
		config["read_only"] = *r.Read_only
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_elb_attachment struct {
	Elb      string
	Instance string
}

func Aws_elb_attachmentMapper(r *Aws_elb_attachment) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["elb"] = r.Elb
	config["instance"] = r.Instance
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_opsworks_rails_app_layer struct {
	Auto_assign_public_ips      *bool
	Install_updates_on_boot     *bool
	Stack_id                    string
	Rubygems_version            *string
	Custom_instance_profile_arn *string
	Drain_elb_on_shutdown       *bool
	Use_ebs_optimized_instances *bool
	App_server                  *string
	Manage_bundler              *bool
	Custom_json                 *string
	Name                        *string
	Passenger_version           *string
	Auto_assign_elastic_ips     *bool
	Elastic_load_balancer       *string
	Auto_healing                *bool
	Bundler_version             *string
	Ruby_version                *string
}

func Aws_opsworks_rails_app_layerMapper(r *Aws_opsworks_rails_app_layer) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Bundler_version != nil {
		config["bundler_version"] = *r.Bundler_version
	}
	if r.Ruby_version != nil {
		config["ruby_version"] = *r.Ruby_version
	}
	if r.Auto_assign_elastic_ips != nil {
		config["auto_assign_elastic_ips"] = *r.Auto_assign_elastic_ips
	}
	if r.Elastic_load_balancer != nil {
		config["elastic_load_balancer"] = *r.Elastic_load_balancer
	}
	if r.Auto_healing != nil {
		config["auto_healing"] = *r.Auto_healing
	}
	if r.Install_updates_on_boot != nil {
		config["install_updates_on_boot"] = *r.Install_updates_on_boot
	}
	config["stack_id"] = r.Stack_id
	if r.Rubygems_version != nil {
		config["rubygems_version"] = *r.Rubygems_version
	}
	if r.Auto_assign_public_ips != nil {
		config["auto_assign_public_ips"] = *r.Auto_assign_public_ips
	}
	if r.App_server != nil {
		config["app_server"] = *r.App_server
	}
	if r.Manage_bundler != nil {
		config["manage_bundler"] = *r.Manage_bundler
	}
	if r.Custom_instance_profile_arn != nil {
		config["custom_instance_profile_arn"] = *r.Custom_instance_profile_arn
	}
	if r.Drain_elb_on_shutdown != nil {
		config["drain_elb_on_shutdown"] = *r.Drain_elb_on_shutdown
	}
	if r.Use_ebs_optimized_instances != nil {
		config["use_ebs_optimized_instances"] = *r.Use_ebs_optimized_instances
	}
	if r.Name != nil {
		config["name"] = *r.Name
	}
	if r.Passenger_version != nil {
		config["passenger_version"] = *r.Passenger_version
	}
	if r.Custom_json != nil {
		config["custom_json"] = *r.Custom_json
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_organizations_policy struct {
	Resource_type *string
	Arn           string
	Content       string
	Description   *string
	Name          string
}

func Aws_organizations_policyMapper(r *Aws_organizations_policy) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["arn"] = r.Arn
	config["content"] = r.Content
	if r.Description != nil {
		config["description"] = *r.Description
	}
	config["name"] = r.Name
	if r.Resource_type != nil {
		config["resource_type"] = *r.Resource_type
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_ses_domain_mail_from struct {
	Domain                 string
	Mail_from_domain       string
	Behavior_on_mx_failure *string
}

func Aws_ses_domain_mail_fromMapper(r *Aws_ses_domain_mail_from) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Behavior_on_mx_failure != nil {
		config["behavior_on_mx_failure"] = *r.Behavior_on_mx_failure
	}
	config["domain"] = r.Domain
	config["mail_from_domain"] = r.Mail_from_domain
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_ses_receipt_rule_set struct {
	Rule_set_name string
}

func Aws_ses_receipt_rule_setMapper(r *Aws_ses_receipt_rule_set) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["rule_set_name"] = r.Rule_set_name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_config_configuration_recorder_status struct {
	Name       string
	Is_enabled bool
}

func Aws_config_configuration_recorder_statusMapper(r *Aws_config_configuration_recorder_status) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	config["is_enabled"] = r.Is_enabled
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_db_snapshot struct {
	Db_instance_identifier        string
	Availability_zone             string
	Engine_version                string
	Storage_type                  string
	Kms_key_id                    string
	License_model                 string
	Snapshot_type                 string
	Status                        string
	Db_snapshot_identifier        string
	Db_snapshot_arn               string
	Engine                        string
	Source_db_snapshot_identifier string
	Source_region                 string
	Tags                          *map[string]interface{}
	Encrypted                     bool
	Option_group_name             string
	Vpc_id                        string
}

func Aws_db_snapshotMapper(r *Aws_db_snapshot) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["source_region"] = r.Source_region
	config["snapshot_type"] = r.Snapshot_type
	config["status"] = r.Status
	config["db_snapshot_identifier"] = r.Db_snapshot_identifier
	config["db_snapshot_arn"] = r.Db_snapshot_arn
	config["engine"] = r.Engine
	config["source_db_snapshot_identifier"] = r.Source_db_snapshot_identifier
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["encrypted"] = r.Encrypted
	config["option_group_name"] = r.Option_group_name
	config["vpc_id"] = r.Vpc_id
	config["storage_type"] = r.Storage_type
	config["db_instance_identifier"] = r.Db_instance_identifier
	config["availability_zone"] = r.Availability_zone
	config["engine_version"] = r.Engine_version
	config["kms_key_id"] = r.Kms_key_id
	config["license_model"] = r.License_model
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_dax_cluster struct {
	Cluster_name           string
	Subnet_group_name      *string
	Cluster_address        string
	Node_type              string
	Notification_topic_arn *string
	Configuration_endpoint string
	Iam_role_arn           string
	Description            *string
	Maintenance_window     *string
	Arn                    string
	Parameter_group_name   *string
	Tags                   *map[string]interface{}
}

func Aws_dax_clusterMapper(r *Aws_dax_cluster) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["cluster_name"] = r.Cluster_name
	if r.Subnet_group_name != nil {
		config["subnet_group_name"] = *r.Subnet_group_name
	}
	config["configuration_endpoint"] = r.Configuration_endpoint
	config["cluster_address"] = r.Cluster_address
	config["node_type"] = r.Node_type
	if r.Notification_topic_arn != nil {
		config["notification_topic_arn"] = *r.Notification_topic_arn
	}
	config["iam_role_arn"] = r.Iam_role_arn
	if r.Description != nil {
		config["description"] = *r.Description
	}
	if r.Maintenance_window != nil {
		config["maintenance_window"] = *r.Maintenance_window
	}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["arn"] = r.Arn
	if r.Parameter_group_name != nil {
		config["parameter_group_name"] = *r.Parameter_group_name
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_waf_regex_match_set struct {
	Name string
}

func Aws_waf_regex_match_setMapper(r *Aws_waf_regex_match_set) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_dms_replication_instance struct {
	Availability_zone            *string
	Replication_instance_class   string
	Engine_version               *string
	Multi_az                     *bool
	Publicly_accessible          *bool
	Replication_instance_id      string
	Auto_minor_version_upgrade   *bool
	Replication_instance_arn     string
	Tags                         *map[string]interface{}
	Apply_immediately            *bool
	Kms_key_arn                  *string
	Preferred_maintenance_window *string
	Replication_subnet_group_id  *string
}

func Aws_dms_replication_instanceMapper(r *Aws_dms_replication_instance) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Preferred_maintenance_window != nil {
		config["preferred_maintenance_window"] = *r.Preferred_maintenance_window
	}
	if r.Replication_subnet_group_id != nil {
		config["replication_subnet_group_id"] = *r.Replication_subnet_group_id
	}
	if r.Apply_immediately != nil {
		config["apply_immediately"] = *r.Apply_immediately
	}
	if r.Kms_key_arn != nil {
		config["kms_key_arn"] = *r.Kms_key_arn
	}
	if r.Availability_zone != nil {
		config["availability_zone"] = *r.Availability_zone
	}
	config["replication_instance_class"] = r.Replication_instance_class
	config["replication_instance_id"] = r.Replication_instance_id
	if r.Engine_version != nil {
		config["engine_version"] = *r.Engine_version
	}
	if r.Multi_az != nil {
		config["multi_az"] = *r.Multi_az
	}
	if r.Publicly_accessible != nil {
		config["publicly_accessible"] = *r.Publicly_accessible
	}
	if r.Auto_minor_version_upgrade != nil {
		config["auto_minor_version_upgrade"] = *r.Auto_minor_version_upgrade
	}
	config["replication_instance_arn"] = r.Replication_instance_arn
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_emr_cluster struct {
	Additional_info                   *string
	Keep_job_flow_alive_when_no_steps *bool
	Tags                              *map[string]interface{}
	Scale_down_behavior               *string
	Name                              string
	Configurations_json               *string
	Release_label                     string
	Master_public_dns                 string
	Master_instance_type              *string
	Log_uri                           *string
	Security_configuration            *string
	Core_instance_type                *string
	Configurations                    *string
	Visible_to_all_users              *bool
	Custom_ami_id                     *string
	Autoscaling_role                  *string
	Cluster_state                     string
	Termination_protection            *bool
	Service_role                      string
}

func Aws_emr_clusterMapper(r *Aws_emr_cluster) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Core_instance_type != nil {
		config["core_instance_type"] = *r.Core_instance_type
	}
	if r.Configurations != nil {
		config["configurations"] = *r.Configurations
	}
	if r.Visible_to_all_users != nil {
		config["visible_to_all_users"] = *r.Visible_to_all_users
	}
	if r.Custom_ami_id != nil {
		config["custom_ami_id"] = *r.Custom_ami_id
	}
	if r.Autoscaling_role != nil {
		config["autoscaling_role"] = *r.Autoscaling_role
	}
	config["cluster_state"] = r.Cluster_state
	if r.Termination_protection != nil {
		config["termination_protection"] = *r.Termination_protection
	}
	config["service_role"] = r.Service_role
	if r.Additional_info != nil {
		config["additional_info"] = *r.Additional_info
	}
	if r.Keep_job_flow_alive_when_no_steps != nil {
		config["keep_job_flow_alive_when_no_steps"] = *r.Keep_job_flow_alive_when_no_steps
	}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	if r.Scale_down_behavior != nil {
		config["scale_down_behavior"] = *r.Scale_down_behavior
	}
	config["name"] = r.Name
	if r.Configurations_json != nil {
		config["configurations_json"] = *r.Configurations_json
	}
	config["release_label"] = r.Release_label
	config["master_public_dns"] = r.Master_public_dns
	if r.Master_instance_type != nil {
		config["master_instance_type"] = *r.Master_instance_type
	}
	if r.Log_uri != nil {
		config["log_uri"] = *r.Log_uri
	}
	if r.Security_configuration != nil {
		config["security_configuration"] = *r.Security_configuration
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_internet_gateway struct {
	Vpc_id   *string
	Tags     *map[string]interface{}
	Owner_id string
}

func Aws_internet_gatewayMapper(r *Aws_internet_gateway) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Vpc_id != nil {
		config["vpc_id"] = *r.Vpc_id
	}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["owner_id"] = r.Owner_id
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_ami struct {
	Description          *string
	Tags                 *map[string]interface{}
	Image_location       *string
	Ena_support          *bool
	Manage_ebs_snapshots bool
	Root_device_name     *string
	Root_snapshot_id     string
	Architecture         *string
	Name                 string
	Ramdisk_id           *string
	Kernel_id            *string
	Sriov_net_support    *string
	Virtualization_type  *string
}

func Aws_amiMapper(r *Aws_ami) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Virtualization_type != nil {
		config["virtualization_type"] = *r.Virtualization_type
	}
	if r.Kernel_id != nil {
		config["kernel_id"] = *r.Kernel_id
	}
	if r.Sriov_net_support != nil {
		config["sriov_net_support"] = *r.Sriov_net_support
	}
	if r.Description != nil {
		config["description"] = *r.Description
	}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	if r.Root_device_name != nil {
		config["root_device_name"] = *r.Root_device_name
	}
	if r.Image_location != nil {
		config["image_location"] = *r.Image_location
	}
	if r.Ena_support != nil {
		config["ena_support"] = *r.Ena_support
	}
	config["manage_ebs_snapshots"] = r.Manage_ebs_snapshots
	if r.Ramdisk_id != nil {
		config["ramdisk_id"] = *r.Ramdisk_id
	}
	config["root_snapshot_id"] = r.Root_snapshot_id
	if r.Architecture != nil {
		config["architecture"] = *r.Architecture
	}
	config["name"] = r.Name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_api_gateway_authorizer struct {
	Name                           string
	Identity_validation_expression *string
	Rest_api_id                    string
	Resource_type                  *string
	Authorizer_credentials         *string
	Authorizer_uri                 *string
	Identity_source                *string
}

func Aws_api_gateway_authorizerMapper(r *Aws_api_gateway_authorizer) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	if r.Identity_validation_expression != nil {
		config["identity_validation_expression"] = *r.Identity_validation_expression
	}
	if r.Authorizer_uri != nil {
		config["authorizer_uri"] = *r.Authorizer_uri
	}
	if r.Identity_source != nil {
		config["identity_source"] = *r.Identity_source
	}
	config["rest_api_id"] = r.Rest_api_id
	if r.Resource_type != nil {
		config["resource_type"] = *r.Resource_type
	}
	if r.Authorizer_credentials != nil {
		config["authorizer_credentials"] = *r.Authorizer_credentials
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_api_gateway_usage_plan struct {
	Product_code *string
	Name         string
	Description  *string
}

func Aws_api_gateway_usage_planMapper(r *Aws_api_gateway_usage_plan) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	if r.Description != nil {
		config["description"] = *r.Description
	}
	if r.Product_code != nil {
		config["product_code"] = *r.Product_code
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_cloudwatch_log_destination struct {
	Target_arn string
	Arn        string
	Name       string
	Role_arn   string
}

func Aws_cloudwatch_log_destinationMapper(r *Aws_cloudwatch_log_destination) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["role_arn"] = r.Role_arn
	config["target_arn"] = r.Target_arn
	config["arn"] = r.Arn
	config["name"] = r.Name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_codedeploy_app struct {
	Name             string
	Compute_platform *string
	Unique_id        *string
}

func Aws_codedeploy_appMapper(r *Aws_codedeploy_app) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	if r.Compute_platform != nil {
		config["compute_platform"] = *r.Compute_platform
	}
	if r.Unique_id != nil {
		config["unique_id"] = *r.Unique_id
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_key_pair struct {
	Key_name_prefix *string
	Public_key      string
	Fingerprint     string
	Key_name        *string
}

func Aws_key_pairMapper(r *Aws_key_pair) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Key_name != nil {
		config["key_name"] = *r.Key_name
	}
	if r.Key_name_prefix != nil {
		config["key_name_prefix"] = *r.Key_name_prefix
	}
	config["public_key"] = r.Public_key
	config["fingerprint"] = r.Fingerprint
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_route_table_association struct {
	Route_table_id string
	Subnet_id      string
}

func Aws_route_table_associationMapper(r *Aws_route_table_association) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["route_table_id"] = r.Route_table_id
	config["subnet_id"] = r.Subnet_id
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_api_gateway_api_key struct {
	Description       *string
	Enabled           *bool
	Created_date      string
	Last_updated_date string
	Value             *string
	Name              string
}

func Aws_api_gateway_api_keyMapper(r *Aws_api_gateway_api_key) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Description != nil {
		config["description"] = *r.Description
	}
	if r.Enabled != nil {
		config["enabled"] = *r.Enabled
	}
	config["created_date"] = r.Created_date
	config["last_updated_date"] = r.Last_updated_date
	if r.Value != nil {
		config["value"] = *r.Value
	}
	config["name"] = r.Name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_api_gateway_usage_plan_key struct {
	Key_type      string
	Usage_plan_id string
	Name          string
	Value         string
	Key_id        string
}

func Aws_api_gateway_usage_plan_keyMapper(r *Aws_api_gateway_usage_plan_key) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["key_id"] = r.Key_id
	config["key_type"] = r.Key_type
	config["usage_plan_id"] = r.Usage_plan_id
	config["name"] = r.Name
	config["value"] = r.Value
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_appautoscaling_target struct {
	Service_namespace  string
	Resource_id        string
	Role_arn           *string
	Scalable_dimension string
}

func Aws_appautoscaling_targetMapper(r *Aws_appautoscaling_target) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["resource_id"] = r.Resource_id
	if r.Role_arn != nil {
		config["role_arn"] = *r.Role_arn
	}
	config["scalable_dimension"] = r.Scalable_dimension
	config["service_namespace"] = r.Service_namespace
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_api_gateway_request_validator struct {
	Rest_api_id                 string
	Name                        string
	Validate_request_body       *bool
	Validate_request_parameters *bool
}

func Aws_api_gateway_request_validatorMapper(r *Aws_api_gateway_request_validator) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["rest_api_id"] = r.Rest_api_id
	config["name"] = r.Name
	if r.Validate_request_body != nil {
		config["validate_request_body"] = *r.Validate_request_body
	}
	if r.Validate_request_parameters != nil {
		config["validate_request_parameters"] = *r.Validate_request_parameters
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_service_discovery_private_dns_namespace struct {
	Name        string
	Description *string
	Vpc         string
	Arn         string
	Hosted_zone string
}

func Aws_service_discovery_private_dns_namespaceMapper(r *Aws_service_discovery_private_dns_namespace) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	if r.Description != nil {
		config["description"] = *r.Description
	}
	config["vpc"] = r.Vpc
	config["arn"] = r.Arn
	config["hosted_zone"] = r.Hosted_zone
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_docdb_cluster_parameter_group struct {
	Arn         string
	Name        *string
	Name_prefix *string
	Family      string
	Description *string
	Tags        *map[string]interface{}
}

func Aws_docdb_cluster_parameter_groupMapper(r *Aws_docdb_cluster_parameter_group) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["arn"] = r.Arn
	if r.Name != nil {
		config["name"] = *r.Name
	}
	if r.Name_prefix != nil {
		config["name_prefix"] = *r.Name_prefix
	}
	config["family"] = r.Family
	if r.Description != nil {
		config["description"] = *r.Description
	}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_route53_record struct {
	Multivalue_answer_routing_policy *bool
	Fqdn                             string
	Resource_type                    string
	Set_identifier                   *string
	Failover                         *string
	Name                             string
	Zone_id                          string
	Allow_overwrite                  *bool
	Health_check_id                  *string
}

func Aws_route53_recordMapper(r *Aws_route53_record) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Multivalue_answer_routing_policy != nil {
		config["multivalue_answer_routing_policy"] = *r.Multivalue_answer_routing_policy
	}
	config["fqdn"] = r.Fqdn
	config["resource_type"] = r.Resource_type
	if r.Set_identifier != nil {
		config["set_identifier"] = *r.Set_identifier
	}
	if r.Failover != nil {
		config["failover"] = *r.Failover
	}
	config["name"] = r.Name
	config["zone_id"] = r.Zone_id
	if r.Allow_overwrite != nil {
		config["allow_overwrite"] = *r.Allow_overwrite
	}
	if r.Health_check_id != nil {
		config["health_check_id"] = *r.Health_check_id
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_batch_job_definition struct {
	Name                 string
	Container_properties *string
	Parameters           *map[string]interface{}
	Resource_type        string
	Arn                  string
}

func Aws_batch_job_definitionMapper(r *Aws_batch_job_definition) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["resource_type"] = r.Resource_type
	config["arn"] = r.Arn
	config["name"] = r.Name
	if r.Container_properties != nil {
		config["container_properties"] = *r.Container_properties
	}
	if r.Parameters != nil {
		config["parameters"] = *r.Parameters
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_batch_job_queue struct {
	Name  string
	State string
	Arn   string
}

func Aws_batch_job_queueMapper(r *Aws_batch_job_queue) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	config["state"] = r.State
	config["arn"] = r.Arn
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_glue_crawler struct {
	Table_prefix           *string
	Name                   string
	Database_name          string
	Role                   string
	Description            *string
	Schedule               *string
	Security_configuration *string
	Configuration          *string
}

func Aws_glue_crawlerMapper(r *Aws_glue_crawler) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Table_prefix != nil {
		config["table_prefix"] = *r.Table_prefix
	}
	config["name"] = r.Name
	config["database_name"] = r.Database_name
	config["role"] = r.Role
	if r.Description != nil {
		config["description"] = *r.Description
	}
	if r.Schedule != nil {
		config["schedule"] = *r.Schedule
	}
	if r.Security_configuration != nil {
		config["security_configuration"] = *r.Security_configuration
	}
	if r.Configuration != nil {
		config["configuration"] = *r.Configuration
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_iam_service_linked_role struct {
	Create_date      string
	Unique_id        string
	Custom_suffix    *string
	Description      *string
	Aws_service_name string
	Name             string
	Path             string
	Arn              string
}

func Aws_iam_service_linked_roleMapper(r *Aws_iam_service_linked_role) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["unique_id"] = r.Unique_id
	if r.Custom_suffix != nil {
		config["custom_suffix"] = *r.Custom_suffix
	}
	if r.Description != nil {
		config["description"] = *r.Description
	}
	config["aws_service_name"] = r.Aws_service_name
	config["name"] = r.Name
	config["path"] = r.Path
	config["arn"] = r.Arn
	config["create_date"] = r.Create_date
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_lightsail_key_pair struct {
	Name_prefix           *string
	Pgp_key               *string
	Arn                   string
	Fingerprint           string
	Public_key            *string
	Private_key           string
	Encrypted_fingerprint string
	Encrypted_private_key string
	Name                  *string
}

func Aws_lightsail_key_pairMapper(r *Aws_lightsail_key_pair) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["fingerprint"] = r.Fingerprint
	if r.Public_key != nil {
		config["public_key"] = *r.Public_key
	}
	config["private_key"] = r.Private_key
	config["encrypted_fingerprint"] = r.Encrypted_fingerprint
	config["encrypted_private_key"] = r.Encrypted_private_key
	if r.Name != nil {
		config["name"] = *r.Name
	}
	if r.Pgp_key != nil {
		config["pgp_key"] = *r.Pgp_key
	}
	config["arn"] = r.Arn
	if r.Name_prefix != nil {
		config["name_prefix"] = *r.Name_prefix
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_pinpoint_baidu_channel struct {
	Application_id string
	Enabled        *bool
	Api_key        string
	Secret_key     string
}

func Aws_pinpoint_baidu_channelMapper(r *Aws_pinpoint_baidu_channel) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["secret_key"] = r.Secret_key
	config["application_id"] = r.Application_id
	if r.Enabled != nil {
		config["enabled"] = *r.Enabled
	}
	config["api_key"] = r.Api_key
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_lb struct {
	Enable_http2                     *bool
	Dns_name                         string
	Tags                             *map[string]interface{}
	Ip_address_type                  *string
	Vpc_id                           string
	Name                             *string
	Internal                         *bool
	Enable_cross_zone_load_balancing *bool
	Arn                              string
	Arn_suffix                       string
	Load_balancer_type               *string
	Name_prefix                      *string
	Enable_deletion_protection       *bool
	Zone_id                          string
}

func Aws_lbMapper(r *Aws_lb) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["arn"] = r.Arn
	config["arn_suffix"] = r.Arn_suffix
	if r.Load_balancer_type != nil {
		config["load_balancer_type"] = *r.Load_balancer_type
	}
	if r.Name_prefix != nil {
		config["name_prefix"] = *r.Name_prefix
	}
	if r.Enable_deletion_protection != nil {
		config["enable_deletion_protection"] = *r.Enable_deletion_protection
	}
	config["zone_id"] = r.Zone_id
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	if r.Enable_http2 != nil {
		config["enable_http2"] = *r.Enable_http2
	}
	config["dns_name"] = r.Dns_name
	if r.Enable_cross_zone_load_balancing != nil {
		config["enable_cross_zone_load_balancing"] = *r.Enable_cross_zone_load_balancing
	}
	if r.Ip_address_type != nil {
		config["ip_address_type"] = *r.Ip_address_type
	}
	config["vpc_id"] = r.Vpc_id
	if r.Name != nil {
		config["name"] = *r.Name
	}
	if r.Internal != nil {
		config["internal"] = *r.Internal
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_api_gateway_base_path_mapping struct {
	Base_path   *string
	Stage_name  *string
	Domain_name string
	Api_id      string
}

func Aws_api_gateway_base_path_mappingMapper(r *Aws_api_gateway_base_path_mapping) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Stage_name != nil {
		config["stage_name"] = *r.Stage_name
	}
	config["domain_name"] = r.Domain_name
	config["api_id"] = r.Api_id
	if r.Base_path != nil {
		config["base_path"] = *r.Base_path
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_cloudtrail struct {
	Tags                          *map[string]interface{}
	S3_key_prefix                 *string
	Cloud_watch_logs_group_arn    *string
	Arn                           string
	Sns_topic_name                *string
	Name                          string
	Enable_logging                *bool
	Is_multi_region_trail         *bool
	Enable_log_file_validation    *bool
	Kms_key_id                    *string
	Home_region                   string
	S3_bucket_name                string
	Cloud_watch_logs_role_arn     *string
	Is_organization_trail         *bool
	Include_global_service_events *bool
}

func Aws_cloudtrailMapper(r *Aws_cloudtrail) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.S3_key_prefix != nil {
		config["s3_key_prefix"] = *r.S3_key_prefix
	}
	if r.Cloud_watch_logs_group_arn != nil {
		config["cloud_watch_logs_group_arn"] = *r.Cloud_watch_logs_group_arn
	}
	config["arn"] = r.Arn
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["name"] = r.Name
	if r.Enable_logging != nil {
		config["enable_logging"] = *r.Enable_logging
	}
	if r.Is_multi_region_trail != nil {
		config["is_multi_region_trail"] = *r.Is_multi_region_trail
	}
	if r.Sns_topic_name != nil {
		config["sns_topic_name"] = *r.Sns_topic_name
	}
	config["s3_bucket_name"] = r.S3_bucket_name
	if r.Cloud_watch_logs_role_arn != nil {
		config["cloud_watch_logs_role_arn"] = *r.Cloud_watch_logs_role_arn
	}
	if r.Is_organization_trail != nil {
		config["is_organization_trail"] = *r.Is_organization_trail
	}
	if r.Enable_log_file_validation != nil {
		config["enable_log_file_validation"] = *r.Enable_log_file_validation
	}
	if r.Kms_key_id != nil {
		config["kms_key_id"] = *r.Kms_key_id
	}
	config["home_region"] = r.Home_region
	if r.Include_global_service_events != nil {
		config["include_global_service_events"] = *r.Include_global_service_events
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_neptune_cluster_snapshot struct {
	Snapshot_type                  string
	Status                         string
	Vpc_id                         string
	Db_cluster_identifier          string
	License_model                  string
	Db_cluster_snapshot_identifier string
	Db_cluster_snapshot_arn        string
	Engine                         string
	Engine_version                 string
	Kms_key_id                     string
	Source_db_cluster_snapshot_arn string
	Storage_encrypted              bool
}

func Aws_neptune_cluster_snapshotMapper(r *Aws_neptune_cluster_snapshot) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["vpc_id"] = r.Vpc_id
	config["db_cluster_identifier"] = r.Db_cluster_identifier
	config["status"] = r.Status
	config["license_model"] = r.License_model
	config["db_cluster_snapshot_identifier"] = r.Db_cluster_snapshot_identifier
	config["engine"] = r.Engine
	config["engine_version"] = r.Engine_version
	config["kms_key_id"] = r.Kms_key_id
	config["source_db_cluster_snapshot_arn"] = r.Source_db_cluster_snapshot_arn
	config["storage_encrypted"] = r.Storage_encrypted
	config["db_cluster_snapshot_arn"] = r.Db_cluster_snapshot_arn
	config["snapshot_type"] = r.Snapshot_type
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_rds_cluster_endpoint struct {
	Endpoint                    string
	Arn                         string
	Cluster_endpoint_identifier string
	Cluster_identifier          string
	Custom_endpoint_type        string
}

func Aws_rds_cluster_endpointMapper(r *Aws_rds_cluster_endpoint) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["arn"] = r.Arn
	config["cluster_endpoint_identifier"] = r.Cluster_endpoint_identifier
	config["cluster_identifier"] = r.Cluster_identifier
	config["custom_endpoint_type"] = r.Custom_endpoint_type
	config["endpoint"] = r.Endpoint
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_ses_domain_identity struct {
	Domain             string
	Verification_token string
	Arn                string
}

func Aws_ses_domain_identityMapper(r *Aws_ses_domain_identity) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["domain"] = r.Domain
	config["verification_token"] = r.Verification_token
	config["arn"] = r.Arn
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_opsworks_stack struct {
	Use_custom_cookbooks          *bool
	Name                          string
	Configuration_manager_version *string
	Default_availability_zone     *string
	Default_root_device_type      *string
	Default_subnet_id             *string
	Manage_berkshelf              *bool
	Berkshelf_version             *string
	Stack_endpoint                string
	Tags                          *map[string]interface{}
	Use_opsworks_security_groups  *bool
	Service_role_arn              string
	Custom_json                   *string
	Default_ssh_key_name          *string
	Hostname_theme                *string
	Configuration_manager_name    *string
	Default_os                    *string
	Vpc_id                        *string
	Agent_version                 *string
	Arn                           string
	Region                        string
	Default_instance_profile_arn  string
	Color                         *string
}

func Aws_opsworks_stackMapper(r *Aws_opsworks_stack) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["arn"] = r.Arn
	config["region"] = r.Region
	config["default_instance_profile_arn"] = r.Default_instance_profile_arn
	if r.Color != nil {
		config["color"] = *r.Color
	}
	if r.Configuration_manager_name != nil {
		config["configuration_manager_name"] = *r.Configuration_manager_name
	}
	if r.Default_os != nil {
		config["default_os"] = *r.Default_os
	}
	if r.Vpc_id != nil {
		config["vpc_id"] = *r.Vpc_id
	}
	if r.Agent_version != nil {
		config["agent_version"] = *r.Agent_version
	}
	if r.Configuration_manager_version != nil {
		config["configuration_manager_version"] = *r.Configuration_manager_version
	}
	if r.Default_availability_zone != nil {
		config["default_availability_zone"] = *r.Default_availability_zone
	}
	if r.Default_root_device_type != nil {
		config["default_root_device_type"] = *r.Default_root_device_type
	}
	if r.Default_subnet_id != nil {
		config["default_subnet_id"] = *r.Default_subnet_id
	}
	if r.Use_custom_cookbooks != nil {
		config["use_custom_cookbooks"] = *r.Use_custom_cookbooks
	}
	config["name"] = r.Name
	if r.Berkshelf_version != nil {
		config["berkshelf_version"] = *r.Berkshelf_version
	}
	config["stack_endpoint"] = r.Stack_endpoint
	if r.Manage_berkshelf != nil {
		config["manage_berkshelf"] = *r.Manage_berkshelf
	}
	if r.Custom_json != nil {
		config["custom_json"] = *r.Custom_json
	}
	if r.Default_ssh_key_name != nil {
		config["default_ssh_key_name"] = *r.Default_ssh_key_name
	}
	if r.Hostname_theme != nil {
		config["hostname_theme"] = *r.Hostname_theme
	}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	if r.Use_opsworks_security_groups != nil {
		config["use_opsworks_security_groups"] = *r.Use_opsworks_security_groups
	}
	config["service_role_arn"] = r.Service_role_arn
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_vpn_connection_route struct {
	Destination_cidr_block string
	Vpn_connection_id      string
}

func Aws_vpn_connection_routeMapper(r *Aws_vpn_connection_route) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["destination_cidr_block"] = r.Destination_cidr_block
	config["vpn_connection_id"] = r.Vpn_connection_id
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_service_discovery_public_dns_namespace struct {
	Name        string
	Description *string
	Arn         string
	Hosted_zone string
}

func Aws_service_discovery_public_dns_namespaceMapper(r *Aws_service_discovery_public_dns_namespace) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["arn"] = r.Arn
	config["hosted_zone"] = r.Hosted_zone
	config["name"] = r.Name
	if r.Description != nil {
		config["description"] = *r.Description
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_waf_rate_based_rule struct {
	Name        string
	Metric_name string
	Rate_key    string
}

func Aws_waf_rate_based_ruleMapper(r *Aws_waf_rate_based_rule) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	config["metric_name"] = r.Metric_name
	config["rate_key"] = r.Rate_key
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_ec2_transit_gateway_route_table_association struct {
	Resource_id                    string
	Resource_type                  string
	Transit_gateway_attachment_id  string
	Transit_gateway_route_table_id string
}

func Aws_ec2_transit_gateway_route_table_associationMapper(r *Aws_ec2_transit_gateway_route_table_association) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["transit_gateway_attachment_id"] = r.Transit_gateway_attachment_id
	config["transit_gateway_route_table_id"] = r.Transit_gateway_route_table_id
	config["resource_id"] = r.Resource_id
	config["resource_type"] = r.Resource_type
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_vpc struct {
	Ipv6_association_id              string
	Enable_classiclink_dns_support   *bool
	Default_security_group_id        string
	Default_route_table_id           string
	Ipv6_cidr_block                  string
	Cidr_block                       string
	Enable_dns_hostnames             *bool
	Dhcp_options_id                  string
	Default_network_acl_id           string
	Owner_id                         string
	Instance_tenancy                 *string
	Enable_dns_support               *bool
	Main_route_table_id              string
	Tags                             *map[string]interface{}
	Enable_classiclink               *bool
	Assign_generated_ipv6_cidr_block *bool
	Arn                              string
}

func Aws_vpcMapper(r *Aws_vpc) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Enable_classiclink != nil {
		config["enable_classiclink"] = *r.Enable_classiclink
	}
	if r.Assign_generated_ipv6_cidr_block != nil {
		config["assign_generated_ipv6_cidr_block"] = *r.Assign_generated_ipv6_cidr_block
	}
	config["arn"] = r.Arn
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	if r.Enable_classiclink_dns_support != nil {
		config["enable_classiclink_dns_support"] = *r.Enable_classiclink_dns_support
	}
	config["default_security_group_id"] = r.Default_security_group_id
	config["default_route_table_id"] = r.Default_route_table_id
	config["ipv6_association_id"] = r.Ipv6_association_id
	config["cidr_block"] = r.Cidr_block
	if r.Enable_dns_hostnames != nil {
		config["enable_dns_hostnames"] = *r.Enable_dns_hostnames
	}
	config["dhcp_options_id"] = r.Dhcp_options_id
	config["ipv6_cidr_block"] = r.Ipv6_cidr_block
	if r.Instance_tenancy != nil {
		config["instance_tenancy"] = *r.Instance_tenancy
	}
	if r.Enable_dns_support != nil {
		config["enable_dns_support"] = *r.Enable_dns_support
	}
	config["main_route_table_id"] = r.Main_route_table_id
	config["default_network_acl_id"] = r.Default_network_acl_id
	config["owner_id"] = r.Owner_id
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_alb_listener_certificate struct {
	Listener_arn    string
	Certificate_arn string
}

func Aws_alb_listener_certificateMapper(r *Aws_alb_listener_certificate) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["listener_arn"] = r.Listener_arn
	config["certificate_arn"] = r.Certificate_arn
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_acm_certificate_validation struct {
	Certificate_arn string
}

func Aws_acm_certificate_validationMapper(r *Aws_acm_certificate_validation) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["certificate_arn"] = r.Certificate_arn
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_budgets_budget struct {
	Time_unit         string
	Cost_filters      *map[string]interface{}
	Name_prefix       *string
	Limit_amount      string
	Limit_unit        string
	Time_period_start string
	Time_period_end   *string
	Account_id        *string
	Name              *string
	Budget_type       string
}

func Aws_budgets_budgetMapper(r *Aws_budgets_budget) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Name_prefix != nil {
		config["name_prefix"] = *r.Name_prefix
	}
	config["limit_amount"] = r.Limit_amount
	config["time_unit"] = r.Time_unit
	if r.Cost_filters != nil {
		config["cost_filters"] = *r.Cost_filters
	}
	if r.Time_period_end != nil {
		config["time_period_end"] = *r.Time_period_end
	}
	if r.Account_id != nil {
		config["account_id"] = *r.Account_id
	}
	if r.Name != nil {
		config["name"] = *r.Name
	}
	config["budget_type"] = r.Budget_type
	config["limit_unit"] = r.Limit_unit
	config["time_period_start"] = r.Time_period_start
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_lightsail_static_ip struct {
	Name         string
	Ip_address   string
	Arn          string
	Support_code string
}

func Aws_lightsail_static_ipMapper(r *Aws_lightsail_static_ip) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	config["ip_address"] = r.Ip_address
	config["arn"] = r.Arn
	config["support_code"] = r.Support_code
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_simpledb_domain struct {
	Name string
}

func Aws_simpledb_domainMapper(r *Aws_simpledb_domain) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_vpn_connection struct {
	Tunnel2_inside_cidr            *string
	Tags                           *map[string]interface{}
	Tunnel1_cgw_inside_address     string
	Tunnel1_vgw_inside_address     string
	Tunnel2_bgp_asn                string
	Vpn_gateway_id                 *string
	Resource_type                  string
	Tunnel1_bgp_asn                string
	Tunnel2_vgw_inside_address     string
	Transit_gateway_id             *string
	Tunnel1_inside_cidr            *string
	Tunnel2_preshared_key          *string
	Customer_gateway_configuration *string
	Tunnel2_cgw_inside_address     string
	Customer_gateway_id            string
	Static_routes_only             *bool
	Tunnel1_preshared_key          *string
	Tunnel1_address                string
	Tunnel2_address                string
}

func Aws_vpn_connectionMapper(r *Aws_vpn_connection) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Transit_gateway_id != nil {
		config["transit_gateway_id"] = *r.Transit_gateway_id
	}
	if r.Tunnel1_inside_cidr != nil {
		config["tunnel1_inside_cidr"] = *r.Tunnel1_inside_cidr
	}
	if r.Tunnel2_preshared_key != nil {
		config["tunnel2_preshared_key"] = *r.Tunnel2_preshared_key
	}
	if r.Customer_gateway_configuration != nil {
		config["customer_gateway_configuration"] = *r.Customer_gateway_configuration
	}
	config["tunnel2_cgw_inside_address"] = r.Tunnel2_cgw_inside_address
	config["customer_gateway_id"] = r.Customer_gateway_id
	if r.Static_routes_only != nil {
		config["static_routes_only"] = *r.Static_routes_only
	}
	if r.Tunnel1_preshared_key != nil {
		config["tunnel1_preshared_key"] = *r.Tunnel1_preshared_key
	}
	config["tunnel1_address"] = r.Tunnel1_address
	config["tunnel2_address"] = r.Tunnel2_address
	if r.Tunnel2_inside_cidr != nil {
		config["tunnel2_inside_cidr"] = *r.Tunnel2_inside_cidr
	}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["tunnel1_cgw_inside_address"] = r.Tunnel1_cgw_inside_address
	config["tunnel1_vgw_inside_address"] = r.Tunnel1_vgw_inside_address
	config["tunnel2_bgp_asn"] = r.Tunnel2_bgp_asn
	if r.Vpn_gateway_id != nil {
		config["vpn_gateway_id"] = *r.Vpn_gateway_id
	}
	config["resource_type"] = r.Resource_type
	config["tunnel1_bgp_asn"] = r.Tunnel1_bgp_asn
	config["tunnel2_vgw_inside_address"] = r.Tunnel2_vgw_inside_address
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_opsworks_user_profile struct {
	User_arn              string
	Allow_self_management *bool
	Ssh_username          string
	Ssh_public_key        *string
}

func Aws_opsworks_user_profileMapper(r *Aws_opsworks_user_profile) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["user_arn"] = r.User_arn
	if r.Allow_self_management != nil {
		config["allow_self_management"] = *r.Allow_self_management
	}
	config["ssh_username"] = r.Ssh_username
	if r.Ssh_public_key != nil {
		config["ssh_public_key"] = *r.Ssh_public_key
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_db_cluster_snapshot struct {
	Db_cluster_snapshot_arn        string
	Source_db_cluster_snapshot_arn string
	Status                         string
	Storage_encrypted              bool
	Kms_key_id                     string
	Snapshot_type                  string
	Vpc_id                         string
	Db_cluster_snapshot_identifier string
	Db_cluster_identifier          string
	Engine_version                 string
	Engine                         string
	License_model                  string
}

func Aws_db_cluster_snapshotMapper(r *Aws_db_cluster_snapshot) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["db_cluster_identifier"] = r.Db_cluster_identifier
	config["engine_version"] = r.Engine_version
	config["db_cluster_snapshot_identifier"] = r.Db_cluster_snapshot_identifier
	config["license_model"] = r.License_model
	config["engine"] = r.Engine
	config["db_cluster_snapshot_arn"] = r.Db_cluster_snapshot_arn
	config["source_db_cluster_snapshot_arn"] = r.Source_db_cluster_snapshot_arn
	config["status"] = r.Status
	config["kms_key_id"] = r.Kms_key_id
	config["snapshot_type"] = r.Snapshot_type
	config["vpc_id"] = r.Vpc_id
	config["storage_encrypted"] = r.Storage_encrypted
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_iot_topic_rule struct {
	Enabled     bool
	Sql         string
	Sql_version string
	Arn         string
	Description *string
	Name        string
}

func Aws_iot_topic_ruleMapper(r *Aws_iot_topic_rule) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["enabled"] = r.Enabled
	config["sql"] = r.Sql
	config["sql_version"] = r.Sql_version
	config["arn"] = r.Arn
	if r.Description != nil {
		config["description"] = *r.Description
	}
	config["name"] = r.Name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_ses_template struct {
	Name    string
	Html    *string
	Subject *string
	Text    *string
}

func Aws_ses_templateMapper(r *Aws_ses_template) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	if r.Html != nil {
		config["html"] = *r.Html
	}
	if r.Subject != nil {
		config["subject"] = *r.Subject
	}
	if r.Text != nil {
		config["text"] = *r.Text
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_volume_attachment struct {
	Device_name  string
	Instance_id  string
	Volume_id    string
	Force_detach *bool
	Skip_destroy *bool
}

func Aws_volume_attachmentMapper(r *Aws_volume_attachment) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["instance_id"] = r.Instance_id
	config["volume_id"] = r.Volume_id
	if r.Force_detach != nil {
		config["force_detach"] = *r.Force_detach
	}
	if r.Skip_destroy != nil {
		config["skip_destroy"] = *r.Skip_destroy
	}
	config["device_name"] = r.Device_name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_api_gateway_integration struct {
	Cache_namespace            *string
	Uri                        *string
	Request_parameters_in_json *string
	Content_handling           *string
	Credentials                *string
	Integration_http_method    *string
	Passthrough_behavior       *string
	Resource_id                string
	Http_method                string
	Resource_type              string
	Rest_api_id                string
	Connection_type            *string
	Connection_id              *string
	Request_templates          *map[string]interface{}
	Request_parameters         *map[string]interface{}
}

func Aws_api_gateway_integrationMapper(r *Aws_api_gateway_integration) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["resource_id"] = r.Resource_id
	config["http_method"] = r.Http_method
	config["resource_type"] = r.Resource_type
	if r.Credentials != nil {
		config["credentials"] = *r.Credentials
	}
	if r.Integration_http_method != nil {
		config["integration_http_method"] = *r.Integration_http_method
	}
	if r.Passthrough_behavior != nil {
		config["passthrough_behavior"] = *r.Passthrough_behavior
	}
	config["rest_api_id"] = r.Rest_api_id
	if r.Connection_type != nil {
		config["connection_type"] = *r.Connection_type
	}
	if r.Connection_id != nil {
		config["connection_id"] = *r.Connection_id
	}
	if r.Request_templates != nil {
		config["request_templates"] = *r.Request_templates
	}
	if r.Request_parameters != nil {
		config["request_parameters"] = *r.Request_parameters
	}
	if r.Uri != nil {
		config["uri"] = *r.Uri
	}
	if r.Request_parameters_in_json != nil {
		config["request_parameters_in_json"] = *r.Request_parameters_in_json
	}
	if r.Content_handling != nil {
		config["content_handling"] = *r.Content_handling
	}
	if r.Cache_namespace != nil {
		config["cache_namespace"] = *r.Cache_namespace
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_codedeploy_deployment_group struct {
	Deployment_config_name *string
	App_name               string
	Service_role_arn       string
	Deployment_group_name  string
}

func Aws_codedeploy_deployment_groupMapper(r *Aws_codedeploy_deployment_group) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["app_name"] = r.App_name
	config["service_role_arn"] = r.Service_role_arn
	config["deployment_group_name"] = r.Deployment_group_name
	if r.Deployment_config_name != nil {
		config["deployment_config_name"] = *r.Deployment_config_name
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_iam_group_policy struct {
	Name        *string
	Name_prefix *string
	Group       string
	Policy      string
}

func Aws_iam_group_policyMapper(r *Aws_iam_group_policy) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["policy"] = r.Policy
	if r.Name != nil {
		config["name"] = *r.Name
	}
	if r.Name_prefix != nil {
		config["name_prefix"] = *r.Name_prefix
	}
	config["group"] = r.Group
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_lambda_layer_version struct {
	S3_bucket         *string
	S3_object_version *string
	License_info      *string
	Layer_name        string
	Description       *string
	Filename          *string
	Source_code_hash  *string
	Arn               string
	Layer_arn         string
	Created_date      string
	Version           string
	S3_key            *string
}

func Aws_lambda_layer_versionMapper(r *Aws_lambda_layer_version) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["arn"] = r.Arn
	if r.Source_code_hash != nil {
		config["source_code_hash"] = *r.Source_code_hash
	}
	config["version"] = r.Version
	if r.S3_key != nil {
		config["s3_key"] = *r.S3_key
	}
	config["layer_arn"] = r.Layer_arn
	config["created_date"] = r.Created_date
	if r.License_info != nil {
		config["license_info"] = *r.License_info
	}
	config["layer_name"] = r.Layer_name
	if r.S3_bucket != nil {
		config["s3_bucket"] = *r.S3_bucket
	}
	if r.S3_object_version != nil {
		config["s3_object_version"] = *r.S3_object_version
	}
	if r.Filename != nil {
		config["filename"] = *r.Filename
	}
	if r.Description != nil {
		config["description"] = *r.Description
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_media_package_channel struct {
	Arn         string
	Channel_id  string
	Description *string
}

func Aws_media_package_channelMapper(r *Aws_media_package_channel) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Description != nil {
		config["description"] = *r.Description
	}
	config["arn"] = r.Arn
	config["channel_id"] = r.Channel_id
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_storagegateway_working_storage struct {
	Disk_id     string
	Gateway_arn string
}

func Aws_storagegateway_working_storageMapper(r *Aws_storagegateway_working_storage) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["disk_id"] = r.Disk_id
	config["gateway_arn"] = r.Gateway_arn
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_autoscaling_policy struct {
	Autoscaling_group_name  string
	Metric_aggregation_type *string
	Arn                     string
	Name                    string
	Adjustment_type         *string
	Policy_type             *string
}

func Aws_autoscaling_policyMapper(r *Aws_autoscaling_policy) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["arn"] = r.Arn
	config["name"] = r.Name
	if r.Adjustment_type != nil {
		config["adjustment_type"] = *r.Adjustment_type
	}
	if r.Policy_type != nil {
		config["policy_type"] = *r.Policy_type
	}
	if r.Metric_aggregation_type != nil {
		config["metric_aggregation_type"] = *r.Metric_aggregation_type
	}
	config["autoscaling_group_name"] = r.Autoscaling_group_name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_cognito_user_pool_client struct {
	Name                                 string
	User_pool_id                         string
	Generate_secret                      *bool
	Allowed_oauth_flows_user_pool_client *bool
	Default_redirect_uri                 *string
	Client_secret                        string
}

func Aws_cognito_user_pool_clientMapper(r *Aws_cognito_user_pool_client) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	config["user_pool_id"] = r.User_pool_id
	if r.Allowed_oauth_flows_user_pool_client != nil {
		config["allowed_oauth_flows_user_pool_client"] = *r.Allowed_oauth_flows_user_pool_client
	}
	if r.Default_redirect_uri != nil {
		config["default_redirect_uri"] = *r.Default_redirect_uri
	}
	if r.Generate_secret != nil {
		config["generate_secret"] = *r.Generate_secret
	}
	config["client_secret"] = r.Client_secret
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_codebuild_project struct {
	Badge_url      string
	Tags           *map[string]interface{}
	Description    *string
	Badge_enabled  *bool
	Encryption_key *string
	Arn            string
	Service_role   string
	Name           string
}

func Aws_codebuild_projectMapper(r *Aws_codebuild_project) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["arn"] = r.Arn
	config["service_role"] = r.Service_role
	config["name"] = r.Name
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	if r.Description != nil {
		config["description"] = *r.Description
	}
	if r.Badge_enabled != nil {
		config["badge_enabled"] = *r.Badge_enabled
	}
	config["badge_url"] = r.Badge_url
	if r.Encryption_key != nil {
		config["encryption_key"] = *r.Encryption_key
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_emr_security_configuration struct {
	Name          *string
	Name_prefix   *string
	Configuration string
	Creation_date string
}

func Aws_emr_security_configurationMapper(r *Aws_emr_security_configuration) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["configuration"] = r.Configuration
	config["creation_date"] = r.Creation_date
	if r.Name != nil {
		config["name"] = *r.Name
	}
	if r.Name_prefix != nil {
		config["name_prefix"] = *r.Name_prefix
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_iam_role struct {
	Tags                  *map[string]interface{}
	Name                  *string
	Name_prefix           *string
	Path                  *string
	Permissions_boundary  *string
	Description           *string
	Assume_role_policy    string
	Arn                   string
	Unique_id             string
	Force_detach_policies *bool
	Create_date           string
}

func Aws_iam_roleMapper(r *Aws_iam_role) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Permissions_boundary != nil {
		config["permissions_boundary"] = *r.Permissions_boundary
	}
	if r.Description != nil {
		config["description"] = *r.Description
	}
	config["assume_role_policy"] = r.Assume_role_policy
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	if r.Name != nil {
		config["name"] = *r.Name
	}
	if r.Name_prefix != nil {
		config["name_prefix"] = *r.Name_prefix
	}
	if r.Path != nil {
		config["path"] = *r.Path
	}
	config["create_date"] = r.Create_date
	config["arn"] = r.Arn
	config["unique_id"] = r.Unique_id
	if r.Force_detach_policies != nil {
		config["force_detach_policies"] = *r.Force_detach_policies
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_autoscaling_lifecycle_hook struct {
	Notification_metadata   *string
	Notification_target_arn *string
	Role_arn                *string
	Name                    string
	Autoscaling_group_name  string
	Default_result          *string
	Lifecycle_transition    string
}

func Aws_autoscaling_lifecycle_hookMapper(r *Aws_autoscaling_lifecycle_hook) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Notification_metadata != nil {
		config["notification_metadata"] = *r.Notification_metadata
	}
	if r.Notification_target_arn != nil {
		config["notification_target_arn"] = *r.Notification_target_arn
	}
	if r.Role_arn != nil {
		config["role_arn"] = *r.Role_arn
	}
	config["name"] = r.Name
	config["autoscaling_group_name"] = r.Autoscaling_group_name
	if r.Default_result != nil {
		config["default_result"] = *r.Default_result
	}
	config["lifecycle_transition"] = r.Lifecycle_transition
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_secretsmanager_secret_version struct {
	Arn           string
	Secret_id     string
	Secret_string *string
	Secret_binary *string
	Version_id    string
}

func Aws_secretsmanager_secret_versionMapper(r *Aws_secretsmanager_secret_version) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["arn"] = r.Arn
	config["secret_id"] = r.Secret_id
	if r.Secret_string != nil {
		config["secret_string"] = *r.Secret_string
	}
	if r.Secret_binary != nil {
		config["secret_binary"] = *r.Secret_binary
	}
	config["version_id"] = r.Version_id
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_ses_event_destination struct {
	Name                   string
	Configuration_set_name string
	Enabled                *bool
}

func Aws_ses_event_destinationMapper(r *Aws_ses_event_destination) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	config["configuration_set_name"] = r.Configuration_set_name
	if r.Enabled != nil {
		config["enabled"] = *r.Enabled
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_storagegateway_cache struct {
	Gateway_arn string
	Disk_id     string
}

func Aws_storagegateway_cacheMapper(r *Aws_storagegateway_cache) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["disk_id"] = r.Disk_id
	config["gateway_arn"] = r.Gateway_arn
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_appmesh_route struct {
	Arn                 string
	Created_date        string
	Last_updated_date   string
	Name                string
	Mesh_name           string
	Virtual_router_name string
}

func Aws_appmesh_routeMapper(r *Aws_appmesh_route) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	config["mesh_name"] = r.Mesh_name
	config["virtual_router_name"] = r.Virtual_router_name
	config["arn"] = r.Arn
	config["created_date"] = r.Created_date
	config["last_updated_date"] = r.Last_updated_date
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_codebuild_webhook struct {
	Url           string
	Project_name  string
	Branch_filter *string
	Payload_url   string
	Secret        string
}

func Aws_codebuild_webhookMapper(r *Aws_codebuild_webhook) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["url"] = r.Url
	config["project_name"] = r.Project_name
	if r.Branch_filter != nil {
		config["branch_filter"] = *r.Branch_filter
	}
	config["payload_url"] = r.Payload_url
	config["secret"] = r.Secret
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_glue_catalog_database struct {
	Catalog_id   *string
	Name         string
	Description  *string
	Location_uri *string
	Parameters   *map[string]interface{}
}

func Aws_glue_catalog_databaseMapper(r *Aws_glue_catalog_database) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Description != nil {
		config["description"] = *r.Description
	}
	if r.Location_uri != nil {
		config["location_uri"] = *r.Location_uri
	}
	if r.Parameters != nil {
		config["parameters"] = *r.Parameters
	}
	if r.Catalog_id != nil {
		config["catalog_id"] = *r.Catalog_id
	}
	config["name"] = r.Name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_glue_catalog_table struct {
	View_original_text *string
	Catalog_id         *string
	Name               string
	Owner              *string
	Table_type         *string
	Database_name      string
	Description        *string
	Parameters         *map[string]interface{}
	View_expanded_text *string
}

func Aws_glue_catalog_tableMapper(r *Aws_glue_catalog_table) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["database_name"] = r.Database_name
	if r.Description != nil {
		config["description"] = *r.Description
	}
	if r.Parameters != nil {
		config["parameters"] = *r.Parameters
	}
	if r.View_expanded_text != nil {
		config["view_expanded_text"] = *r.View_expanded_text
	}
	if r.Catalog_id != nil {
		config["catalog_id"] = *r.Catalog_id
	}
	config["name"] = r.Name
	if r.Owner != nil {
		config["owner"] = *r.Owner
	}
	if r.Table_type != nil {
		config["table_type"] = *r.Table_type
	}
	if r.View_original_text != nil {
		config["view_original_text"] = *r.View_original_text
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_lambda_function struct {
	Source_code_hash  *string
	Publish           *bool
	Arn               string
	Qualified_arn     string
	S3_key            *string
	Version           string
	Last_modified     string
	Tags              *map[string]interface{}
	Runtime           string
	Invoke_arn        string
	Kms_key_arn       *string
	S3_bucket         *string
	S3_object_version *string
	Description       *string
	Function_name     string
	Handler           string
	Role              string
	Filename          *string
}

func Aws_lambda_functionMapper(r *Aws_lambda_function) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.S3_key != nil {
		config["s3_key"] = *r.S3_key
	}
	config["arn"] = r.Arn
	config["qualified_arn"] = r.Qualified_arn
	config["version"] = r.Version
	config["last_modified"] = r.Last_modified
	config["runtime"] = r.Runtime
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	if r.S3_bucket != nil {
		config["s3_bucket"] = *r.S3_bucket
	}
	config["invoke_arn"] = r.Invoke_arn
	if r.Kms_key_arn != nil {
		config["kms_key_arn"] = *r.Kms_key_arn
	}
	if r.Filename != nil {
		config["filename"] = *r.Filename
	}
	if r.S3_object_version != nil {
		config["s3_object_version"] = *r.S3_object_version
	}
	if r.Description != nil {
		config["description"] = *r.Description
	}
	config["function_name"] = r.Function_name
	config["handler"] = r.Handler
	config["role"] = r.Role
	if r.Publish != nil {
		config["publish"] = *r.Publish
	}
	if r.Source_code_hash != nil {
		config["source_code_hash"] = *r.Source_code_hash
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_globalaccelerator_accelerator struct {
	Name            string
	Ip_address_type *string
	Enabled         *bool
}

func Aws_globalaccelerator_acceleratorMapper(r *Aws_globalaccelerator_accelerator) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	if r.Ip_address_type != nil {
		config["ip_address_type"] = *r.Ip_address_type
	}
	if r.Enabled != nil {
		config["enabled"] = *r.Enabled
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_transfer_ssh_key struct {
	Body      string
	Server_id string
	User_name string
}

func Aws_transfer_ssh_keyMapper(r *Aws_transfer_ssh_key) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["body"] = r.Body
	config["server_id"] = r.Server_id
	config["user_name"] = r.User_name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_glue_security_configuration struct {
	Name string
}

func Aws_glue_security_configurationMapper(r *Aws_glue_security_configuration) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_inspector_assessment_target struct {
	Name               string
	Arn                string
	Resource_group_arn *string
}

func Aws_inspector_assessment_targetMapper(r *Aws_inspector_assessment_target) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	config["arn"] = r.Arn
	if r.Resource_group_arn != nil {
		config["resource_group_arn"] = *r.Resource_group_arn
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_network_acl_rule struct {
	Icmp_type       *string
	Icmp_code       *string
	Egress          *bool
	Rule_action     string
	Cidr_block      *string
	Network_acl_id  string
	Protocol        string
	Ipv6_cidr_block *string
}

func Aws_network_acl_ruleMapper(r *Aws_network_acl_rule) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Icmp_type != nil {
		config["icmp_type"] = *r.Icmp_type
	}
	if r.Icmp_code != nil {
		config["icmp_code"] = *r.Icmp_code
	}
	if r.Egress != nil {
		config["egress"] = *r.Egress
	}
	config["rule_action"] = r.Rule_action
	if r.Cidr_block != nil {
		config["cidr_block"] = *r.Cidr_block
	}
	config["network_acl_id"] = r.Network_acl_id
	config["protocol"] = r.Protocol
	if r.Ipv6_cidr_block != nil {
		config["ipv6_cidr_block"] = *r.Ipv6_cidr_block
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_api_gateway_rest_api struct {
	Name             string
	Api_key_source   *string
	Created_date     string
	Description      *string
	Policy           *string
	Body             *string
	Root_resource_id string
	Execution_arn    string
}

func Aws_api_gateway_rest_apiMapper(r *Aws_api_gateway_rest_api) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	if r.Api_key_source != nil {
		config["api_key_source"] = *r.Api_key_source
	}
	config["created_date"] = r.Created_date
	if r.Description != nil {
		config["description"] = *r.Description
	}
	if r.Policy != nil {
		config["policy"] = *r.Policy
	}
	if r.Body != nil {
		config["body"] = *r.Body
	}
	config["root_resource_id"] = r.Root_resource_id
	config["execution_arn"] = r.Execution_arn
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_db_instance struct {
	Address                             string
	Instance_class                      string
	Maintenance_window                  *string
	Hosted_zone_id                      string
	Snapshot_identifier                 *string
	Identifier_prefix                   *string
	Backup_window                       *string
	Status                              string
	Allow_major_version_upgrade         *bool
	Domain_iam_role_name                *string
	Password                            *string
	Character_set_name                  *string
	Resource_id                         string
	Ca_cert_identifier                  string
	Username                            *string
	Apply_immediately                   *bool
	Monitoring_role_arn                 *string
	Publicly_accessible                 *bool
	Copy_tags_to_snapshot               *bool
	Kms_key_id                          *string
	Name                                *string
	Parameter_group_name                *string
	Arn                                 string
	Db_subnet_group_name                *string
	Domain                              *string
	Deletion_protection                 *bool
	Engine                              *string
	Availability_zone                   *string
	Multi_az                            *bool
	Option_group_name                   *string
	Storage_encrypted                   *bool
	Replicate_source_db                 *string
	Engine_version                      *string
	Storage_type                        *string
	License_model                       *string
	Auto_minor_version_upgrade          *bool
	Iam_database_authentication_enabled *bool
	Skip_final_snapshot                 *bool
	Tags                                *map[string]interface{}
	Identifier                          *string
	Final_snapshot_identifier           *string
	Endpoint                            string
	Timezone                            *string
}

func Aws_db_instanceMapper(r *Aws_db_instance) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Identifier != nil {
		config["identifier"] = *r.Identifier
	}
	if r.Final_snapshot_identifier != nil {
		config["final_snapshot_identifier"] = *r.Final_snapshot_identifier
	}
	config["endpoint"] = r.Endpoint
	if r.Timezone != nil {
		config["timezone"] = *r.Timezone
	}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["address"] = r.Address
	config["instance_class"] = r.Instance_class
	if r.Maintenance_window != nil {
		config["maintenance_window"] = *r.Maintenance_window
	}
	config["hosted_zone_id"] = r.Hosted_zone_id
	if r.Snapshot_identifier != nil {
		config["snapshot_identifier"] = *r.Snapshot_identifier
	}
	if r.Identifier_prefix != nil {
		config["identifier_prefix"] = *r.Identifier_prefix
	}
	if r.Backup_window != nil {
		config["backup_window"] = *r.Backup_window
	}
	config["status"] = r.Status
	if r.Allow_major_version_upgrade != nil {
		config["allow_major_version_upgrade"] = *r.Allow_major_version_upgrade
	}
	if r.Password != nil {
		config["password"] = *r.Password
	}
	if r.Character_set_name != nil {
		config["character_set_name"] = *r.Character_set_name
	}
	config["resource_id"] = r.Resource_id
	config["ca_cert_identifier"] = r.Ca_cert_identifier
	if r.Domain_iam_role_name != nil {
		config["domain_iam_role_name"] = *r.Domain_iam_role_name
	}
	if r.Username != nil {
		config["username"] = *r.Username
	}
	if r.Apply_immediately != nil {
		config["apply_immediately"] = *r.Apply_immediately
	}
	if r.Monitoring_role_arn != nil {
		config["monitoring_role_arn"] = *r.Monitoring_role_arn
	}
	if r.Publicly_accessible != nil {
		config["publicly_accessible"] = *r.Publicly_accessible
	}
	if r.Copy_tags_to_snapshot != nil {
		config["copy_tags_to_snapshot"] = *r.Copy_tags_to_snapshot
	}
	if r.Kms_key_id != nil {
		config["kms_key_id"] = *r.Kms_key_id
	}
	if r.Name != nil {
		config["name"] = *r.Name
	}
	if r.Parameter_group_name != nil {
		config["parameter_group_name"] = *r.Parameter_group_name
	}
	config["arn"] = r.Arn
	if r.Db_subnet_group_name != nil {
		config["db_subnet_group_name"] = *r.Db_subnet_group_name
	}
	if r.Domain != nil {
		config["domain"] = *r.Domain
	}
	if r.Deletion_protection != nil {
		config["deletion_protection"] = *r.Deletion_protection
	}
	if r.Engine != nil {
		config["engine"] = *r.Engine
	}
	if r.Availability_zone != nil {
		config["availability_zone"] = *r.Availability_zone
	}
	if r.Multi_az != nil {
		config["multi_az"] = *r.Multi_az
	}
	if r.Option_group_name != nil {
		config["option_group_name"] = *r.Option_group_name
	}
	if r.Storage_encrypted != nil {
		config["storage_encrypted"] = *r.Storage_encrypted
	}
	if r.Replicate_source_db != nil {
		config["replicate_source_db"] = *r.Replicate_source_db
	}
	if r.Engine_version != nil {
		config["engine_version"] = *r.Engine_version
	}
	if r.Storage_type != nil {
		config["storage_type"] = *r.Storage_type
	}
	if r.License_model != nil {
		config["license_model"] = *r.License_model
	}
	if r.Auto_minor_version_upgrade != nil {
		config["auto_minor_version_upgrade"] = *r.Auto_minor_version_upgrade
	}
	if r.Iam_database_authentication_enabled != nil {
		config["iam_database_authentication_enabled"] = *r.Iam_database_authentication_enabled
	}
	if r.Skip_final_snapshot != nil {
		config["skip_final_snapshot"] = *r.Skip_final_snapshot
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_directory_service_conditional_forwarder struct {
	Directory_id       string
	Remote_domain_name string
}

func Aws_directory_service_conditional_forwarderMapper(r *Aws_directory_service_conditional_forwarder) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["directory_id"] = r.Directory_id
	config["remote_domain_name"] = r.Remote_domain_name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_ec2_transit_gateway_route_table_propagation struct {
	Transit_gateway_attachment_id  string
	Transit_gateway_route_table_id string
	Resource_id                    string
	Resource_type                  string
}

func Aws_ec2_transit_gateway_route_table_propagationMapper(r *Aws_ec2_transit_gateway_route_table_propagation) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["resource_id"] = r.Resource_id
	config["resource_type"] = r.Resource_type
	config["transit_gateway_attachment_id"] = r.Transit_gateway_attachment_id
	config["transit_gateway_route_table_id"] = r.Transit_gateway_route_table_id
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_glacier_vault_lock struct {
	Ignore_deletion_error *bool
	Policy                string
	Vault_name            string
	Complete_lock         bool
}

func Aws_glacier_vault_lockMapper(r *Aws_glacier_vault_lock) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Ignore_deletion_error != nil {
		config["ignore_deletion_error"] = *r.Ignore_deletion_error
	}
	config["policy"] = r.Policy
	config["vault_name"] = r.Vault_name
	config["complete_lock"] = r.Complete_lock
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_wafregional_byte_match_set struct {
	Name string
}

func Aws_wafregional_byte_match_setMapper(r *Aws_wafregional_byte_match_set) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_lb_listener struct {
	Certificate_arn   *string
	Arn               string
	Load_balancer_arn string
	Protocol          *string
	Ssl_policy        *string
}

func Aws_lb_listenerMapper(r *Aws_lb_listener) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Ssl_policy != nil {
		config["ssl_policy"] = *r.Ssl_policy
	}
	if r.Certificate_arn != nil {
		config["certificate_arn"] = *r.Certificate_arn
	}
	config["arn"] = r.Arn
	config["load_balancer_arn"] = r.Load_balancer_arn
	if r.Protocol != nil {
		config["protocol"] = *r.Protocol
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_autoscaling_group struct {
	Force_delete              *bool
	Arn                       string
	Name                      *string
	Launch_configuration      *string
	Metrics_granularity       *string
	Name_prefix               *string
	Placement_group           *string
	Wait_for_capacity_timeout *string
	Service_linked_role_arn   *string
	Health_check_type         *string
	Protect_from_scale_in     *bool
}

func Aws_autoscaling_groupMapper(r *Aws_autoscaling_group) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Name != nil {
		config["name"] = *r.Name
	}
	if r.Name_prefix != nil {
		config["name_prefix"] = *r.Name_prefix
	}
	if r.Launch_configuration != nil {
		config["launch_configuration"] = *r.Launch_configuration
	}
	if r.Metrics_granularity != nil {
		config["metrics_granularity"] = *r.Metrics_granularity
	}
	if r.Placement_group != nil {
		config["placement_group"] = *r.Placement_group
	}
	if r.Wait_for_capacity_timeout != nil {
		config["wait_for_capacity_timeout"] = *r.Wait_for_capacity_timeout
	}
	if r.Service_linked_role_arn != nil {
		config["service_linked_role_arn"] = *r.Service_linked_role_arn
	}
	if r.Health_check_type != nil {
		config["health_check_type"] = *r.Health_check_type
	}
	if r.Protect_from_scale_in != nil {
		config["protect_from_scale_in"] = *r.Protect_from_scale_in
	}
	if r.Force_delete != nil {
		config["force_delete"] = *r.Force_delete
	}
	config["arn"] = r.Arn
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_elastic_beanstalk_application struct {
	Name        string
	Description *string
}

func Aws_elastic_beanstalk_applicationMapper(r *Aws_elastic_beanstalk_application) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	if r.Description != nil {
		config["description"] = *r.Description
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_sfn_state_machine struct {
	Definition    string
	Name          string
	Role_arn      string
	Creation_date string
	Status        string
	Tags          *map[string]interface{}
}

func Aws_sfn_state_machineMapper(r *Aws_sfn_state_machine) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["definition"] = r.Definition
	config["name"] = r.Name
	config["role_arn"] = r.Role_arn
	config["creation_date"] = r.Creation_date
	config["status"] = r.Status
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_cognito_user_pool struct {
	Mfa_configuration          *string
	Sms_authentication_message *string
	Tags                       *map[string]interface{}
	Name                       string
	Arn                        string
	Email_verification_subject *string
	Email_verification_message *string
	Endpoint                   string
	Sms_verification_message   *string
	Creation_date              string
	Last_modified_date         string
}

func Aws_cognito_user_poolMapper(r *Aws_cognito_user_pool) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Mfa_configuration != nil {
		config["mfa_configuration"] = *r.Mfa_configuration
	}
	if r.Sms_authentication_message != nil {
		config["sms_authentication_message"] = *r.Sms_authentication_message
	}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["arn"] = r.Arn
	if r.Email_verification_subject != nil {
		config["email_verification_subject"] = *r.Email_verification_subject
	}
	if r.Email_verification_message != nil {
		config["email_verification_message"] = *r.Email_verification_message
	}
	config["endpoint"] = r.Endpoint
	config["name"] = r.Name
	config["creation_date"] = r.Creation_date
	config["last_modified_date"] = r.Last_modified_date
	if r.Sms_verification_message != nil {
		config["sms_verification_message"] = *r.Sms_verification_message
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_guardduty_member struct {
	Account_id                 string
	Detector_id                string
	Email                      string
	Relationship_status        string
	Invite                     *bool
	Disable_email_notification *bool
	Invitation_message         *string
}

func Aws_guardduty_memberMapper(r *Aws_guardduty_member) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Disable_email_notification != nil {
		config["disable_email_notification"] = *r.Disable_email_notification
	}
	if r.Invitation_message != nil {
		config["invitation_message"] = *r.Invitation_message
	}
	config["account_id"] = r.Account_id
	config["detector_id"] = r.Detector_id
	config["email"] = r.Email
	config["relationship_status"] = r.Relationship_status
	if r.Invite != nil {
		config["invite"] = *r.Invite
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_s3_bucket_public_access_block struct {
	Block_public_policy     *bool
	Ignore_public_acls      *bool
	Restrict_public_buckets *bool
	Bucket                  string
	Block_public_acls       *bool
}

func Aws_s3_bucket_public_access_blockMapper(r *Aws_s3_bucket_public_access_block) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["bucket"] = r.Bucket
	if r.Block_public_acls != nil {
		config["block_public_acls"] = *r.Block_public_acls
	}
	if r.Block_public_policy != nil {
		config["block_public_policy"] = *r.Block_public_policy
	}
	if r.Ignore_public_acls != nil {
		config["ignore_public_acls"] = *r.Ignore_public_acls
	}
	if r.Restrict_public_buckets != nil {
		config["restrict_public_buckets"] = *r.Restrict_public_buckets
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_ssm_activation struct {
	Expiration_date *string
	Iam_role        string
	Activation_code string
	Name            *string
	Description     *string
	Expired         string
}

func Aws_ssm_activationMapper(r *Aws_ssm_activation) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["activation_code"] = r.Activation_code
	if r.Name != nil {
		config["name"] = *r.Name
	}
	if r.Description != nil {
		config["description"] = *r.Description
	}
	config["expired"] = r.Expired
	if r.Expiration_date != nil {
		config["expiration_date"] = *r.Expiration_date
	}
	config["iam_role"] = r.Iam_role
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_cloudfront_public_key struct {
	Caller_reference string
	Comment          *string
	Encoded_key      string
	Etag             string
	Name             *string
	Name_prefix      *string
}

func Aws_cloudfront_public_keyMapper(r *Aws_cloudfront_public_key) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Comment != nil {
		config["comment"] = *r.Comment
	}
	config["encoded_key"] = r.Encoded_key
	config["etag"] = r.Etag
	if r.Name != nil {
		config["name"] = *r.Name
	}
	if r.Name_prefix != nil {
		config["name_prefix"] = *r.Name_prefix
	}
	config["caller_reference"] = r.Caller_reference
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_kms_key struct {
	Is_enabled          *bool
	Tags                *map[string]interface{}
	Arn                 string
	Key_id              string
	Description         *string
	Key_usage           *string
	Policy              *string
	Enable_key_rotation *bool
}

func Aws_kms_keyMapper(r *Aws_kms_key) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["arn"] = r.Arn
	config["key_id"] = r.Key_id
	if r.Description != nil {
		config["description"] = *r.Description
	}
	if r.Key_usage != nil {
		config["key_usage"] = *r.Key_usage
	}
	if r.Is_enabled != nil {
		config["is_enabled"] = *r.Is_enabled
	}
	if r.Policy != nil {
		config["policy"] = *r.Policy
	}
	if r.Enable_key_rotation != nil {
		config["enable_key_rotation"] = *r.Enable_key_rotation
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_pinpoint_apns_sandbox_channel struct {
	Token_key                     *string
	Token_key_id                  *string
	Application_id                string
	Certificate                   *string
	Private_key                   *string
	Team_id                       *string
	Bundle_id                     *string
	Default_authentication_method *string
	Enabled                       *bool
}

func Aws_pinpoint_apns_sandbox_channelMapper(r *Aws_pinpoint_apns_sandbox_channel) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["application_id"] = r.Application_id
	if r.Certificate != nil {
		config["certificate"] = *r.Certificate
	}
	if r.Private_key != nil {
		config["private_key"] = *r.Private_key
	}
	if r.Token_key != nil {
		config["token_key"] = *r.Token_key
	}
	if r.Token_key_id != nil {
		config["token_key_id"] = *r.Token_key_id
	}
	if r.Bundle_id != nil {
		config["bundle_id"] = *r.Bundle_id
	}
	if r.Default_authentication_method != nil {
		config["default_authentication_method"] = *r.Default_authentication_method
	}
	if r.Enabled != nil {
		config["enabled"] = *r.Enabled
	}
	if r.Team_id != nil {
		config["team_id"] = *r.Team_id
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_dynamodb_table_item struct {
	Range_key  *string
	Item       string
	Table_name string
	Hash_key   string
}

func Aws_dynamodb_table_itemMapper(r *Aws_dynamodb_table_item) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["table_name"] = r.Table_name
	config["hash_key"] = r.Hash_key
	if r.Range_key != nil {
		config["range_key"] = *r.Range_key
	}
	config["item"] = r.Item
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_opsworks_php_app_layer struct {
	Install_updates_on_boot     *bool
	Custom_json                 *string
	Auto_assign_elastic_ips     *bool
	Elastic_load_balancer       *string
	Auto_healing                *bool
	Drain_elb_on_shutdown       *bool
	Name                        *string
	Custom_instance_profile_arn *string
	Stack_id                    string
	Use_ebs_optimized_instances *bool
	Auto_assign_public_ips      *bool
}

func Aws_opsworks_php_app_layerMapper(r *Aws_opsworks_php_app_layer) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Use_ebs_optimized_instances != nil {
		config["use_ebs_optimized_instances"] = *r.Use_ebs_optimized_instances
	}
	if r.Auto_assign_public_ips != nil {
		config["auto_assign_public_ips"] = *r.Auto_assign_public_ips
	}
	if r.Custom_json != nil {
		config["custom_json"] = *r.Custom_json
	}
	if r.Install_updates_on_boot != nil {
		config["install_updates_on_boot"] = *r.Install_updates_on_boot
	}
	if r.Elastic_load_balancer != nil {
		config["elastic_load_balancer"] = *r.Elastic_load_balancer
	}
	if r.Auto_healing != nil {
		config["auto_healing"] = *r.Auto_healing
	}
	if r.Drain_elb_on_shutdown != nil {
		config["drain_elb_on_shutdown"] = *r.Drain_elb_on_shutdown
	}
	if r.Auto_assign_elastic_ips != nil {
		config["auto_assign_elastic_ips"] = *r.Auto_assign_elastic_ips
	}
	config["stack_id"] = r.Stack_id
	if r.Name != nil {
		config["name"] = *r.Name
	}
	if r.Custom_instance_profile_arn != nil {
		config["custom_instance_profile_arn"] = *r.Custom_instance_profile_arn
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_snapshot_create_volume_permission struct {
	Snapshot_id string
	Account_id  string
}

func Aws_snapshot_create_volume_permissionMapper(r *Aws_snapshot_create_volume_permission) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["snapshot_id"] = r.Snapshot_id
	config["account_id"] = r.Account_id
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_alb_target_group_attachment struct {
	Target_group_arn  string
	Target_id         string
	Availability_zone *string
}

func Aws_alb_target_group_attachmentMapper(r *Aws_alb_target_group_attachment) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["target_group_arn"] = r.Target_group_arn
	config["target_id"] = r.Target_id
	if r.Availability_zone != nil {
		config["availability_zone"] = *r.Availability_zone
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_lb_target_group struct {
	Vpc_id            *string
	Name_prefix       *string
	Protocol          *string
	Proxy_protocol_v2 *bool
	Arn               string
	Name              *string
	Target_type       *string
	Arn_suffix        string
	Tags              *map[string]interface{}
}

func Aws_lb_target_groupMapper(r *Aws_lb_target_group) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["arn_suffix"] = r.Arn_suffix
	if r.Name != nil {
		config["name"] = *r.Name
	}
	if r.Target_type != nil {
		config["target_type"] = *r.Target_type
	}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	if r.Name_prefix != nil {
		config["name_prefix"] = *r.Name_prefix
	}
	if r.Vpc_id != nil {
		config["vpc_id"] = *r.Vpc_id
	}
	config["arn"] = r.Arn
	if r.Protocol != nil {
		config["protocol"] = *r.Protocol
	}
	if r.Proxy_protocol_v2 != nil {
		config["proxy_protocol_v2"] = *r.Proxy_protocol_v2
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_ami_launch_permission struct {
	Image_id   string
	Account_id string
}

func Aws_ami_launch_permissionMapper(r *Aws_ami_launch_permission) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["image_id"] = r.Image_id
	config["account_id"] = r.Account_id
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_api_gateway_integration_response struct {
	Rest_api_id                 string
	Http_method                 string
	Response_parameters         *map[string]interface{}
	Response_parameters_in_json *string
	Content_handling            *string
	Resource_id                 string
	Status_code                 string
	Selection_pattern           *string
	Response_templates          *map[string]interface{}
}

func Aws_api_gateway_integration_responseMapper(r *Aws_api_gateway_integration_response) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["status_code"] = r.Status_code
	if r.Selection_pattern != nil {
		config["selection_pattern"] = *r.Selection_pattern
	}
	if r.Response_templates != nil {
		config["response_templates"] = *r.Response_templates
	}
	config["resource_id"] = r.Resource_id
	config["http_method"] = r.Http_method
	if r.Response_parameters != nil {
		config["response_parameters"] = *r.Response_parameters
	}
	if r.Response_parameters_in_json != nil {
		config["response_parameters_in_json"] = *r.Response_parameters_in_json
	}
	if r.Content_handling != nil {
		config["content_handling"] = *r.Content_handling
	}
	config["rest_api_id"] = r.Rest_api_id
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_cloudwatch_log_metric_filter struct {
	Name           string
	Pattern        string
	Log_group_name string
}

func Aws_cloudwatch_log_metric_filterMapper(r *Aws_cloudwatch_log_metric_filter) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	config["pattern"] = r.Pattern
	config["log_group_name"] = r.Log_group_name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_opsworks_instance struct {
	Delete_ebs                   *bool
	Elastic_ip                   *string
	Hostname                     *string
	Instance_type                *string
	Os                           *string
	Ssh_host_dsa_key_fingerprint *string
	Delete_eip                   *bool
	Last_service_error_id        *string
	Private_dns                  *string
	Registered_by                *string
	Reported_os_version          *string
	State                        *string
	Status                       *string
	Ec2_instance_id              string
	Public_ip                    *string
	Ssh_key_name                 *string
	Agent_version                *string
	Ecs_cluster_arn              *string
	Public_dns                   *string
	Reported_os_family           *string
	Root_device_type             *string
	Ami_id                       *string
	Platform                     *string
	Root_device_volume_id        *string
	Auto_scaling_type            *string
	Created_at                   *string
	Ebs_optimized                *bool
	Private_ip                   *string
	Ssh_host_rsa_key_fingerprint *string
	Tenancy                      *string
	Architecture                 *string
	Availability_zone            *string
	Instance_profile_arn         *string
	Stack_id                     string
	Virtualization_type          *string
	Infrastructure_class         *string
	Install_updates_on_boot      *bool
	Reported_agent_version       *string
	Reported_os_name             *string
	Subnet_id                    *string
}

func Aws_opsworks_instanceMapper(r *Aws_opsworks_instance) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Ebs_optimized != nil {
		config["ebs_optimized"] = *r.Ebs_optimized
	}
	if r.Private_ip != nil {
		config["private_ip"] = *r.Private_ip
	}
	if r.Ssh_host_rsa_key_fingerprint != nil {
		config["ssh_host_rsa_key_fingerprint"] = *r.Ssh_host_rsa_key_fingerprint
	}
	if r.Tenancy != nil {
		config["tenancy"] = *r.Tenancy
	}
	if r.Auto_scaling_type != nil {
		config["auto_scaling_type"] = *r.Auto_scaling_type
	}
	if r.Created_at != nil {
		config["created_at"] = *r.Created_at
	}
	if r.Instance_profile_arn != nil {
		config["instance_profile_arn"] = *r.Instance_profile_arn
	}
	config["stack_id"] = r.Stack_id
	if r.Virtualization_type != nil {
		config["virtualization_type"] = *r.Virtualization_type
	}
	if r.Architecture != nil {
		config["architecture"] = *r.Architecture
	}
	if r.Availability_zone != nil {
		config["availability_zone"] = *r.Availability_zone
	}
	if r.Reported_agent_version != nil {
		config["reported_agent_version"] = *r.Reported_agent_version
	}
	if r.Reported_os_name != nil {
		config["reported_os_name"] = *r.Reported_os_name
	}
	if r.Subnet_id != nil {
		config["subnet_id"] = *r.Subnet_id
	}
	if r.Infrastructure_class != nil {
		config["infrastructure_class"] = *r.Infrastructure_class
	}
	if r.Install_updates_on_boot != nil {
		config["install_updates_on_boot"] = *r.Install_updates_on_boot
	}
	if r.Hostname != nil {
		config["hostname"] = *r.Hostname
	}
	if r.Instance_type != nil {
		config["instance_type"] = *r.Instance_type
	}
	if r.Os != nil {
		config["os"] = *r.Os
	}
	if r.Ssh_host_dsa_key_fingerprint != nil {
		config["ssh_host_dsa_key_fingerprint"] = *r.Ssh_host_dsa_key_fingerprint
	}
	if r.Delete_ebs != nil {
		config["delete_ebs"] = *r.Delete_ebs
	}
	if r.Elastic_ip != nil {
		config["elastic_ip"] = *r.Elastic_ip
	}
	if r.Private_dns != nil {
		config["private_dns"] = *r.Private_dns
	}
	if r.Registered_by != nil {
		config["registered_by"] = *r.Registered_by
	}
	if r.Reported_os_version != nil {
		config["reported_os_version"] = *r.Reported_os_version
	}
	if r.State != nil {
		config["state"] = *r.State
	}
	if r.Status != nil {
		config["status"] = *r.Status
	}
	if r.Delete_eip != nil {
		config["delete_eip"] = *r.Delete_eip
	}
	if r.Last_service_error_id != nil {
		config["last_service_error_id"] = *r.Last_service_error_id
	}
	if r.Ssh_key_name != nil {
		config["ssh_key_name"] = *r.Ssh_key_name
	}
	config["ec2_instance_id"] = r.Ec2_instance_id
	if r.Public_ip != nil {
		config["public_ip"] = *r.Public_ip
	}
	if r.Public_dns != nil {
		config["public_dns"] = *r.Public_dns
	}
	if r.Reported_os_family != nil {
		config["reported_os_family"] = *r.Reported_os_family
	}
	if r.Root_device_type != nil {
		config["root_device_type"] = *r.Root_device_type
	}
	if r.Agent_version != nil {
		config["agent_version"] = *r.Agent_version
	}
	if r.Ecs_cluster_arn != nil {
		config["ecs_cluster_arn"] = *r.Ecs_cluster_arn
	}
	if r.Root_device_volume_id != nil {
		config["root_device_volume_id"] = *r.Root_device_volume_id
	}
	if r.Ami_id != nil {
		config["ami_id"] = *r.Ami_id
	}
	if r.Platform != nil {
		config["platform"] = *r.Platform
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_vpc_endpoint_service_allowed_principal struct {
	Vpc_endpoint_service_id string
	Principal_arn           string
}

func Aws_vpc_endpoint_service_allowed_principalMapper(r *Aws_vpc_endpoint_service_allowed_principal) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["vpc_endpoint_service_id"] = r.Vpc_endpoint_service_id
	config["principal_arn"] = r.Principal_arn
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_dx_bgp_peer struct {
	Amazon_address       *string
	Bgp_auth_key         *string
	Customer_address     *string
	Bgp_status           string
	Address_family       string
	Virtual_interface_id string
}

func Aws_dx_bgp_peerMapper(r *Aws_dx_bgp_peer) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["address_family"] = r.Address_family
	config["virtual_interface_id"] = r.Virtual_interface_id
	if r.Amazon_address != nil {
		config["amazon_address"] = *r.Amazon_address
	}
	if r.Bgp_auth_key != nil {
		config["bgp_auth_key"] = *r.Bgp_auth_key
	}
	if r.Customer_address != nil {
		config["customer_address"] = *r.Customer_address
	}
	config["bgp_status"] = r.Bgp_status
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_licensemanager_license_configuration struct {
	License_count_hard_limit *bool
	License_counting_type    string
	Name                     string
	Tags                     *map[string]interface{}
	Description              *string
}

func Aws_licensemanager_license_configurationMapper(r *Aws_licensemanager_license_configuration) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Description != nil {
		config["description"] = *r.Description
	}
	if r.License_count_hard_limit != nil {
		config["license_count_hard_limit"] = *r.License_count_hard_limit
	}
	config["license_counting_type"] = r.License_counting_type
	config["name"] = r.Name
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_rds_cluster_parameter_group struct {
	Description *string
	Tags        *map[string]interface{}
	Arn         string
	Name        *string
	Name_prefix *string
	Family      string
}

func Aws_rds_cluster_parameter_groupMapper(r *Aws_rds_cluster_parameter_group) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["arn"] = r.Arn
	if r.Name != nil {
		config["name"] = *r.Name
	}
	if r.Name_prefix != nil {
		config["name_prefix"] = *r.Name_prefix
	}
	config["family"] = r.Family
	if r.Description != nil {
		config["description"] = *r.Description
	}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_pinpoint_event_stream struct {
	Application_id         string
	Destination_stream_arn string
	Role_arn               string
}

func Aws_pinpoint_event_streamMapper(r *Aws_pinpoint_event_stream) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["destination_stream_arn"] = r.Destination_stream_arn
	config["role_arn"] = r.Role_arn
	config["application_id"] = r.Application_id
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_api_gateway_deployment struct {
	Variables         *map[string]interface{}
	Created_date      string
	Invoke_url        string
	Execution_arn     string
	Rest_api_id       string
	Stage_name        string
	Description       *string
	Stage_description *string
}

func Aws_api_gateway_deploymentMapper(r *Aws_api_gateway_deployment) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["stage_name"] = r.Stage_name
	if r.Description != nil {
		config["description"] = *r.Description
	}
	if r.Stage_description != nil {
		config["stage_description"] = *r.Stage_description
	}
	if r.Variables != nil {
		config["variables"] = *r.Variables
	}
	config["created_date"] = r.Created_date
	config["invoke_url"] = r.Invoke_url
	config["execution_arn"] = r.Execution_arn
	config["rest_api_id"] = r.Rest_api_id
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_athena_named_query struct {
	Name        string
	Query       string
	Database    string
	Description *string
}

func Aws_athena_named_queryMapper(r *Aws_athena_named_query) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	config["query"] = r.Query
	config["database"] = r.Database
	if r.Description != nil {
		config["description"] = *r.Description
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_opsworks_nodejs_app_layer struct {
	Auto_healing                *bool
	Name                        *string
	Auto_assign_elastic_ips     *bool
	Custom_json                 *string
	Drain_elb_on_shutdown       *bool
	Install_updates_on_boot     *bool
	Stack_id                    string
	Elastic_load_balancer       *string
	Use_ebs_optimized_instances *bool
	Nodejs_version              *string
	Auto_assign_public_ips      *bool
	Custom_instance_profile_arn *string
}

func Aws_opsworks_nodejs_app_layerMapper(r *Aws_opsworks_nodejs_app_layer) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Drain_elb_on_shutdown != nil {
		config["drain_elb_on_shutdown"] = *r.Drain_elb_on_shutdown
	}
	if r.Install_updates_on_boot != nil {
		config["install_updates_on_boot"] = *r.Install_updates_on_boot
	}
	config["stack_id"] = r.Stack_id
	if r.Elastic_load_balancer != nil {
		config["elastic_load_balancer"] = *r.Elastic_load_balancer
	}
	if r.Use_ebs_optimized_instances != nil {
		config["use_ebs_optimized_instances"] = *r.Use_ebs_optimized_instances
	}
	if r.Nodejs_version != nil {
		config["nodejs_version"] = *r.Nodejs_version
	}
	if r.Auto_assign_public_ips != nil {
		config["auto_assign_public_ips"] = *r.Auto_assign_public_ips
	}
	if r.Custom_instance_profile_arn != nil {
		config["custom_instance_profile_arn"] = *r.Custom_instance_profile_arn
	}
	if r.Auto_healing != nil {
		config["auto_healing"] = *r.Auto_healing
	}
	if r.Name != nil {
		config["name"] = *r.Name
	}
	if r.Auto_assign_elastic_ips != nil {
		config["auto_assign_elastic_ips"] = *r.Auto_assign_elastic_ips
	}
	if r.Custom_json != nil {
		config["custom_json"] = *r.Custom_json
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_default_vpc struct {
	Enable_dns_hostnames             *bool
	Enable_dns_support               *bool
	Enable_classiclink               *bool
	Enable_classiclink_dns_support   *bool
	Default_security_group_id        string
	Ipv6_association_id              string
	Arn                              string
	Instance_tenancy                 string
	Tags                             *map[string]interface{}
	Owner_id                         string
	Assign_generated_ipv6_cidr_block bool
	Default_route_table_id           string
	Cidr_block                       string
	Main_route_table_id              string
	Default_network_acl_id           string
	Dhcp_options_id                  string
	Ipv6_cidr_block                  string
}

func Aws_default_vpcMapper(r *Aws_default_vpc) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["assign_generated_ipv6_cidr_block"] = r.Assign_generated_ipv6_cidr_block
	config["default_route_table_id"] = r.Default_route_table_id
	config["dhcp_options_id"] = r.Dhcp_options_id
	config["ipv6_cidr_block"] = r.Ipv6_cidr_block
	config["cidr_block"] = r.Cidr_block
	config["main_route_table_id"] = r.Main_route_table_id
	config["default_network_acl_id"] = r.Default_network_acl_id
	if r.Enable_classiclink_dns_support != nil {
		config["enable_classiclink_dns_support"] = *r.Enable_classiclink_dns_support
	}
	config["default_security_group_id"] = r.Default_security_group_id
	config["ipv6_association_id"] = r.Ipv6_association_id
	config["arn"] = r.Arn
	if r.Enable_dns_hostnames != nil {
		config["enable_dns_hostnames"] = *r.Enable_dns_hostnames
	}
	if r.Enable_dns_support != nil {
		config["enable_dns_support"] = *r.Enable_dns_support
	}
	if r.Enable_classiclink != nil {
		config["enable_classiclink"] = *r.Enable_classiclink
	}
	config["instance_tenancy"] = r.Instance_tenancy
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["owner_id"] = r.Owner_id
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_elasticache_replication_group struct {
	Replication_group_description  string
	Auth_token                     *string
	Engine                         *string
	Replication_group_id           string
	Snapshot_window                *string
	Transit_encryption_enabled     *bool
	Auto_minor_version_upgrade     *bool
	Automatic_failover_enabled     *bool
	Subnet_group_name              *string
	Snapshot_name                  *string
	Maintenance_window             *string
	Parameter_group_name           *string
	At_rest_encryption_enabled     *bool
	Node_type                      *string
	Notification_topic_arn         *string
	Apply_immediately              *bool
	Tags                           *map[string]interface{}
	Configuration_endpoint_address string
	Engine_version                 *string
	Primary_endpoint_address       string
}

func Aws_elasticache_replication_groupMapper(r *Aws_elasticache_replication_group) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Node_type != nil {
		config["node_type"] = *r.Node_type
	}
	if r.At_rest_encryption_enabled != nil {
		config["at_rest_encryption_enabled"] = *r.At_rest_encryption_enabled
	}
	if r.Notification_topic_arn != nil {
		config["notification_topic_arn"] = *r.Notification_topic_arn
	}
	if r.Apply_immediately != nil {
		config["apply_immediately"] = *r.Apply_immediately
	}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["configuration_endpoint_address"] = r.Configuration_endpoint_address
	config["primary_endpoint_address"] = r.Primary_endpoint_address
	if r.Engine_version != nil {
		config["engine_version"] = *r.Engine_version
	}
	if r.Engine != nil {
		config["engine"] = *r.Engine
	}
	config["replication_group_description"] = r.Replication_group_description
	if r.Auth_token != nil {
		config["auth_token"] = *r.Auth_token
	}
	if r.Automatic_failover_enabled != nil {
		config["automatic_failover_enabled"] = *r.Automatic_failover_enabled
	}
	config["replication_group_id"] = r.Replication_group_id
	if r.Snapshot_window != nil {
		config["snapshot_window"] = *r.Snapshot_window
	}
	if r.Transit_encryption_enabled != nil {
		config["transit_encryption_enabled"] = *r.Transit_encryption_enabled
	}
	if r.Auto_minor_version_upgrade != nil {
		config["auto_minor_version_upgrade"] = *r.Auto_minor_version_upgrade
	}
	if r.Snapshot_name != nil {
		config["snapshot_name"] = *r.Snapshot_name
	}
	if r.Subnet_group_name != nil {
		config["subnet_group_name"] = *r.Subnet_group_name
	}
	if r.Parameter_group_name != nil {
		config["parameter_group_name"] = *r.Parameter_group_name
	}
	if r.Maintenance_window != nil {
		config["maintenance_window"] = *r.Maintenance_window
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_guardduty_ipset struct {
	Detector_id string
	Name        string
	Format      string
	Location    string
	Activate    bool
}

func Aws_guardduty_ipsetMapper(r *Aws_guardduty_ipset) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["activate"] = r.Activate
	config["detector_id"] = r.Detector_id
	config["name"] = r.Name
	config["format"] = r.Format
	config["location"] = r.Location
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_network_interface_sg_attachment struct {
	Security_group_id    string
	Network_interface_id string
}

func Aws_network_interface_sg_attachmentMapper(r *Aws_network_interface_sg_attachment) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["network_interface_id"] = r.Network_interface_id
	config["security_group_id"] = r.Security_group_id
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_wafregional_rule_group struct {
	Name        string
	Metric_name string
}

func Aws_wafregional_rule_groupMapper(r *Aws_wafregional_rule_group) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	config["metric_name"] = r.Metric_name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_api_gateway_stage struct {
	Cache_cluster_size    *string
	Invoke_url            string
	Variables             *map[string]interface{}
	Execution_arn         string
	Stage_name            string
	Xray_tracing_enabled  *bool
	Cache_cluster_enabled *bool
	Client_certificate_id *string
	Description           *string
	Tags                  *map[string]interface{}
	Deployment_id         string
	Documentation_version *string
	Rest_api_id           string
}

func Aws_api_gateway_stageMapper(r *Aws_api_gateway_stage) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["rest_api_id"] = r.Rest_api_id
	config["deployment_id"] = r.Deployment_id
	if r.Documentation_version != nil {
		config["documentation_version"] = *r.Documentation_version
	}
	if r.Cache_cluster_size != nil {
		config["cache_cluster_size"] = *r.Cache_cluster_size
	}
	config["invoke_url"] = r.Invoke_url
	if r.Variables != nil {
		config["variables"] = *r.Variables
	}
	config["execution_arn"] = r.Execution_arn
	config["stage_name"] = r.Stage_name
	if r.Xray_tracing_enabled != nil {
		config["xray_tracing_enabled"] = *r.Xray_tracing_enabled
	}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	if r.Cache_cluster_enabled != nil {
		config["cache_cluster_enabled"] = *r.Cache_cluster_enabled
	}
	if r.Client_certificate_id != nil {
		config["client_certificate_id"] = *r.Client_certificate_id
	}
	if r.Description != nil {
		config["description"] = *r.Description
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_dx_hosted_private_virtual_interface_accepter struct {
	Virtual_interface_id string
	Vpn_gateway_id       *string
	Dx_gateway_id        *string
	Tags                 *map[string]interface{}
	Arn                  string
}

func Aws_dx_hosted_private_virtual_interface_accepterMapper(r *Aws_dx_hosted_private_virtual_interface_accepter) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["arn"] = r.Arn
	config["virtual_interface_id"] = r.Virtual_interface_id
	if r.Vpn_gateway_id != nil {
		config["vpn_gateway_id"] = *r.Vpn_gateway_id
	}
	if r.Dx_gateway_id != nil {
		config["dx_gateway_id"] = *r.Dx_gateway_id
	}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_iot_policy_attachment struct {
	Policy string
	Target string
}

func Aws_iot_policy_attachmentMapper(r *Aws_iot_policy_attachment) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["policy"] = r.Policy
	config["target"] = r.Target
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_kinesis_analytics_application struct {
	Status                string
	Name                  string
	Arn                   string
	Create_timestamp      string
	Description           *string
	Last_update_timestamp string
	Code                  *string
}

func Aws_kinesis_analytics_applicationMapper(r *Aws_kinesis_analytics_application) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Code != nil {
		config["code"] = *r.Code
	}
	if r.Description != nil {
		config["description"] = *r.Description
	}
	config["last_update_timestamp"] = r.Last_update_timestamp
	config["status"] = r.Status
	config["name"] = r.Name
	config["arn"] = r.Arn
	config["create_timestamp"] = r.Create_timestamp
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_vpc_endpoint_subnet_association struct {
	Subnet_id       string
	Vpc_endpoint_id string
}

func Aws_vpc_endpoint_subnet_associationMapper(r *Aws_vpc_endpoint_subnet_association) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["vpc_endpoint_id"] = r.Vpc_endpoint_id
	config["subnet_id"] = r.Subnet_id
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_directory_service_directory struct {
	Password          string
	Edition           *string
	Name              string
	Size              *string
	Alias             *string
	Short_name        *string
	Security_group_id string
	Tags              *map[string]interface{}
	Access_url        string
	Resource_type     *string
	Description       *string
	Enable_sso        *bool
}

func Aws_directory_service_directoryMapper(r *Aws_directory_service_directory) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["password"] = r.Password
	if r.Edition != nil {
		config["edition"] = *r.Edition
	}
	config["security_group_id"] = r.Security_group_id
	config["name"] = r.Name
	if r.Size != nil {
		config["size"] = *r.Size
	}
	if r.Alias != nil {
		config["alias"] = *r.Alias
	}
	if r.Short_name != nil {
		config["short_name"] = *r.Short_name
	}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["access_url"] = r.Access_url
	if r.Resource_type != nil {
		config["resource_type"] = *r.Resource_type
	}
	if r.Description != nil {
		config["description"] = *r.Description
	}
	if r.Enable_sso != nil {
		config["enable_sso"] = *r.Enable_sso
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_ecs_task_definition struct {
	Arn                   string
	Container_definitions string
	Ipc_mode              *string
	Tags                  *map[string]interface{}
	Cpu                   *string
	Pid_mode              *string
	Family                string
	Execution_role_arn    *string
	Memory                *string
	Network_mode          *string
	Task_role_arn         *string
}

func Aws_ecs_task_definitionMapper(r *Aws_ecs_task_definition) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Pid_mode != nil {
		config["pid_mode"] = *r.Pid_mode
	}
	if r.Cpu != nil {
		config["cpu"] = *r.Cpu
	}
	if r.Execution_role_arn != nil {
		config["execution_role_arn"] = *r.Execution_role_arn
	}
	if r.Memory != nil {
		config["memory"] = *r.Memory
	}
	if r.Network_mode != nil {
		config["network_mode"] = *r.Network_mode
	}
	config["family"] = r.Family
	if r.Task_role_arn != nil {
		config["task_role_arn"] = *r.Task_role_arn
	}
	config["container_definitions"] = r.Container_definitions
	if r.Ipc_mode != nil {
		config["ipc_mode"] = *r.Ipc_mode
	}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["arn"] = r.Arn
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_licensemanager_association struct {
	License_configuration_arn string
	Resource_arn              string
}

func Aws_licensemanager_associationMapper(r *Aws_licensemanager_association) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["license_configuration_arn"] = r.License_configuration_arn
	config["resource_arn"] = r.Resource_arn
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_ses_configuration_set struct {
	Name string
}

func Aws_ses_configuration_setMapper(r *Aws_ses_configuration_set) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_ssm_association struct {
	Name                string
	Parameters          *map[string]interface{}
	Schedule_expression *string
	Document_version    *string
	Association_id      string
	Instance_id         *string
	Association_name    *string
}

func Aws_ssm_associationMapper(r *Aws_ssm_association) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["association_id"] = r.Association_id
	if r.Instance_id != nil {
		config["instance_id"] = *r.Instance_id
	}
	if r.Association_name != nil {
		config["association_name"] = *r.Association_name
	}
	config["name"] = r.Name
	if r.Parameters != nil {
		config["parameters"] = *r.Parameters
	}
	if r.Schedule_expression != nil {
		config["schedule_expression"] = *r.Schedule_expression
	}
	if r.Document_version != nil {
		config["document_version"] = *r.Document_version
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_pinpoint_apns_voip_channel struct {
	Default_authentication_method *string
	Team_id                       *string
	Token_key                     *string
	Certificate                   *string
	Bundle_id                     *string
	Enabled                       *bool
	Private_key                   *string
	Token_key_id                  *string
	Application_id                string
}

func Aws_pinpoint_apns_voip_channelMapper(r *Aws_pinpoint_apns_voip_channel) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Bundle_id != nil {
		config["bundle_id"] = *r.Bundle_id
	}
	if r.Enabled != nil {
		config["enabled"] = *r.Enabled
	}
	if r.Private_key != nil {
		config["private_key"] = *r.Private_key
	}
	if r.Token_key_id != nil {
		config["token_key_id"] = *r.Token_key_id
	}
	config["application_id"] = r.Application_id
	if r.Default_authentication_method != nil {
		config["default_authentication_method"] = *r.Default_authentication_method
	}
	if r.Team_id != nil {
		config["team_id"] = *r.Team_id
	}
	if r.Token_key != nil {
		config["token_key"] = *r.Token_key
	}
	if r.Certificate != nil {
		config["certificate"] = *r.Certificate
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_api_gateway_vpc_link struct {
	Name        string
	Description *string
}

func Aws_api_gateway_vpc_linkMapper(r *Aws_api_gateway_vpc_link) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	if r.Description != nil {
		config["description"] = *r.Description
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_glacier_vault struct {
	Arn           string
	Access_policy *string
	Tags          *map[string]interface{}
	Name          string
	Location      string
}

func Aws_glacier_vaultMapper(r *Aws_glacier_vault) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["name"] = r.Name
	config["location"] = r.Location
	config["arn"] = r.Arn
	if r.Access_policy != nil {
		config["access_policy"] = *r.Access_policy
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_rds_cluster_instance struct {
	Publicly_accessible             *bool
	Instance_class                  string
	Engine_version                  *string
	Availability_zone               *string
	Performance_insights_kms_key_id *string
	Identifier                      *string
	Writer                          bool
	Cluster_identifier              string
	Db_parameter_group_name         *string
	Apply_immediately               *bool
	Copy_tags_to_snapshot           *bool
	Tags                            *map[string]interface{}
	Arn                             string
	Kms_key_id                      string
	Dbi_resource_id                 string
	Preferred_maintenance_window    *string
	Db_subnet_group_name            *string
	Engine                          *string
	Identifier_prefix               *string
	Monitoring_role_arn             *string
	Performance_insights_enabled    *bool
	Auto_minor_version_upgrade      *bool
	Preferred_backup_window         *string
	Storage_encrypted               bool
	Endpoint                        string
}

func Aws_rds_cluster_instanceMapper(r *Aws_rds_cluster_instance) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["endpoint"] = r.Endpoint
	config["instance_class"] = r.Instance_class
	if r.Engine_version != nil {
		config["engine_version"] = *r.Engine_version
	}
	if r.Availability_zone != nil {
		config["availability_zone"] = *r.Availability_zone
	}
	if r.Performance_insights_kms_key_id != nil {
		config["performance_insights_kms_key_id"] = *r.Performance_insights_kms_key_id
	}
	if r.Identifier != nil {
		config["identifier"] = *r.Identifier
	}
	if r.Publicly_accessible != nil {
		config["publicly_accessible"] = *r.Publicly_accessible
	}
	config["cluster_identifier"] = r.Cluster_identifier
	if r.Db_parameter_group_name != nil {
		config["db_parameter_group_name"] = *r.Db_parameter_group_name
	}
	if r.Apply_immediately != nil {
		config["apply_immediately"] = *r.Apply_immediately
	}
	if r.Copy_tags_to_snapshot != nil {
		config["copy_tags_to_snapshot"] = *r.Copy_tags_to_snapshot
	}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["arn"] = r.Arn
	config["writer"] = r.Writer
	config["dbi_resource_id"] = r.Dbi_resource_id
	if r.Preferred_maintenance_window != nil {
		config["preferred_maintenance_window"] = *r.Preferred_maintenance_window
	}
	if r.Db_subnet_group_name != nil {
		config["db_subnet_group_name"] = *r.Db_subnet_group_name
	}
	config["kms_key_id"] = r.Kms_key_id
	if r.Identifier_prefix != nil {
		config["identifier_prefix"] = *r.Identifier_prefix
	}
	if r.Engine != nil {
		config["engine"] = *r.Engine
	}
	if r.Performance_insights_enabled != nil {
		config["performance_insights_enabled"] = *r.Performance_insights_enabled
	}
	if r.Auto_minor_version_upgrade != nil {
		config["auto_minor_version_upgrade"] = *r.Auto_minor_version_upgrade
	}
	if r.Monitoring_role_arn != nil {
		config["monitoring_role_arn"] = *r.Monitoring_role_arn
	}
	config["storage_encrypted"] = r.Storage_encrypted
	if r.Preferred_backup_window != nil {
		config["preferred_backup_window"] = *r.Preferred_backup_window
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_redshift_subnet_group struct {
	Name        string
	Description *string
	Tags        *map[string]interface{}
}

func Aws_redshift_subnet_groupMapper(r *Aws_redshift_subnet_group) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	if r.Description != nil {
		config["description"] = *r.Description
	}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_lightsail_domain struct {
	Domain_name string
	Arn         string
}

func Aws_lightsail_domainMapper(r *Aws_lightsail_domain) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["domain_name"] = r.Domain_name
	config["arn"] = r.Arn
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_media_store_container struct {
	Name     string
	Arn      string
	Endpoint string
}

func Aws_media_store_containerMapper(r *Aws_media_store_container) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	config["arn"] = r.Arn
	config["endpoint"] = r.Endpoint
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_service_discovery_http_namespace struct {
	Name        string
	Description *string
	Arn         string
}

func Aws_service_discovery_http_namespaceMapper(r *Aws_service_discovery_http_namespace) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	if r.Description != nil {
		config["description"] = *r.Description
	}
	config["arn"] = r.Arn
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_cognito_user_group struct {
	Name         string
	Role_arn     *string
	User_pool_id string
	Description  *string
}

func Aws_cognito_user_groupMapper(r *Aws_cognito_user_group) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Description != nil {
		config["description"] = *r.Description
	}
	config["name"] = r.Name
	if r.Role_arn != nil {
		config["role_arn"] = *r.Role_arn
	}
	config["user_pool_id"] = r.User_pool_id
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_dms_replication_subnet_group struct {
	Replication_subnet_group_arn         string
	Replication_subnet_group_description string
	Replication_subnet_group_id          string
	Tags                                 *map[string]interface{}
	Vpc_id                               string
}

func Aws_dms_replication_subnet_groupMapper(r *Aws_dms_replication_subnet_group) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["replication_subnet_group_id"] = r.Replication_subnet_group_id
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["vpc_id"] = r.Vpc_id
	config["replication_subnet_group_arn"] = r.Replication_subnet_group_arn
	config["replication_subnet_group_description"] = r.Replication_subnet_group_description
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_opsworks_permission struct {
	Level      *string
	Stack_id   *string
	Allow_ssh  *bool
	Allow_sudo *bool
	User_arn   string
}

func Aws_opsworks_permissionMapper(r *Aws_opsworks_permission) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Allow_ssh != nil {
		config["allow_ssh"] = *r.Allow_ssh
	}
	if r.Allow_sudo != nil {
		config["allow_sudo"] = *r.Allow_sudo
	}
	config["user_arn"] = r.User_arn
	if r.Level != nil {
		config["level"] = *r.Level
	}
	if r.Stack_id != nil {
		config["stack_id"] = *r.Stack_id
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_vpc_dhcp_options_association struct {
	Vpc_id          string
	Dhcp_options_id string
}

func Aws_vpc_dhcp_options_associationMapper(r *Aws_vpc_dhcp_options_association) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["vpc_id"] = r.Vpc_id
	config["dhcp_options_id"] = r.Dhcp_options_id
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_inspector_assessment_template struct {
	Name       string
	Target_arn string
	Arn        string
}

func Aws_inspector_assessment_templateMapper(r *Aws_inspector_assessment_template) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["arn"] = r.Arn
	config["name"] = r.Name
	config["target_arn"] = r.Target_arn
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_kinesis_stream struct {
	Encryption_type *string
	Kms_key_id      *string
	Arn             *string
	Tags            *map[string]interface{}
	Name            string
}

func Aws_kinesis_streamMapper(r *Aws_kinesis_stream) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Encryption_type != nil {
		config["encryption_type"] = *r.Encryption_type
	}
	if r.Kms_key_id != nil {
		config["kms_key_id"] = *r.Kms_key_id
	}
	if r.Arn != nil {
		config["arn"] = *r.Arn
	}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["name"] = r.Name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_vpc_ipv4_cidr_block_association struct {
	Vpc_id     string
	Cidr_block string
}

func Aws_vpc_ipv4_cidr_block_associationMapper(r *Aws_vpc_ipv4_cidr_block_association) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["vpc_id"] = r.Vpc_id
	config["cidr_block"] = r.Cidr_block
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_waf_geo_match_set struct {
	Name string
}

func Aws_waf_geo_match_setMapper(r *Aws_waf_geo_match_set) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_pinpoint_apns_channel struct {
	Certificate                   *string
	Default_authentication_method *string
	Private_key                   *string
	Token_key                     *string
	Bundle_id                     *string
	Enabled                       *bool
	Team_id                       *string
	Token_key_id                  *string
	Application_id                string
}

func Aws_pinpoint_apns_channelMapper(r *Aws_pinpoint_apns_channel) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Bundle_id != nil {
		config["bundle_id"] = *r.Bundle_id
	}
	if r.Certificate != nil {
		config["certificate"] = *r.Certificate
	}
	if r.Default_authentication_method != nil {
		config["default_authentication_method"] = *r.Default_authentication_method
	}
	if r.Private_key != nil {
		config["private_key"] = *r.Private_key
	}
	if r.Token_key != nil {
		config["token_key"] = *r.Token_key
	}
	config["application_id"] = r.Application_id
	if r.Enabled != nil {
		config["enabled"] = *r.Enabled
	}
	if r.Team_id != nil {
		config["team_id"] = *r.Team_id
	}
	if r.Token_key_id != nil {
		config["token_key_id"] = *r.Token_key_id
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_db_parameter_group struct {
	Name_prefix *string
	Family      string
	Description *string
	Tags        *map[string]interface{}
	Arn         string
	Name        *string
}

func Aws_db_parameter_groupMapper(r *Aws_db_parameter_group) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Description != nil {
		config["description"] = *r.Description
	}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["arn"] = r.Arn
	if r.Name != nil {
		config["name"] = *r.Name
	}
	if r.Name_prefix != nil {
		config["name_prefix"] = *r.Name_prefix
	}
	config["family"] = r.Family
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_ec2_fleet struct {
	Replace_unhealthy_instances         *bool
	Terminate_instances                 *bool
	Terminate_instances_with_expiration *bool
	Resource_type                       *string
	Tags                                *map[string]interface{}
	Excess_capacity_termination_policy  *string
}

func Aws_ec2_fleetMapper(r *Aws_ec2_fleet) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Terminate_instances != nil {
		config["terminate_instances"] = *r.Terminate_instances
	}
	if r.Terminate_instances_with_expiration != nil {
		config["terminate_instances_with_expiration"] = *r.Terminate_instances_with_expiration
	}
	if r.Resource_type != nil {
		config["resource_type"] = *r.Resource_type
	}
	if r.Replace_unhealthy_instances != nil {
		config["replace_unhealthy_instances"] = *r.Replace_unhealthy_instances
	}
	if r.Excess_capacity_termination_policy != nil {
		config["excess_capacity_termination_policy"] = *r.Excess_capacity_termination_policy
	}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_gamelift_game_session_queue struct {
	Name string
	Arn  string
}

func Aws_gamelift_game_session_queueMapper(r *Aws_gamelift_game_session_queue) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	config["arn"] = r.Arn
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_spot_datafeed_subscription struct {
	Bucket string
	Prefix *string
}

func Aws_spot_datafeed_subscriptionMapper(r *Aws_spot_datafeed_subscription) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["bucket"] = r.Bucket
	if r.Prefix != nil {
		config["prefix"] = *r.Prefix
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_alb_target_group struct {
	Tags              *map[string]interface{}
	Arn               string
	Arn_suffix        string
	Protocol          *string
	Name              *string
	Name_prefix       *string
	Proxy_protocol_v2 *bool
	Target_type       *string
	Vpc_id            *string
}

func Aws_alb_target_groupMapper(r *Aws_alb_target_group) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Vpc_id != nil {
		config["vpc_id"] = *r.Vpc_id
	}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["arn"] = r.Arn
	config["arn_suffix"] = r.Arn_suffix
	if r.Protocol != nil {
		config["protocol"] = *r.Protocol
	}
	if r.Name != nil {
		config["name"] = *r.Name
	}
	if r.Name_prefix != nil {
		config["name_prefix"] = *r.Name_prefix
	}
	if r.Proxy_protocol_v2 != nil {
		config["proxy_protocol_v2"] = *r.Proxy_protocol_v2
	}
	if r.Target_type != nil {
		config["target_type"] = *r.Target_type
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_appautoscaling_policy struct {
	Resource_id             string
	Service_namespace       string
	Arn                     string
	Policy_type             *string
	Scalable_dimension      string
	Adjustment_type         *string
	Name                    string
	Metric_aggregation_type *string
}

func Aws_appautoscaling_policyMapper(r *Aws_appautoscaling_policy) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	if r.Metric_aggregation_type != nil {
		config["metric_aggregation_type"] = *r.Metric_aggregation_type
	}
	config["resource_id"] = r.Resource_id
	config["service_namespace"] = r.Service_namespace
	config["arn"] = r.Arn
	if r.Policy_type != nil {
		config["policy_type"] = *r.Policy_type
	}
	config["scalable_dimension"] = r.Scalable_dimension
	if r.Adjustment_type != nil {
		config["adjustment_type"] = *r.Adjustment_type
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_redshift_snapshot_copy_grant struct {
	Snapshot_copy_grant_name string
	Kms_key_id               *string
	Tags                     *map[string]interface{}
}

func Aws_redshift_snapshot_copy_grantMapper(r *Aws_redshift_snapshot_copy_grant) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["snapshot_copy_grant_name"] = r.Snapshot_copy_grant_name
	if r.Kms_key_id != nil {
		config["kms_key_id"] = *r.Kms_key_id
	}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_default_security_group struct {
	Vpc_id                 *string
	Name                   string
	Arn                    string
	Owner_id               string
	Tags                   *map[string]interface{}
	Revoke_rules_on_delete *bool
}

func Aws_default_security_groupMapper(r *Aws_default_security_group) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	if r.Revoke_rules_on_delete != nil {
		config["revoke_rules_on_delete"] = *r.Revoke_rules_on_delete
	}
	config["name"] = r.Name
	config["arn"] = r.Arn
	config["owner_id"] = r.Owner_id
	if r.Vpc_id != nil {
		config["vpc_id"] = *r.Vpc_id
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_ebs_snapshot_copy struct {
	Source_snapshot_id     string
	Volume_id              string
	Description            *string
	Owner_id               string
	Owner_alias            string
	Encrypted              *bool
	Kms_key_id             *string
	Source_region          string
	Data_encryption_key_id string
	Tags                   *map[string]interface{}
}

func Aws_ebs_snapshot_copyMapper(r *Aws_ebs_snapshot_copy) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["source_region"] = r.Source_region
	config["source_snapshot_id"] = r.Source_snapshot_id
	config["volume_id"] = r.Volume_id
	if r.Description != nil {
		config["description"] = *r.Description
	}
	config["owner_id"] = r.Owner_id
	config["owner_alias"] = r.Owner_alias
	if r.Encrypted != nil {
		config["encrypted"] = *r.Encrypted
	}
	if r.Kms_key_id != nil {
		config["kms_key_id"] = *r.Kms_key_id
	}
	config["data_encryption_key_id"] = r.Data_encryption_key_id
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_spot_fleet_request struct {
	Iam_fleet_role                      string
	Excess_capacity_termination_policy  *string
	Replace_unhealthy_instances         *bool
	Terminate_instances_with_expiration *bool
	Fleet_type                          *string
	Spot_request_state                  string
	Allocation_strategy                 *string
	Spot_price                          *string
	Valid_from                          *string
	Valid_until                         *string
	Wait_for_fulfillment                *bool
	Instance_interruption_behaviour     *string
	Client_token                        string
}

func Aws_spot_fleet_requestMapper(r *Aws_spot_fleet_request) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["iam_fleet_role"] = r.Iam_fleet_role
	if r.Excess_capacity_termination_policy != nil {
		config["excess_capacity_termination_policy"] = *r.Excess_capacity_termination_policy
	}
	if r.Replace_unhealthy_instances != nil {
		config["replace_unhealthy_instances"] = *r.Replace_unhealthy_instances
	}
	if r.Terminate_instances_with_expiration != nil {
		config["terminate_instances_with_expiration"] = *r.Terminate_instances_with_expiration
	}
	if r.Fleet_type != nil {
		config["fleet_type"] = *r.Fleet_type
	}
	config["spot_request_state"] = r.Spot_request_state
	if r.Valid_until != nil {
		config["valid_until"] = *r.Valid_until
	}
	if r.Allocation_strategy != nil {
		config["allocation_strategy"] = *r.Allocation_strategy
	}
	if r.Spot_price != nil {
		config["spot_price"] = *r.Spot_price
	}
	if r.Valid_from != nil {
		config["valid_from"] = *r.Valid_from
	}
	if r.Wait_for_fulfillment != nil {
		config["wait_for_fulfillment"] = *r.Wait_for_fulfillment
	}
	if r.Instance_interruption_behaviour != nil {
		config["instance_interruption_behaviour"] = *r.Instance_interruption_behaviour
	}
	config["client_token"] = r.Client_token
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_nat_gateway struct {
	Public_ip            string
	Tags                 *map[string]interface{}
	Allocation_id        string
	Subnet_id            string
	Network_interface_id string
	Private_ip           string
}

func Aws_nat_gatewayMapper(r *Aws_nat_gateway) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["private_ip"] = r.Private_ip
	config["public_ip"] = r.Public_ip
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["allocation_id"] = r.Allocation_id
	config["subnet_id"] = r.Subnet_id
	config["network_interface_id"] = r.Network_interface_id
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_organizations_account struct {
	Name                       string
	Email                      string
	Iam_user_access_to_billing *string
	Role_name                  *string
	Arn                        string
	Joined_method              string
	Joined_timestamp           string
	Status                     string
}

func Aws_organizations_accountMapper(r *Aws_organizations_account) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["joined_timestamp"] = r.Joined_timestamp
	config["status"] = r.Status
	config["name"] = r.Name
	config["email"] = r.Email
	if r.Iam_user_access_to_billing != nil {
		config["iam_user_access_to_billing"] = *r.Iam_user_access_to_billing
	}
	if r.Role_name != nil {
		config["role_name"] = *r.Role_name
	}
	config["arn"] = r.Arn
	config["joined_method"] = r.Joined_method
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_ebs_snapshot struct {
	Volume_id              string
	Description            *string
	Owner_alias            string
	Encrypted              bool
	Kms_key_id             string
	Data_encryption_key_id string
	Tags                   *map[string]interface{}
	Owner_id               string
}

func Aws_ebs_snapshotMapper(r *Aws_ebs_snapshot) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["owner_id"] = r.Owner_id
	config["kms_key_id"] = r.Kms_key_id
	config["data_encryption_key_id"] = r.Data_encryption_key_id
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["volume_id"] = r.Volume_id
	if r.Description != nil {
		config["description"] = *r.Description
	}
	config["owner_alias"] = r.Owner_alias
	config["encrypted"] = r.Encrypted
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_elasticache_subnet_group struct {
	Description *string
	Name        string
}

func Aws_elasticache_subnet_groupMapper(r *Aws_elasticache_subnet_group) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Description != nil {
		config["description"] = *r.Description
	}
	config["name"] = r.Name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_elasticsearch_domain struct {
	Access_policies       *string
	Domain_id             string
	Endpoint              string
	Tags                  *map[string]interface{}
	Arn                   string
	Elasticsearch_version *string
	Domain_name           string
	Advanced_options      *map[string]interface{}
	Kibana_endpoint       string
}

func Aws_elasticsearch_domainMapper(r *Aws_elasticsearch_domain) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["arn"] = r.Arn
	if r.Elasticsearch_version != nil {
		config["elasticsearch_version"] = *r.Elasticsearch_version
	}
	config["domain_name"] = r.Domain_name
	if r.Advanced_options != nil {
		config["advanced_options"] = *r.Advanced_options
	}
	config["kibana_endpoint"] = r.Kibana_endpoint
	if r.Access_policies != nil {
		config["access_policies"] = *r.Access_policies
	}
	config["domain_id"] = r.Domain_id
	config["endpoint"] = r.Endpoint
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_elb struct {
	Name_prefix               *string
	Internal                  *bool
	Dns_name                  string
	Zone_id                   string
	Name                      *string
	Arn                       string
	Cross_zone_load_balancing *bool
	Source_security_group     *string
	Source_security_group_id  string
	Connection_draining       *bool
	Tags                      *map[string]interface{}
}

func Aws_elbMapper(r *Aws_elb) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Source_security_group != nil {
		config["source_security_group"] = *r.Source_security_group
	}
	config["source_security_group_id"] = r.Source_security_group_id
	if r.Connection_draining != nil {
		config["connection_draining"] = *r.Connection_draining
	}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	if r.Name_prefix != nil {
		config["name_prefix"] = *r.Name_prefix
	}
	if r.Internal != nil {
		config["internal"] = *r.Internal
	}
	config["dns_name"] = r.Dns_name
	if r.Name != nil {
		config["name"] = *r.Name
	}
	config["arn"] = r.Arn
	if r.Cross_zone_load_balancing != nil {
		config["cross_zone_load_balancing"] = *r.Cross_zone_load_balancing
	}
	config["zone_id"] = r.Zone_id
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_iam_user_policy_attachment struct {
	User       string
	Policy_arn string
}

func Aws_iam_user_policy_attachmentMapper(r *Aws_iam_user_policy_attachment) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["user"] = r.User
	config["policy_arn"] = r.Policy_arn
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_gamelift_fleet struct {
	Arn                                string
	Build_id                           string
	New_game_session_protection_policy *string
	Ec2_instance_type                  string
	Name                               string
	Description                        *string
	Operating_system                   string
}

func Aws_gamelift_fleetMapper(r *Aws_gamelift_fleet) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["arn"] = r.Arn
	config["build_id"] = r.Build_id
	if r.New_game_session_protection_policy != nil {
		config["new_game_session_protection_policy"] = *r.New_game_session_protection_policy
	}
	config["ec2_instance_type"] = r.Ec2_instance_type
	config["name"] = r.Name
	if r.Description != nil {
		config["description"] = *r.Description
	}
	config["operating_system"] = r.Operating_system
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_iam_role_policy struct {
	Policy      string
	Name        *string
	Name_prefix *string
	Role        string
}

func Aws_iam_role_policyMapper(r *Aws_iam_role_policy) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["policy"] = r.Policy
	if r.Name != nil {
		config["name"] = *r.Name
	}
	if r.Name_prefix != nil {
		config["name_prefix"] = *r.Name_prefix
	}
	config["role"] = r.Role
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_ssm_patch_group struct {
	Baseline_id string
	Patch_group string
}

func Aws_ssm_patch_groupMapper(r *Aws_ssm_patch_group) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["baseline_id"] = r.Baseline_id
	config["patch_group"] = r.Patch_group
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_lb_target_group_attachment struct {
	Availability_zone *string
	Target_group_arn  string
	Target_id         string
}

func Aws_lb_target_group_attachmentMapper(r *Aws_lb_target_group_attachment) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Availability_zone != nil {
		config["availability_zone"] = *r.Availability_zone
	}
	config["target_group_arn"] = r.Target_group_arn
	config["target_id"] = r.Target_id
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_dx_hosted_private_virtual_interface struct {
	Amazon_address      *string
	Owner_account_id    string
	Arn                 string
	Address_family      string
	Bgp_auth_key        *string
	Customer_address    *string
	Jumbo_frame_capable bool
	Connection_id       string
	Name                string
}

func Aws_dx_hosted_private_virtual_interfaceMapper(r *Aws_dx_hosted_private_virtual_interface) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["owner_account_id"] = r.Owner_account_id
	config["arn"] = r.Arn
	config["address_family"] = r.Address_family
	if r.Amazon_address != nil {
		config["amazon_address"] = *r.Amazon_address
	}
	if r.Bgp_auth_key != nil {
		config["bgp_auth_key"] = *r.Bgp_auth_key
	}
	if r.Customer_address != nil {
		config["customer_address"] = *r.Customer_address
	}
	config["jumbo_frame_capable"] = r.Jumbo_frame_capable
	config["connection_id"] = r.Connection_id
	config["name"] = r.Name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_ec2_transit_gateway_route struct {
	Destination_cidr_block         string
	Transit_gateway_attachment_id  string
	Transit_gateway_route_table_id string
}

func Aws_ec2_transit_gateway_routeMapper(r *Aws_ec2_transit_gateway_route) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["destination_cidr_block"] = r.Destination_cidr_block
	config["transit_gateway_attachment_id"] = r.Transit_gateway_attachment_id
	config["transit_gateway_route_table_id"] = r.Transit_gateway_route_table_id
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_opsworks_application struct {
	Data_source_type          *string
	Description               *string
	Enable_ssl                *bool
	Document_root             *string
	Aws_flow_ruby_settings    *string
	Data_source_database_name *string
	Rails_env                 *string
	Auto_bundle_on_deploy     *string
	Data_source_arn           *string
	Name                      string
	Short_name                *string
	Resource_type             string
	Stack_id                  string
}

func Aws_opsworks_applicationMapper(r *Aws_opsworks_application) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["resource_type"] = r.Resource_type
	config["stack_id"] = r.Stack_id
	if r.Rails_env != nil {
		config["rails_env"] = *r.Rails_env
	}
	if r.Auto_bundle_on_deploy != nil {
		config["auto_bundle_on_deploy"] = *r.Auto_bundle_on_deploy
	}
	if r.Data_source_arn != nil {
		config["data_source_arn"] = *r.Data_source_arn
	}
	config["name"] = r.Name
	if r.Short_name != nil {
		config["short_name"] = *r.Short_name
	}
	if r.Data_source_type != nil {
		config["data_source_type"] = *r.Data_source_type
	}
	if r.Description != nil {
		config["description"] = *r.Description
	}
	if r.Data_source_database_name != nil {
		config["data_source_database_name"] = *r.Data_source_database_name
	}
	if r.Enable_ssl != nil {
		config["enable_ssl"] = *r.Enable_ssl
	}
	if r.Document_root != nil {
		config["document_root"] = *r.Document_root
	}
	if r.Aws_flow_ruby_settings != nil {
		config["aws_flow_ruby_settings"] = *r.Aws_flow_ruby_settings
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_api_gateway_domain_name struct {
	Regional_zone_id          string
	Certificate_private_key   *string
	Regional_certificate_name *string
	Certificate_chain         *string
	Cloudfront_zone_id        string
	Certificate_name          *string
	Domain_name               string
	Cloudfront_domain_name    string
	Certificate_upload_date   string
	Regional_certificate_arn  *string
	Regional_domain_name      string
	Certificate_body          *string
	Certificate_arn           *string
}

func Aws_api_gateway_domain_nameMapper(r *Aws_api_gateway_domain_name) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Certificate_private_key != nil {
		config["certificate_private_key"] = *r.Certificate_private_key
	}
	if r.Regional_certificate_name != nil {
		config["regional_certificate_name"] = *r.Regional_certificate_name
	}
	config["regional_zone_id"] = r.Regional_zone_id
	if r.Certificate_chain != nil {
		config["certificate_chain"] = *r.Certificate_chain
	}
	if r.Certificate_name != nil {
		config["certificate_name"] = *r.Certificate_name
	}
	config["domain_name"] = r.Domain_name
	config["cloudfront_zone_id"] = r.Cloudfront_zone_id
	if r.Regional_certificate_arn != nil {
		config["regional_certificate_arn"] = *r.Regional_certificate_arn
	}
	config["regional_domain_name"] = r.Regional_domain_name
	if r.Certificate_body != nil {
		config["certificate_body"] = *r.Certificate_body
	}
	if r.Certificate_arn != nil {
		config["certificate_arn"] = *r.Certificate_arn
	}
	config["cloudfront_domain_name"] = r.Cloudfront_domain_name
	config["certificate_upload_date"] = r.Certificate_upload_date
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_elastic_beanstalk_application_version struct {
	Application  string
	Description  *string
	Bucket       string
	Key          string
	Name         string
	Force_delete *bool
}

func Aws_elastic_beanstalk_application_versionMapper(r *Aws_elastic_beanstalk_application_version) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	if r.Force_delete != nil {
		config["force_delete"] = *r.Force_delete
	}
	config["application"] = r.Application
	if r.Description != nil {
		config["description"] = *r.Description
	}
	config["bucket"] = r.Bucket
	config["key"] = r.Key
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_iam_user_login_profile struct {
	Key_fingerprint         string
	Encrypted_password      string
	User                    string
	Pgp_key                 string
	Password_reset_required *bool
}

func Aws_iam_user_login_profileMapper(r *Aws_iam_user_login_profile) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["user"] = r.User
	config["pgp_key"] = r.Pgp_key
	if r.Password_reset_required != nil {
		config["password_reset_required"] = *r.Password_reset_required
	}
	config["key_fingerprint"] = r.Key_fingerprint
	config["encrypted_password"] = r.Encrypted_password
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_default_network_acl struct {
	Vpc_id                 string
	Default_network_acl_id string
	Tags                   *map[string]interface{}
	Owner_id               string
}

func Aws_default_network_aclMapper(r *Aws_default_network_acl) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["owner_id"] = r.Owner_id
	config["vpc_id"] = r.Vpc_id
	config["default_network_acl_id"] = r.Default_network_acl_id
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_neptune_subnet_group struct {
	Name_prefix *string
	Description *string
	Tags        *map[string]interface{}
	Arn         string
	Name        *string
}

func Aws_neptune_subnet_groupMapper(r *Aws_neptune_subnet_group) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["arn"] = r.Arn
	if r.Name != nil {
		config["name"] = *r.Name
	}
	if r.Name_prefix != nil {
		config["name_prefix"] = *r.Name_prefix
	}
	if r.Description != nil {
		config["description"] = *r.Description
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_iam_account_password_policy struct {
	Allow_users_to_change_password *bool
	Hard_expiry                    *bool
	Require_uppercase_characters   *bool
	Expire_passwords               bool
	Require_lowercase_characters   *bool
	Require_numbers                *bool
	Require_symbols                *bool
}

func Aws_iam_account_password_policyMapper(r *Aws_iam_account_password_policy) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Allow_users_to_change_password != nil {
		config["allow_users_to_change_password"] = *r.Allow_users_to_change_password
	}
	if r.Hard_expiry != nil {
		config["hard_expiry"] = *r.Hard_expiry
	}
	config["expire_passwords"] = r.Expire_passwords
	if r.Require_lowercase_characters != nil {
		config["require_lowercase_characters"] = *r.Require_lowercase_characters
	}
	if r.Require_numbers != nil {
		config["require_numbers"] = *r.Require_numbers
	}
	if r.Require_symbols != nil {
		config["require_symbols"] = *r.Require_symbols
	}
	if r.Require_uppercase_characters != nil {
		config["require_uppercase_characters"] = *r.Require_uppercase_characters
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_ses_domain_identity_verification struct {
	Arn    string
	Domain string
}

func Aws_ses_domain_identity_verificationMapper(r *Aws_ses_domain_identity_verification) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["arn"] = r.Arn
	config["domain"] = r.Domain
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_wafregional_geo_match_set struct {
	Name string
}

func Aws_wafregional_geo_match_setMapper(r *Aws_wafregional_geo_match_set) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_ecr_repository_policy struct {
	Repository  string
	Policy      string
	Registry_id string
}

func Aws_ecr_repository_policyMapper(r *Aws_ecr_repository_policy) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["registry_id"] = r.Registry_id
	config["repository"] = r.Repository
	config["policy"] = r.Policy
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_iam_group_membership struct {
	Group string
	Name  string
}

func Aws_iam_group_membershipMapper(r *Aws_iam_group_membership) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	config["group"] = r.Group
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_lb_cookie_stickiness_policy struct {
	Name          string
	Load_balancer string
}

func Aws_lb_cookie_stickiness_policyMapper(r *Aws_lb_cookie_stickiness_policy) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	config["load_balancer"] = r.Load_balancer
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_servicecatalog_portfolio struct {
	Arn           string
	Created_time  string
	Name          string
	Description   *string
	Provider_name *string
	Tags          *map[string]interface{}
}

func Aws_servicecatalog_portfolioMapper(r *Aws_servicecatalog_portfolio) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["arn"] = r.Arn
	config["created_time"] = r.Created_time
	config["name"] = r.Name
	if r.Description != nil {
		config["description"] = *r.Description
	}
	if r.Provider_name != nil {
		config["provider_name"] = *r.Provider_name
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_ssm_maintenance_window_target struct {
	Window_id         string
	Resource_type     string
	Owner_information *string
}

func Aws_ssm_maintenance_window_targetMapper(r *Aws_ssm_maintenance_window_target) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["window_id"] = r.Window_id
	config["resource_type"] = r.Resource_type
	if r.Owner_information != nil {
		config["owner_information"] = *r.Owner_information
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_cognito_user_pool_domain struct {
	Domain                      string
	Certificate_arn             *string
	User_pool_id                string
	Aws_account_id              string
	Cloudfront_distribution_arn string
	S3_bucket                   string
	Version                     string
}

func Aws_cognito_user_pool_domainMapper(r *Aws_cognito_user_pool_domain) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["user_pool_id"] = r.User_pool_id
	config["aws_account_id"] = r.Aws_account_id
	config["cloudfront_distribution_arn"] = r.Cloudfront_distribution_arn
	config["s3_bucket"] = r.S3_bucket
	config["version"] = r.Version
	config["domain"] = r.Domain
	if r.Certificate_arn != nil {
		config["certificate_arn"] = *r.Certificate_arn
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_cognito_resource_server struct {
	Name         string
	User_pool_id string
	Identifier   string
}

func Aws_cognito_resource_serverMapper(r *Aws_cognito_resource_server) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["user_pool_id"] = r.User_pool_id
	config["identifier"] = r.Identifier
	config["name"] = r.Name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_iam_role_policy_attachment struct {
	Role       string
	Policy_arn string
}

func Aws_iam_role_policy_attachmentMapper(r *Aws_iam_role_policy_attachment) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["policy_arn"] = r.Policy_arn
	config["role"] = r.Role
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_sns_topic_subscription struct {
	Endpoint               string
	Topic_arn              string
	Delivery_policy        *string
	Raw_message_delivery   *bool
	Arn                    string
	Protocol               string
	Filter_policy          *string
	Endpoint_auto_confirms *bool
}

func Aws_sns_topic_subscriptionMapper(r *Aws_sns_topic_subscription) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Endpoint_auto_confirms != nil {
		config["endpoint_auto_confirms"] = *r.Endpoint_auto_confirms
	}
	if r.Filter_policy != nil {
		config["filter_policy"] = *r.Filter_policy
	}
	config["protocol"] = r.Protocol
	config["endpoint"] = r.Endpoint
	config["topic_arn"] = r.Topic_arn
	if r.Delivery_policy != nil {
		config["delivery_policy"] = *r.Delivery_policy
	}
	if r.Raw_message_delivery != nil {
		config["raw_message_delivery"] = *r.Raw_message_delivery
	}
	config["arn"] = r.Arn
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_wafregional_regex_match_set struct {
	Name string
}

func Aws_wafregional_regex_match_setMapper(r *Aws_wafregional_regex_match_set) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_api_gateway_method_response struct {
	Http_method                 string
	Status_code                 string
	Response_models             *map[string]interface{}
	Response_parameters         *map[string]interface{}
	Response_parameters_in_json *string
	Rest_api_id                 string
	Resource_id                 string
}

func Aws_api_gateway_method_responseMapper(r *Aws_api_gateway_method_response) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["rest_api_id"] = r.Rest_api_id
	config["resource_id"] = r.Resource_id
	config["http_method"] = r.Http_method
	config["status_code"] = r.Status_code
	if r.Response_models != nil {
		config["response_models"] = *r.Response_models
	}
	if r.Response_parameters != nil {
		config["response_parameters"] = *r.Response_parameters
	}
	if r.Response_parameters_in_json != nil {
		config["response_parameters_in_json"] = *r.Response_parameters_in_json
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_elasticache_cluster struct {
	Az_mode                *string
	Parameter_group_name   *string
	Snapshot_name          *string
	Cluster_id             string
	Node_type              *string
	Availability_zone      *string
	Cluster_address        string
	Notification_topic_arn *string
	Replication_group_id   *string
	Snapshot_window        *string
	Subnet_group_name      *string
	Apply_immediately      *bool
	Configuration_endpoint string
	Engine                 *string
	Engine_version         *string
	Maintenance_window     *string
	Tags                   *map[string]interface{}
}

func Aws_elasticache_clusterMapper(r *Aws_elasticache_cluster) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Maintenance_window != nil {
		config["maintenance_window"] = *r.Maintenance_window
	}
	if r.Subnet_group_name != nil {
		config["subnet_group_name"] = *r.Subnet_group_name
	}
	if r.Apply_immediately != nil {
		config["apply_immediately"] = *r.Apply_immediately
	}
	config["configuration_endpoint"] = r.Configuration_endpoint
	if r.Engine != nil {
		config["engine"] = *r.Engine
	}
	if r.Engine_version != nil {
		config["engine_version"] = *r.Engine_version
	}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	if r.Az_mode != nil {
		config["az_mode"] = *r.Az_mode
	}
	if r.Parameter_group_name != nil {
		config["parameter_group_name"] = *r.Parameter_group_name
	}
	if r.Snapshot_name != nil {
		config["snapshot_name"] = *r.Snapshot_name
	}
	config["cluster_id"] = r.Cluster_id
	if r.Node_type != nil {
		config["node_type"] = *r.Node_type
	}
	if r.Snapshot_window != nil {
		config["snapshot_window"] = *r.Snapshot_window
	}
	if r.Availability_zone != nil {
		config["availability_zone"] = *r.Availability_zone
	}
	config["cluster_address"] = r.Cluster_address
	if r.Notification_topic_arn != nil {
		config["notification_topic_arn"] = *r.Notification_topic_arn
	}
	if r.Replication_group_id != nil {
		config["replication_group_id"] = *r.Replication_group_id
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_cloudwatch_log_destination_policy struct {
	Destination_name string
	Access_policy    string
}

func Aws_cloudwatch_log_destination_policyMapper(r *Aws_cloudwatch_log_destination_policy) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["destination_name"] = r.Destination_name
	config["access_policy"] = r.Access_policy
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_emr_instance_group struct {
	Ebs_optimized *bool
	Cluster_id    string
	Instance_type string
	Status        string
	Name          *string
}

func Aws_emr_instance_groupMapper(r *Aws_emr_instance_group) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["status"] = r.Status
	if r.Name != nil {
		config["name"] = *r.Name
	}
	if r.Ebs_optimized != nil {
		config["ebs_optimized"] = *r.Ebs_optimized
	}
	config["cluster_id"] = r.Cluster_id
	config["instance_type"] = r.Instance_type
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_kms_alias struct {
	Arn            string
	Name           *string
	Name_prefix    *string
	Target_key_id  string
	Target_key_arn string
}

func Aws_kms_aliasMapper(r *Aws_kms_alias) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Name != nil {
		config["name"] = *r.Name
	}
	if r.Name_prefix != nil {
		config["name_prefix"] = *r.Name_prefix
	}
	config["target_key_id"] = r.Target_key_id
	config["target_key_arn"] = r.Target_key_arn
	config["arn"] = r.Arn
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_ses_identity_notification_topic struct {
	Identity          string
	Topic_arn         *string
	Notification_type string
}

func Aws_ses_identity_notification_topicMapper(r *Aws_ses_identity_notification_topic) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Topic_arn != nil {
		config["topic_arn"] = *r.Topic_arn
	}
	config["notification_type"] = r.Notification_type
	config["identity"] = r.Identity
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_api_gateway_method_settings struct {
	Rest_api_id string
	Stage_name  string
	Method_path string
}

func Aws_api_gateway_method_settingsMapper(r *Aws_api_gateway_method_settings) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["rest_api_id"] = r.Rest_api_id
	config["stage_name"] = r.Stage_name
	config["method_path"] = r.Method_path
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_network_interface_attachment struct {
	Instance_id          string
	Network_interface_id string
	Attachment_id        string
	Status               string
}

func Aws_network_interface_attachmentMapper(r *Aws_network_interface_attachment) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["instance_id"] = r.Instance_id
	config["network_interface_id"] = r.Network_interface_id
	config["attachment_id"] = r.Attachment_id
	config["status"] = r.Status
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_customer_gateway struct {
	Ip_address    string
	Resource_type string
	Tags          *map[string]interface{}
}

func Aws_customer_gatewayMapper(r *Aws_customer_gateway) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["ip_address"] = r.Ip_address
	config["resource_type"] = r.Resource_type
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_dx_hosted_public_virtual_interface struct {
	Amazon_address   *string
	Arn              string
	Name             string
	Bgp_auth_key     *string
	Customer_address *string
	Owner_account_id string
	Connection_id    string
	Address_family   string
}

func Aws_dx_hosted_public_virtual_interfaceMapper(r *Aws_dx_hosted_public_virtual_interface) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["connection_id"] = r.Connection_id
	config["address_family"] = r.Address_family
	config["owner_account_id"] = r.Owner_account_id
	config["arn"] = r.Arn
	config["name"] = r.Name
	if r.Bgp_auth_key != nil {
		config["bgp_auth_key"] = *r.Bgp_auth_key
	}
	if r.Customer_address != nil {
		config["customer_address"] = *r.Customer_address
	}
	if r.Amazon_address != nil {
		config["amazon_address"] = *r.Amazon_address
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_storagegateway_nfs_file_share struct {
	Read_only               *bool
	Role_arn                string
	Object_acl              *string
	Requester_pays          *bool
	Arn                     string
	Kms_encrypted           *bool
	Guess_mime_type_enabled *bool
	Kms_key_arn             *string
	Location_arn            string
	Default_storage_class   *string
	Gateway_arn             string
	Squash                  *string
	Fileshare_id            string
}

func Aws_storagegateway_nfs_file_shareMapper(r *Aws_storagegateway_nfs_file_share) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Default_storage_class != nil {
		config["default_storage_class"] = *r.Default_storage_class
	}
	config["gateway_arn"] = r.Gateway_arn
	if r.Guess_mime_type_enabled != nil {
		config["guess_mime_type_enabled"] = *r.Guess_mime_type_enabled
	}
	if r.Kms_key_arn != nil {
		config["kms_key_arn"] = *r.Kms_key_arn
	}
	config["location_arn"] = r.Location_arn
	config["fileshare_id"] = r.Fileshare_id
	if r.Squash != nil {
		config["squash"] = *r.Squash
	}
	if r.Object_acl != nil {
		config["object_acl"] = *r.Object_acl
	}
	if r.Read_only != nil {
		config["read_only"] = *r.Read_only
	}
	config["role_arn"] = r.Role_arn
	config["arn"] = r.Arn
	if r.Kms_encrypted != nil {
		config["kms_encrypted"] = *r.Kms_encrypted
	}
	if r.Requester_pays != nil {
		config["requester_pays"] = *r.Requester_pays
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_waf_rule_group struct {
	Name        string
	Metric_name string
}

func Aws_waf_rule_groupMapper(r *Aws_waf_rule_group) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	config["metric_name"] = r.Metric_name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_pinpoint_sms_channel struct {
	Enabled        *bool
	Sender_id      *string
	Short_code     *string
	Application_id string
}

func Aws_pinpoint_sms_channelMapper(r *Aws_pinpoint_sms_channel) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["application_id"] = r.Application_id
	if r.Enabled != nil {
		config["enabled"] = *r.Enabled
	}
	if r.Sender_id != nil {
		config["sender_id"] = *r.Sender_id
	}
	if r.Short_code != nil {
		config["short_code"] = *r.Short_code
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_iam_user_ssh_key struct {
	Public_key        string
	Encoding          string
	Status            *string
	Ssh_public_key_id string
	Fingerprint       string
	Username          string
}

func Aws_iam_user_ssh_keyMapper(r *Aws_iam_user_ssh_key) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["public_key"] = r.Public_key
	config["encoding"] = r.Encoding
	if r.Status != nil {
		config["status"] = *r.Status
	}
	config["ssh_public_key_id"] = r.Ssh_public_key_id
	config["fingerprint"] = r.Fingerprint
	config["username"] = r.Username
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_instance struct {
	Subnet_id                            *string
	Monitoring                           *bool
	Arn                                  string
	Instance_type                        string
	Tags                                 *map[string]interface{}
	Tenancy                              *string
	Host_id                              *string
	Instance_state                       string
	Availability_zone                    *string
	Network_interface_id                 string
	Instance_initiated_shutdown_behavior *string
	Volume_tags                          *map[string]interface{}
	Ami                                  string
	Password_data                        string
	User_data                            *string
	Public_ip                            string
	Disable_api_termination              *bool
	Key_name                             *string
	Get_password_data                    *bool
	Block_device                         *map[string]interface{}
	Private_dns                          string
	Iam_instance_profile                 *string
	Placement_group                      *string
	Source_dest_check                    *bool
	Public_dns                           string
	Primary_network_interface_id         string
	Ebs_optimized                        *bool
	Private_ip                           *string
	User_data_base64                     *string
	Associate_public_ip_address          *bool
}

func Aws_instanceMapper(r *Aws_instance) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Placement_group != nil {
		config["placement_group"] = *r.Placement_group
	}
	if r.Source_dest_check != nil {
		config["source_dest_check"] = *r.Source_dest_check
	}
	config["private_dns"] = r.Private_dns
	if r.Iam_instance_profile != nil {
		config["iam_instance_profile"] = *r.Iam_instance_profile
	}
	if r.Private_ip != nil {
		config["private_ip"] = *r.Private_ip
	}
	if r.User_data_base64 != nil {
		config["user_data_base64"] = *r.User_data_base64
	}
	config["public_dns"] = r.Public_dns
	config["primary_network_interface_id"] = r.Primary_network_interface_id
	if r.Ebs_optimized != nil {
		config["ebs_optimized"] = *r.Ebs_optimized
	}
	if r.Associate_public_ip_address != nil {
		config["associate_public_ip_address"] = *r.Associate_public_ip_address
	}
	config["arn"] = r.Arn
	config["instance_type"] = r.Instance_type
	if r.Subnet_id != nil {
		config["subnet_id"] = *r.Subnet_id
	}
	if r.Monitoring != nil {
		config["monitoring"] = *r.Monitoring
	}
	if r.Tenancy != nil {
		config["tenancy"] = *r.Tenancy
	}
	if r.Host_id != nil {
		config["host_id"] = *r.Host_id
	}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	if r.Availability_zone != nil {
		config["availability_zone"] = *r.Availability_zone
	}
	config["network_interface_id"] = r.Network_interface_id
	config["instance_state"] = r.Instance_state
	config["ami"] = r.Ami
	config["password_data"] = r.Password_data
	if r.Instance_initiated_shutdown_behavior != nil {
		config["instance_initiated_shutdown_behavior"] = *r.Instance_initiated_shutdown_behavior
	}
	if r.Volume_tags != nil {
		config["volume_tags"] = *r.Volume_tags
	}
	if r.Key_name != nil {
		config["key_name"] = *r.Key_name
	}
	if r.Get_password_data != nil {
		config["get_password_data"] = *r.Get_password_data
	}
	if r.User_data != nil {
		config["user_data"] = *r.User_data
	}
	config["public_ip"] = r.Public_ip
	if r.Disable_api_termination != nil {
		config["disable_api_termination"] = *r.Disable_api_termination
	}
	if r.Block_device != nil {
		config["block_device"] = *r.Block_device
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_neptune_event_subscription struct {
	Customer_aws_id string
	Arn             string
	Name            *string
	Name_prefix     *string
	Sns_topic_arn   string
	Source_type     *string
	Enabled         *bool
	Tags            *map[string]interface{}
}

func Aws_neptune_event_subscriptionMapper(r *Aws_neptune_event_subscription) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["arn"] = r.Arn
	if r.Name != nil {
		config["name"] = *r.Name
	}
	if r.Name_prefix != nil {
		config["name_prefix"] = *r.Name_prefix
	}
	config["sns_topic_arn"] = r.Sns_topic_arn
	if r.Source_type != nil {
		config["source_type"] = *r.Source_type
	}
	if r.Enabled != nil {
		config["enabled"] = *r.Enabled
	}
	config["customer_aws_id"] = r.Customer_aws_id
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_opsworks_rds_db_instance struct {
	Stack_id            string
	Rds_db_instance_arn string
	Db_password         string
	Db_user             string
}

func Aws_opsworks_rds_db_instanceMapper(r *Aws_opsworks_rds_db_instance) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["stack_id"] = r.Stack_id
	config["rds_db_instance_arn"] = r.Rds_db_instance_arn
	config["db_password"] = r.Db_password
	config["db_user"] = r.Db_user
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_cloudwatch_event_target struct {
	Input      *string
	Input_path *string
	Rule       string
	Target_id  *string
	Arn        string
	Role_arn   *string
}

func Aws_cloudwatch_event_targetMapper(r *Aws_cloudwatch_event_target) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Role_arn != nil {
		config["role_arn"] = *r.Role_arn
	}
	config["rule"] = r.Rule
	if r.Target_id != nil {
		config["target_id"] = *r.Target_id
	}
	config["arn"] = r.Arn
	if r.Input != nil {
		config["input"] = *r.Input
	}
	if r.Input_path != nil {
		config["input_path"] = *r.Input_path
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_dlm_lifecycle_policy struct {
	Description        string
	Execution_role_arn string
	State              *string
}

func Aws_dlm_lifecycle_policyMapper(r *Aws_dlm_lifecycle_policy) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["execution_role_arn"] = r.Execution_role_arn
	if r.State != nil {
		config["state"] = *r.State
	}
	config["description"] = r.Description
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_iot_policy struct {
	Default_version_id string
	Name               string
	Policy             string
	Arn                string
}

func Aws_iot_policyMapper(r *Aws_iot_policy) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	config["policy"] = r.Policy
	config["arn"] = r.Arn
	config["default_version_id"] = r.Default_version_id
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_macie_s3_bucket_association struct {
	Bucket_name       string
	Prefix            *string
	Member_account_id *string
}

func Aws_macie_s3_bucket_associationMapper(r *Aws_macie_s3_bucket_association) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["bucket_name"] = r.Bucket_name
	if r.Prefix != nil {
		config["prefix"] = *r.Prefix
	}
	if r.Member_account_id != nil {
		config["member_account_id"] = *r.Member_account_id
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_opsworks_java_app_layer struct {
	Drain_elb_on_shutdown       *bool
	Stack_id                    string
	Jvm_version                 *string
	Name                        *string
	App_server_version          *string
	Auto_assign_public_ips      *bool
	Install_updates_on_boot     *bool
	Jvm_type                    *string
	Use_ebs_optimized_instances *bool
	Jvm_options                 *string
	App_server                  *string
	Elastic_load_balancer       *string
	Custom_json                 *string
	Auto_healing                *bool
	Auto_assign_elastic_ips     *bool
	Custom_instance_profile_arn *string
}

func Aws_opsworks_java_app_layerMapper(r *Aws_opsworks_java_app_layer) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Install_updates_on_boot != nil {
		config["install_updates_on_boot"] = *r.Install_updates_on_boot
	}
	if r.Name != nil {
		config["name"] = *r.Name
	}
	if r.App_server_version != nil {
		config["app_server_version"] = *r.App_server_version
	}
	if r.Auto_assign_public_ips != nil {
		config["auto_assign_public_ips"] = *r.Auto_assign_public_ips
	}
	if r.Jvm_type != nil {
		config["jvm_type"] = *r.Jvm_type
	}
	if r.Custom_json != nil {
		config["custom_json"] = *r.Custom_json
	}
	if r.Auto_healing != nil {
		config["auto_healing"] = *r.Auto_healing
	}
	if r.Use_ebs_optimized_instances != nil {
		config["use_ebs_optimized_instances"] = *r.Use_ebs_optimized_instances
	}
	if r.Jvm_options != nil {
		config["jvm_options"] = *r.Jvm_options
	}
	if r.App_server != nil {
		config["app_server"] = *r.App_server
	}
	if r.Elastic_load_balancer != nil {
		config["elastic_load_balancer"] = *r.Elastic_load_balancer
	}
	if r.Auto_assign_elastic_ips != nil {
		config["auto_assign_elastic_ips"] = *r.Auto_assign_elastic_ips
	}
	if r.Custom_instance_profile_arn != nil {
		config["custom_instance_profile_arn"] = *r.Custom_instance_profile_arn
	}
	if r.Drain_elb_on_shutdown != nil {
		config["drain_elb_on_shutdown"] = *r.Drain_elb_on_shutdown
	}
	config["stack_id"] = r.Stack_id
	if r.Jvm_version != nil {
		config["jvm_version"] = *r.Jvm_version
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_ssm_patch_baseline struct {
	Operating_system                  *string
	Approved_patches_compliance_level *string
	Name                              string
	Description                       *string
}

func Aws_ssm_patch_baselineMapper(r *Aws_ssm_patch_baseline) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Operating_system != nil {
		config["operating_system"] = *r.Operating_system
	}
	if r.Approved_patches_compliance_level != nil {
		config["approved_patches_compliance_level"] = *r.Approved_patches_compliance_level
	}
	config["name"] = r.Name
	if r.Description != nil {
		config["description"] = *r.Description
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_vpc_endpoint_connection_notification struct {
	Connection_notification_arn string
	State                       string
	Notification_type           string
	Vpc_endpoint_service_id     *string
	Vpc_endpoint_id             *string
}

func Aws_vpc_endpoint_connection_notificationMapper(r *Aws_vpc_endpoint_connection_notification) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Vpc_endpoint_service_id != nil {
		config["vpc_endpoint_service_id"] = *r.Vpc_endpoint_service_id
	}
	if r.Vpc_endpoint_id != nil {
		config["vpc_endpoint_id"] = *r.Vpc_endpoint_id
	}
	config["connection_notification_arn"] = r.Connection_notification_arn
	config["state"] = r.State
	config["notification_type"] = r.Notification_type
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_api_gateway_documentation_version struct {
	Version     string
	Rest_api_id string
	Description *string
}

func Aws_api_gateway_documentation_versionMapper(r *Aws_api_gateway_documentation_version) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Description != nil {
		config["description"] = *r.Description
	}
	config["version"] = r.Version
	config["rest_api_id"] = r.Rest_api_id
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_ec2_transit_gateway_vpc_attachment struct {
	Dns_support                                     *string
	Transit_gateway_id                              string
	Vpc_id                                          string
	Ipv6_support                                    *string
	Tags                                            *map[string]interface{}
	Transit_gateway_default_route_table_association *bool
	Transit_gateway_default_route_table_propagation *bool
	Vpc_owner_id                                    string
}

func Aws_ec2_transit_gateway_vpc_attachmentMapper(r *Aws_ec2_transit_gateway_vpc_attachment) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Dns_support != nil {
		config["dns_support"] = *r.Dns_support
	}
	config["transit_gateway_id"] = r.Transit_gateway_id
	config["vpc_id"] = r.Vpc_id
	config["vpc_owner_id"] = r.Vpc_owner_id
	if r.Ipv6_support != nil {
		config["ipv6_support"] = *r.Ipv6_support
	}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	if r.Transit_gateway_default_route_table_association != nil {
		config["transit_gateway_default_route_table_association"] = *r.Transit_gateway_default_route_table_association
	}
	if r.Transit_gateway_default_route_table_propagation != nil {
		config["transit_gateway_default_route_table_propagation"] = *r.Transit_gateway_default_route_table_propagation
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_egress_only_internet_gateway struct {
	Vpc_id string
}

func Aws_egress_only_internet_gatewayMapper(r *Aws_egress_only_internet_gateway) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["vpc_id"] = r.Vpc_id
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_kms_grant struct {
	Key_id             string
	Grant_token        string
	Name               *string
	Grantee_principal  string
	Retiring_principal *string
	Retire_on_delete   *bool
	Grant_id           string
}

func Aws_kms_grantMapper(r *Aws_kms_grant) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["key_id"] = r.Key_id
	config["grant_token"] = r.Grant_token
	if r.Name != nil {
		config["name"] = *r.Name
	}
	config["grantee_principal"] = r.Grantee_principal
	if r.Retiring_principal != nil {
		config["retiring_principal"] = *r.Retiring_principal
	}
	if r.Retire_on_delete != nil {
		config["retire_on_delete"] = *r.Retire_on_delete
	}
	config["grant_id"] = r.Grant_id
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_route53_zone_association struct {
	Zone_id    string
	Vpc_id     string
	Vpc_region *string
}

func Aws_route53_zone_associationMapper(r *Aws_route53_zone_association) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["zone_id"] = r.Zone_id
	config["vpc_id"] = r.Vpc_id
	if r.Vpc_region != nil {
		config["vpc_region"] = *r.Vpc_region
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_mq_configuration struct {
	Description    *string
	Engine_type    string
	Engine_version string
	Name           string
	Arn            string
	Data           string
}

func Aws_mq_configurationMapper(r *Aws_mq_configuration) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["data"] = r.Data
	if r.Description != nil {
		config["description"] = *r.Description
	}
	config["engine_type"] = r.Engine_type
	config["engine_version"] = r.Engine_version
	config["name"] = r.Name
	config["arn"] = r.Arn
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_ssm_document struct {
	Default_version string
	Status          string
	Content         string
	Document_type   string
	Description     string
	Hash_type       string
	Latest_version  string
	Owner           string
	Tags            *map[string]interface{}
	Arn             string
	Created_date    string
	Name            string
	Document_format *string
	Schema_version  string
	Hash            string
	Permissions     *map[string]interface{}
}

func Aws_ssm_documentMapper(r *Aws_ssm_document) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["arn"] = r.Arn
	config["created_date"] = r.Created_date
	config["name"] = r.Name
	if r.Document_format != nil {
		config["document_format"] = *r.Document_format
	}
	config["schema_version"] = r.Schema_version
	config["hash"] = r.Hash
	if r.Permissions != nil {
		config["permissions"] = *r.Permissions
	}
	config["default_version"] = r.Default_version
	config["status"] = r.Status
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["content"] = r.Content
	config["document_type"] = r.Document_type
	config["description"] = r.Description
	config["hash_type"] = r.Hash_type
	config["latest_version"] = r.Latest_version
	config["owner"] = r.Owner
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_athena_database struct {
	Name          string
	Bucket        string
	Force_destroy *bool
}

func Aws_athena_databaseMapper(r *Aws_athena_database) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["bucket"] = r.Bucket
	if r.Force_destroy != nil {
		config["force_destroy"] = *r.Force_destroy
	}
	config["name"] = r.Name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_cloudwatch_log_resource_policy struct {
	Policy_document string
	Policy_name     string
}

func Aws_cloudwatch_log_resource_policyMapper(r *Aws_cloudwatch_log_resource_policy) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["policy_name"] = r.Policy_name
	config["policy_document"] = r.Policy_document
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_ssm_maintenance_window struct {
	Name                       string
	Allow_unassociated_targets *bool
	Enabled                    *bool
	End_date                   *string
	Start_date                 *string
	Schedule                   string
	Schedule_timezone          *string
}

func Aws_ssm_maintenance_windowMapper(r *Aws_ssm_maintenance_window) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["schedule"] = r.Schedule
	if r.Schedule_timezone != nil {
		config["schedule_timezone"] = *r.Schedule_timezone
	}
	config["name"] = r.Name
	if r.Allow_unassociated_targets != nil {
		config["allow_unassociated_targets"] = *r.Allow_unassociated_targets
	}
	if r.Enabled != nil {
		config["enabled"] = *r.Enabled
	}
	if r.End_date != nil {
		config["end_date"] = *r.End_date
	}
	if r.Start_date != nil {
		config["start_date"] = *r.Start_date
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_waf_rule struct {
	Name        string
	Metric_name string
}

func Aws_waf_ruleMapper(r *Aws_waf_rule) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	config["metric_name"] = r.Metric_name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_iam_user_group_membership struct {
	User string
}

func Aws_iam_user_group_membershipMapper(r *Aws_iam_user_group_membership) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["user"] = r.User
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_lb_listener_certificate struct {
	Listener_arn    string
	Certificate_arn string
}

func Aws_lb_listener_certificateMapper(r *Aws_lb_listener_certificate) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["listener_arn"] = r.Listener_arn
	config["certificate_arn"] = r.Certificate_arn
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_vpn_gateway struct {
	Amazon_side_asn   *string
	Vpc_id            *string
	Tags              *map[string]interface{}
	Availability_zone *string
}

func Aws_vpn_gatewayMapper(r *Aws_vpn_gateway) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	if r.Availability_zone != nil {
		config["availability_zone"] = *r.Availability_zone
	}
	if r.Amazon_side_asn != nil {
		config["amazon_side_asn"] = *r.Amazon_side_asn
	}
	if r.Vpc_id != nil {
		config["vpc_id"] = *r.Vpc_id
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_wafregional_size_constraint_set struct {
	Name string
}

func Aws_wafregional_size_constraint_setMapper(r *Aws_wafregional_size_constraint_set) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_pinpoint_app struct {
	Name_prefix    *string
	Application_id string
	Name           *string
}

func Aws_pinpoint_appMapper(r *Aws_pinpoint_app) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["application_id"] = r.Application_id
	if r.Name != nil {
		config["name"] = *r.Name
	}
	if r.Name_prefix != nil {
		config["name_prefix"] = *r.Name_prefix
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_config_config_rule struct {
	Name                        string
	Rule_id                     string
	Arn                         string
	Description                 *string
	Input_parameters            *string
	Maximum_execution_frequency *string
}

func Aws_config_config_ruleMapper(r *Aws_config_config_rule) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Description != nil {
		config["description"] = *r.Description
	}
	if r.Input_parameters != nil {
		config["input_parameters"] = *r.Input_parameters
	}
	if r.Maximum_execution_frequency != nil {
		config["maximum_execution_frequency"] = *r.Maximum_execution_frequency
	}
	config["name"] = r.Name
	config["rule_id"] = r.Rule_id
	config["arn"] = r.Arn
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_devicefarm_project struct {
	Arn  string
	Name string
}

func Aws_devicefarm_projectMapper(r *Aws_devicefarm_project) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["arn"] = r.Arn
	config["name"] = r.Name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_elasticache_parameter_group struct {
	Family      string
	Description *string
	Name        string
}

func Aws_elasticache_parameter_groupMapper(r *Aws_elasticache_parameter_group) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	config["family"] = r.Family
	if r.Description != nil {
		config["description"] = *r.Description
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_resourcegroups_group struct {
	Arn         string
	Name        string
	Description *string
}

func Aws_resourcegroups_groupMapper(r *Aws_resourcegroups_group) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	if r.Description != nil {
		config["description"] = *r.Description
	}
	config["arn"] = r.Arn
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_ses_domain_dkim struct {
	Domain string
}

func Aws_ses_domain_dkimMapper(r *Aws_ses_domain_dkim) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["domain"] = r.Domain
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_vpc_dhcp_options struct {
	Tags              *map[string]interface{}
	Owner_id          string
	Domain_name       *string
	Netbios_node_type *string
}

func Aws_vpc_dhcp_optionsMapper(r *Aws_vpc_dhcp_options) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["owner_id"] = r.Owner_id
	if r.Domain_name != nil {
		config["domain_name"] = *r.Domain_name
	}
	if r.Netbios_node_type != nil {
		config["netbios_node_type"] = *r.Netbios_node_type
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_appsync_datasource struct {
	Name             string
	Service_role_arn *string
	Arn              string
	Api_id           string
	Resource_type    string
	Description      *string
}

func Aws_appsync_datasourceMapper(r *Aws_appsync_datasource) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["resource_type"] = r.Resource_type
	if r.Description != nil {
		config["description"] = *r.Description
	}
	if r.Service_role_arn != nil {
		config["service_role_arn"] = *r.Service_role_arn
	}
	config["arn"] = r.Arn
	config["api_id"] = r.Api_id
	config["name"] = r.Name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_codedeploy_deployment_config struct {
	Deployment_config_id   string
	Deployment_config_name string
	Compute_platform       *string
}

func Aws_codedeploy_deployment_configMapper(r *Aws_codedeploy_deployment_config) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["deployment_config_name"] = r.Deployment_config_name
	if r.Compute_platform != nil {
		config["compute_platform"] = *r.Compute_platform
	}
	config["deployment_config_id"] = r.Deployment_config_id
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_datasync_location_s3 struct {
	Uri           string
	Arn           string
	S3_bucket_arn string
	Subdirectory  string
	Tags          *map[string]interface{}
}

func Aws_datasync_location_s3Mapper(r *Aws_datasync_location_s3) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["uri"] = r.Uri
	config["arn"] = r.Arn
	config["s3_bucket_arn"] = r.S3_bucket_arn
	config["subdirectory"] = r.Subdirectory
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_route_table struct {
	Owner_id string
	Vpc_id   string
	Tags     *map[string]interface{}
}

func Aws_route_tableMapper(r *Aws_route_table) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["vpc_id"] = r.Vpc_id
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["owner_id"] = r.Owner_id
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_vpc_peering_connection_accepter struct {
	Tags                      *map[string]interface{}
	Auto_accept               *bool
	Accept_status             string
	Peer_vpc_id               string
	Peer_owner_id             string
	Peer_region               string
	Vpc_peering_connection_id string
	Vpc_id                    string
}

func Aws_vpc_peering_connection_accepterMapper(r *Aws_vpc_peering_connection_accepter) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["peer_region"] = r.Peer_region
	config["vpc_peering_connection_id"] = r.Vpc_peering_connection_id
	config["vpc_id"] = r.Vpc_id
	config["peer_vpc_id"] = r.Peer_vpc_id
	config["peer_owner_id"] = r.Peer_owner_id
	if r.Auto_accept != nil {
		config["auto_accept"] = *r.Auto_accept
	}
	config["accept_status"] = r.Accept_status
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_default_subnet struct {
	Ipv6_cidr_block_association_id  string
	Tags                            *map[string]interface{}
	Owner_id                        string
	Ipv6_cidr_block                 string
	Map_public_ip_on_launch         *bool
	Availability_zone               string
	Availability_zone_id            string
	Assign_ipv6_address_on_creation bool
	Arn                             string
	Vpc_id                          string
	Cidr_block                      string
}

func Aws_default_subnetMapper(r *Aws_default_subnet) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["owner_id"] = r.Owner_id
	config["ipv6_cidr_block"] = r.Ipv6_cidr_block
	if r.Map_public_ip_on_launch != nil {
		config["map_public_ip_on_launch"] = *r.Map_public_ip_on_launch
	}
	config["ipv6_cidr_block_association_id"] = r.Ipv6_cidr_block_association_id
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["assign_ipv6_address_on_creation"] = r.Assign_ipv6_address_on_creation
	config["arn"] = r.Arn
	config["vpc_id"] = r.Vpc_id
	config["cidr_block"] = r.Cidr_block
	config["availability_zone"] = r.Availability_zone
	config["availability_zone_id"] = r.Availability_zone_id
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_codecommit_trigger struct {
	Repository_name  string
	Configuration_id string
}

func Aws_codecommit_triggerMapper(r *Aws_codecommit_trigger) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["repository_name"] = r.Repository_name
	config["configuration_id"] = r.Configuration_id
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_db_event_subscription struct {
	Source_type     *string
	Name_prefix     *string
	Sns_topic       string
	Enabled         *bool
	Customer_aws_id string
	Tags            *map[string]interface{}
	Arn             string
	Name            *string
}

func Aws_db_event_subscriptionMapper(r *Aws_db_event_subscription) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Enabled != nil {
		config["enabled"] = *r.Enabled
	}
	config["customer_aws_id"] = r.Customer_aws_id
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["arn"] = r.Arn
	if r.Name != nil {
		config["name"] = *r.Name
	}
	if r.Source_type != nil {
		config["source_type"] = *r.Source_type
	}
	if r.Name_prefix != nil {
		config["name_prefix"] = *r.Name_prefix
	}
	config["sns_topic"] = r.Sns_topic
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_neptune_cluster struct {
	Engine                               *string
	Hosted_zone_id                       string
	Preferred_maintenance_window         *string
	Tags                                 *map[string]interface{}
	Arn                                  string
	Replication_source_identifier        *string
	Cluster_identifier                   *string
	Endpoint                             string
	Snapshot_identifier                  *string
	Apply_immediately                    *bool
	Iam_database_authentication_enabled  *bool
	Reader_endpoint                      string
	Storage_encrypted                    *bool
	Skip_final_snapshot                  *bool
	Final_snapshot_identifier            *string
	Neptune_subnet_group_name            *string
	Preferred_backup_window              *string
	Cluster_identifier_prefix            *string
	Cluster_resource_id                  string
	Engine_version                       *string
	Kms_key_arn                          *string
	Neptune_cluster_parameter_group_name *string
}

func Aws_neptune_clusterMapper(r *Aws_neptune_cluster) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Final_snapshot_identifier != nil {
		config["final_snapshot_identifier"] = *r.Final_snapshot_identifier
	}
	if r.Neptune_subnet_group_name != nil {
		config["neptune_subnet_group_name"] = *r.Neptune_subnet_group_name
	}
	if r.Engine_version != nil {
		config["engine_version"] = *r.Engine_version
	}
	if r.Kms_key_arn != nil {
		config["kms_key_arn"] = *r.Kms_key_arn
	}
	if r.Neptune_cluster_parameter_group_name != nil {
		config["neptune_cluster_parameter_group_name"] = *r.Neptune_cluster_parameter_group_name
	}
	if r.Preferred_backup_window != nil {
		config["preferred_backup_window"] = *r.Preferred_backup_window
	}
	if r.Cluster_identifier_prefix != nil {
		config["cluster_identifier_prefix"] = *r.Cluster_identifier_prefix
	}
	config["cluster_resource_id"] = r.Cluster_resource_id
	if r.Preferred_maintenance_window != nil {
		config["preferred_maintenance_window"] = *r.Preferred_maintenance_window
	}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	if r.Engine != nil {
		config["engine"] = *r.Engine
	}
	config["hosted_zone_id"] = r.Hosted_zone_id
	if r.Replication_source_identifier != nil {
		config["replication_source_identifier"] = *r.Replication_source_identifier
	}
	config["arn"] = r.Arn
	config["endpoint"] = r.Endpoint
	if r.Cluster_identifier != nil {
		config["cluster_identifier"] = *r.Cluster_identifier
	}
	if r.Snapshot_identifier != nil {
		config["snapshot_identifier"] = *r.Snapshot_identifier
	}
	if r.Apply_immediately != nil {
		config["apply_immediately"] = *r.Apply_immediately
	}
	if r.Iam_database_authentication_enabled != nil {
		config["iam_database_authentication_enabled"] = *r.Iam_database_authentication_enabled
	}
	if r.Skip_final_snapshot != nil {
		config["skip_final_snapshot"] = *r.Skip_final_snapshot
	}
	config["reader_endpoint"] = r.Reader_endpoint
	if r.Storage_encrypted != nil {
		config["storage_encrypted"] = *r.Storage_encrypted
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_route53_zone struct {
	Vpc_id            *string
	Zone_id           string
	Tags              *map[string]interface{}
	Force_destroy     *bool
	Name              string
	Comment           *string
	Vpc_region        *string
	Delegation_set_id *string
}

func Aws_route53_zoneMapper(r *Aws_route53_zone) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Vpc_id != nil {
		config["vpc_id"] = *r.Vpc_id
	}
	config["zone_id"] = r.Zone_id
	if r.Vpc_region != nil {
		config["vpc_region"] = *r.Vpc_region
	}
	if r.Delegation_set_id != nil {
		config["delegation_set_id"] = *r.Delegation_set_id
	}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	if r.Force_destroy != nil {
		config["force_destroy"] = *r.Force_destroy
	}
	config["name"] = r.Name
	if r.Comment != nil {
		config["comment"] = *r.Comment
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_securityhub_account struct {
}

func Aws_securityhub_accountMapper(r *Aws_securityhub_account) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_cloudwatch_event_permission struct {
	Action       *string
	Principal    string
	Statement_id string
}

func Aws_cloudwatch_event_permissionMapper(r *Aws_cloudwatch_event_permission) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["principal"] = r.Principal
	config["statement_id"] = r.Statement_id
	if r.Action != nil {
		config["action"] = *r.Action
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_cloudwatch_log_group struct {
	Name        *string
	Name_prefix *string
	Kms_key_id  *string
	Arn         string
	Tags        *map[string]interface{}
}

func Aws_cloudwatch_log_groupMapper(r *Aws_cloudwatch_log_group) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Name != nil {
		config["name"] = *r.Name
	}
	if r.Name_prefix != nil {
		config["name_prefix"] = *r.Name_prefix
	}
	if r.Kms_key_id != nil {
		config["kms_key_id"] = *r.Kms_key_id
	}
	config["arn"] = r.Arn
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_docdb_subnet_group struct {
	Description *string
	Tags        *map[string]interface{}
	Arn         string
	Name        *string
	Name_prefix *string
}

func Aws_docdb_subnet_groupMapper(r *Aws_docdb_subnet_group) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["arn"] = r.Arn
	if r.Name != nil {
		config["name"] = *r.Name
	}
	if r.Name_prefix != nil {
		config["name_prefix"] = *r.Name_prefix
	}
	if r.Description != nil {
		config["description"] = *r.Description
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_dx_gateway struct {
	Name            string
	Amazon_side_asn string
}

func Aws_dx_gatewayMapper(r *Aws_dx_gateway) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["amazon_side_asn"] = r.Amazon_side_asn
	config["name"] = r.Name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_wafregional_sql_injection_match_set struct {
	Name string
}

func Aws_wafregional_sql_injection_match_setMapper(r *Aws_wafregional_sql_injection_match_set) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_autoscaling_attachment struct {
	Autoscaling_group_name string
	Elb                    *string
	Alb_target_group_arn   *string
}

func Aws_autoscaling_attachmentMapper(r *Aws_autoscaling_attachment) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["autoscaling_group_name"] = r.Autoscaling_group_name
	if r.Elb != nil {
		config["elb"] = *r.Elb
	}
	if r.Alb_target_group_arn != nil {
		config["alb_target_group_arn"] = *r.Alb_target_group_arn
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_iam_access_key struct {
	Status            string
	Secret            string
	Ses_smtp_password string
	Pgp_key           *string
	Key_fingerprint   string
	Encrypted_secret  string
	User              string
}

func Aws_iam_access_keyMapper(r *Aws_iam_access_key) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["secret"] = r.Secret
	config["ses_smtp_password"] = r.Ses_smtp_password
	if r.Pgp_key != nil {
		config["pgp_key"] = *r.Pgp_key
	}
	config["key_fingerprint"] = r.Key_fingerprint
	config["encrypted_secret"] = r.Encrypted_secret
	config["user"] = r.User
	config["status"] = r.Status
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_vpc_peering_connection_options struct {
	Vpc_peering_connection_id string
}

func Aws_vpc_peering_connection_optionsMapper(r *Aws_vpc_peering_connection_options) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["vpc_peering_connection_id"] = r.Vpc_peering_connection_id
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_datasync_location_nfs struct {
	Arn             string
	Server_hostname string
	Subdirectory    string
	Tags            *map[string]interface{}
	Uri             string
}

func Aws_datasync_location_nfsMapper(r *Aws_datasync_location_nfs) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["arn"] = r.Arn
	config["server_hostname"] = r.Server_hostname
	config["subdirectory"] = r.Subdirectory
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["uri"] = r.Uri
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_flow_log struct {
	Log_destination      *string
	Log_destination_type *string
	Log_group_name       *string
	Vpc_id               *string
	Subnet_id            *string
	Eni_id               *string
	Traffic_type         string
	Iam_role_arn         *string
}

func Aws_flow_logMapper(r *Aws_flow_log) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["traffic_type"] = r.Traffic_type
	if r.Iam_role_arn != nil {
		config["iam_role_arn"] = *r.Iam_role_arn
	}
	if r.Log_destination != nil {
		config["log_destination"] = *r.Log_destination
	}
	if r.Log_destination_type != nil {
		config["log_destination_type"] = *r.Log_destination_type
	}
	if r.Log_group_name != nil {
		config["log_group_name"] = *r.Log_group_name
	}
	if r.Vpc_id != nil {
		config["vpc_id"] = *r.Vpc_id
	}
	if r.Subnet_id != nil {
		config["subnet_id"] = *r.Subnet_id
	}
	if r.Eni_id != nil {
		config["eni_id"] = *r.Eni_id
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_iot_thing_type struct {
	Name       string
	Deprecated *bool
	Arn        string
}

func Aws_iot_thing_typeMapper(r *Aws_iot_thing_type) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	if r.Deprecated != nil {
		config["deprecated"] = *r.Deprecated
	}
	config["arn"] = r.Arn
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_pinpoint_gcm_channel struct {
	Application_id string
	Api_key        string
	Enabled        *bool
}

func Aws_pinpoint_gcm_channelMapper(r *Aws_pinpoint_gcm_channel) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["application_id"] = r.Application_id
	config["api_key"] = r.Api_key
	if r.Enabled != nil {
		config["enabled"] = *r.Enabled
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_cloudwatch_metric_alarm struct {
	Alarm_name                            string
	Extended_statistic                    *string
	Evaluate_low_sample_count_percentiles *string
	Arn                                   string
	Comparison_operator                   string
	Metric_name                           string
	Statistic                             *string
	Namespace                             string
	Dimensions                            *map[string]interface{}
	Unit                                  *string
	Treat_missing_data                    *string
	Actions_enabled                       *bool
	Alarm_description                     *string
}

func Aws_cloudwatch_metric_alarmMapper(r *Aws_cloudwatch_metric_alarm) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["namespace"] = r.Namespace
	if r.Dimensions != nil {
		config["dimensions"] = *r.Dimensions
	}
	if r.Actions_enabled != nil {
		config["actions_enabled"] = *r.Actions_enabled
	}
	if r.Alarm_description != nil {
		config["alarm_description"] = *r.Alarm_description
	}
	if r.Unit != nil {
		config["unit"] = *r.Unit
	}
	if r.Treat_missing_data != nil {
		config["treat_missing_data"] = *r.Treat_missing_data
	}
	config["alarm_name"] = r.Alarm_name
	if r.Extended_statistic != nil {
		config["extended_statistic"] = *r.Extended_statistic
	}
	if r.Evaluate_low_sample_count_percentiles != nil {
		config["evaluate_low_sample_count_percentiles"] = *r.Evaluate_low_sample_count_percentiles
	}
	config["arn"] = r.Arn
	config["comparison_operator"] = r.Comparison_operator
	config["metric_name"] = r.Metric_name
	if r.Statistic != nil {
		config["statistic"] = *r.Statistic
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_glue_classifier struct {
	Name string
}

func Aws_glue_classifierMapper(r *Aws_glue_classifier) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_ssm_resource_data_sync struct {
	Name string
}

func Aws_ssm_resource_data_syncMapper(r *Aws_ssm_resource_data_sync) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_default_vpc_dhcp_options struct {
	Owner_id            string
	Domain_name         string
	Domain_name_servers string
	Ntp_servers         string
	Netbios_node_type   *string
	Tags                *map[string]interface{}
}

func Aws_default_vpc_dhcp_optionsMapper(r *Aws_default_vpc_dhcp_options) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["domain_name"] = r.Domain_name
	config["domain_name_servers"] = r.Domain_name_servers
	config["ntp_servers"] = r.Ntp_servers
	if r.Netbios_node_type != nil {
		config["netbios_node_type"] = *r.Netbios_node_type
	}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["owner_id"] = r.Owner_id
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_vpc_endpoint_service struct {
	Acceptance_required bool
	State               string
	Service_type        string
	Service_name        string
	Private_dns_name    string
}

func Aws_vpc_endpoint_serviceMapper(r *Aws_vpc_endpoint_service) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["private_dns_name"] = r.Private_dns_name
	config["service_name"] = r.Service_name
	config["state"] = r.State
	config["service_type"] = r.Service_type
	config["acceptance_required"] = r.Acceptance_required
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_appmesh_virtual_node struct {
	Last_updated_date string
	Name              string
	Mesh_name         string
	Arn               string
	Created_date      string
}

func Aws_appmesh_virtual_nodeMapper(r *Aws_appmesh_virtual_node) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["created_date"] = r.Created_date
	config["last_updated_date"] = r.Last_updated_date
	config["name"] = r.Name
	config["mesh_name"] = r.Mesh_name
	config["arn"] = r.Arn
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_ec2_transit_gateway_route_table struct {
	Default_propagation_route_table bool
	Tags                            *map[string]interface{}
	Transit_gateway_id              string
	Default_association_route_table bool
}

func Aws_ec2_transit_gateway_route_tableMapper(r *Aws_ec2_transit_gateway_route_table) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["default_association_route_table"] = r.Default_association_route_table
	config["default_propagation_route_table"] = r.Default_propagation_route_table
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["transit_gateway_id"] = r.Transit_gateway_id
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_iam_policy_attachment struct {
	Policy_arn string
	Name       string
}

func Aws_iam_policy_attachmentMapper(r *Aws_iam_policy_attachment) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	config["policy_arn"] = r.Policy_arn
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_iam_user_policy struct {
	Policy      string
	Name        *string
	Name_prefix *string
	User        string
}

func Aws_iam_user_policyMapper(r *Aws_iam_user_policy) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["policy"] = r.Policy
	if r.Name != nil {
		config["name"] = *r.Name
	}
	if r.Name_prefix != nil {
		config["name_prefix"] = *r.Name_prefix
	}
	config["user"] = r.User
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_securityhub_standards_subscription struct {
	Standards_arn string
}

func Aws_securityhub_standards_subscriptionMapper(r *Aws_securityhub_standards_subscription) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["standards_arn"] = r.Standards_arn
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_cloudwatch_dashboard struct {
	Dashboard_arn  string
	Dashboard_body string
	Dashboard_name string
}

func Aws_cloudwatch_dashboardMapper(r *Aws_cloudwatch_dashboard) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["dashboard_arn"] = r.Dashboard_arn
	config["dashboard_body"] = r.Dashboard_body
	config["dashboard_name"] = r.Dashboard_name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_network_interface struct {
	Subnet_id         string
	Tags              *map[string]interface{}
	Private_ip        *string
	Private_dns_name  string
	Source_dest_check *bool
	Description       *string
}

func Aws_network_interfaceMapper(r *Aws_network_interface) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	config["subnet_id"] = r.Subnet_id
	config["private_dns_name"] = r.Private_dns_name
	if r.Source_dest_check != nil {
		config["source_dest_check"] = *r.Source_dest_check
	}
	if r.Description != nil {
		config["description"] = *r.Description
	}
	if r.Private_ip != nil {
		config["private_ip"] = *r.Private_ip
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_waf_xss_match_set struct {
	Name string
}

func Aws_waf_xss_match_setMapper(r *Aws_waf_xss_match_set) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_eks_cluster struct {
	Created_at       string
	Version          *string
	Endpoint         string
	Name             string
	Platform_version string
	Role_arn         string
	Arn              string
}

func Aws_eks_clusterMapper(r *Aws_eks_cluster) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["created_at"] = r.Created_at
	if r.Version != nil {
		config["version"] = *r.Version
	}
	config["platform_version"] = r.Platform_version
	config["role_arn"] = r.Role_arn
	config["arn"] = r.Arn
	config["endpoint"] = r.Endpoint
	config["name"] = r.Name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_transfer_user struct {
	Role           string
	Server_id      string
	Tags           *map[string]interface{}
	User_name      string
	Arn            string
	Home_directory *string
	Policy         *string
}

func Aws_transfer_userMapper(r *Aws_transfer_user) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["user_name"] = r.User_name
	config["arn"] = r.Arn
	if r.Home_directory != nil {
		config["home_directory"] = *r.Home_directory
	}
	if r.Policy != nil {
		config["policy"] = *r.Policy
	}
	config["role"] = r.Role
	config["server_id"] = r.Server_id
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_dx_hosted_public_virtual_interface_accepter struct {
	Tags                 *map[string]interface{}
	Arn                  string
	Virtual_interface_id string
}

func Aws_dx_hosted_public_virtual_interface_accepterMapper(r *Aws_dx_hosted_public_virtual_interface_accepter) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["arn"] = r.Arn
	config["virtual_interface_id"] = r.Virtual_interface_id
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_media_store_container_policy struct {
	Container_name string
	Policy         string
}

func Aws_media_store_container_policyMapper(r *Aws_media_store_container_policy) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["container_name"] = r.Container_name
	config["policy"] = r.Policy
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_organizations_organization struct {
	Master_account_email string
	Master_account_id    string
	Feature_set          *string
	Arn                  string
	Master_account_arn   string
}

func Aws_organizations_organizationMapper(r *Aws_organizations_organization) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["arn"] = r.Arn
	config["master_account_arn"] = r.Master_account_arn
	config["master_account_email"] = r.Master_account_email
	config["master_account_id"] = r.Master_account_id
	if r.Feature_set != nil {
		config["feature_set"] = *r.Feature_set
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_waf_web_acl struct {
	Metric_name string
	Name        string
}

func Aws_waf_web_aclMapper(r *Aws_waf_web_acl) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	config["metric_name"] = r.Metric_name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_eip_association struct {
	Instance_id          *string
	Network_interface_id *string
	Private_ip_address   *string
	Public_ip            *string
	Allocation_id        *string
	Allow_reassociation  *bool
}

func Aws_eip_associationMapper(r *Aws_eip_association) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Allocation_id != nil {
		config["allocation_id"] = *r.Allocation_id
	}
	if r.Allow_reassociation != nil {
		config["allow_reassociation"] = *r.Allow_reassociation
	}
	if r.Instance_id != nil {
		config["instance_id"] = *r.Instance_id
	}
	if r.Network_interface_id != nil {
		config["network_interface_id"] = *r.Network_interface_id
	}
	if r.Private_ip_address != nil {
		config["private_ip_address"] = *r.Private_ip_address
	}
	if r.Public_ip != nil {
		config["public_ip"] = *r.Public_ip
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_opsworks_custom_layer struct {
	Drain_elb_on_shutdown       *bool
	Use_ebs_optimized_instances *bool
	Install_updates_on_boot     *bool
	Stack_id                    string
	Short_name                  string
	Name                        string
	Auto_healing                *bool
	Custom_json                 *string
	Elastic_load_balancer       *string
	Auto_assign_elastic_ips     *bool
	Auto_assign_public_ips      *bool
	Custom_instance_profile_arn *string
}

func Aws_opsworks_custom_layerMapper(r *Aws_opsworks_custom_layer) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Auto_assign_elastic_ips != nil {
		config["auto_assign_elastic_ips"] = *r.Auto_assign_elastic_ips
	}
	if r.Auto_assign_public_ips != nil {
		config["auto_assign_public_ips"] = *r.Auto_assign_public_ips
	}
	if r.Custom_instance_profile_arn != nil {
		config["custom_instance_profile_arn"] = *r.Custom_instance_profile_arn
	}
	if r.Elastic_load_balancer != nil {
		config["elastic_load_balancer"] = *r.Elastic_load_balancer
	}
	if r.Install_updates_on_boot != nil {
		config["install_updates_on_boot"] = *r.Install_updates_on_boot
	}
	if r.Drain_elb_on_shutdown != nil {
		config["drain_elb_on_shutdown"] = *r.Drain_elb_on_shutdown
	}
	if r.Use_ebs_optimized_instances != nil {
		config["use_ebs_optimized_instances"] = *r.Use_ebs_optimized_instances
	}
	config["stack_id"] = r.Stack_id
	config["short_name"] = r.Short_name
	config["name"] = r.Name
	if r.Custom_json != nil {
		config["custom_json"] = *r.Custom_json
	}
	if r.Auto_healing != nil {
		config["auto_healing"] = *r.Auto_healing
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_redshift_security_group struct {
	Name        string
	Description *string
}

func Aws_redshift_security_groupMapper(r *Aws_redshift_security_group) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["name"] = r.Name
	if r.Description != nil {
		config["description"] = *r.Description
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_elasticache_security_group struct {
	Description *string
	Name        string
}

func Aws_elasticache_security_groupMapper(r *Aws_elasticache_security_group) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Description != nil {
		config["description"] = *r.Description
	}
	config["name"] = r.Name
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_s3_account_public_access_block struct {
	Restrict_public_buckets *bool
	Account_id              *string
	Block_public_acls       *bool
	Block_public_policy     *bool
	Ignore_public_acls      *bool
}

func Aws_s3_account_public_access_blockMapper(r *Aws_s3_account_public_access_block) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Account_id != nil {
		config["account_id"] = *r.Account_id
	}
	if r.Block_public_acls != nil {
		config["block_public_acls"] = *r.Block_public_acls
	}
	if r.Block_public_policy != nil {
		config["block_public_policy"] = *r.Block_public_policy
	}
	if r.Ignore_public_acls != nil {
		config["ignore_public_acls"] = *r.Ignore_public_acls
	}
	if r.Restrict_public_buckets != nil {
		config["restrict_public_buckets"] = *r.Restrict_public_buckets
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_elasticsearch_domain_policy struct {
	Domain_name     string
	Access_policies string
}

func Aws_elasticsearch_domain_policyMapper(r *Aws_elasticsearch_domain_policy) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["domain_name"] = r.Domain_name
	config["access_policies"] = r.Access_policies
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_iot_thing_principal_attachment struct {
	Principal string
	Thing     string
}

func Aws_iot_thing_principal_attachmentMapper(r *Aws_iot_thing_principal_attachment) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["principal"] = r.Principal
	config["thing"] = r.Thing
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_opsworks_ganglia_layer struct {
	Install_updates_on_boot     *bool
	Stack_id                    string
	Name                        *string
	Drain_elb_on_shutdown       *bool
	Use_ebs_optimized_instances *bool
	Elastic_load_balancer       *string
	Custom_json                 *string
	Auto_healing                *bool
	Auto_assign_elastic_ips     *bool
	Auto_assign_public_ips      *bool
	Custom_instance_profile_arn *string
	Password                    string
	Url                         *string
	Username                    *string
}

func Aws_opsworks_ganglia_layerMapper(r *Aws_opsworks_ganglia_layer) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Url != nil {
		config["url"] = *r.Url
	}
	if r.Username != nil {
		config["username"] = *r.Username
	}
	config["password"] = r.Password
	if r.Install_updates_on_boot != nil {
		config["install_updates_on_boot"] = *r.Install_updates_on_boot
	}
	config["stack_id"] = r.Stack_id
	if r.Name != nil {
		config["name"] = *r.Name
	}
	if r.Elastic_load_balancer != nil {
		config["elastic_load_balancer"] = *r.Elastic_load_balancer
	}
	if r.Custom_json != nil {
		config["custom_json"] = *r.Custom_json
	}
	if r.Auto_healing != nil {
		config["auto_healing"] = *r.Auto_healing
	}
	if r.Drain_elb_on_shutdown != nil {
		config["drain_elb_on_shutdown"] = *r.Drain_elb_on_shutdown
	}
	if r.Use_ebs_optimized_instances != nil {
		config["use_ebs_optimized_instances"] = *r.Use_ebs_optimized_instances
	}
	if r.Auto_assign_elastic_ips != nil {
		config["auto_assign_elastic_ips"] = *r.Auto_assign_elastic_ips
	}
	if r.Auto_assign_public_ips != nil {
		config["auto_assign_public_ips"] = *r.Auto_assign_public_ips
	}
	if r.Custom_instance_profile_arn != nil {
		config["custom_instance_profile_arn"] = *r.Custom_instance_profile_arn
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_api_gateway_documentation_part struct {
	Properties  string
	Rest_api_id string
}

func Aws_api_gateway_documentation_partMapper(r *Aws_api_gateway_documentation_part) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	config["rest_api_id"] = r.Rest_api_id
	config["properties"] = r.Properties
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_cloudfront_distribution struct {
	Http_version           *string
	Active_trusted_signers map[string]interface{}
	Domain_name            string
	Last_modified_time     string
	Comment                *string
	Etag                   string
	Hosted_zone_id         string
	Default_root_object    *string
	Enabled                bool
	Caller_reference       string
	Price_class            *string
	Status                 string
	Retain_on_delete       *bool
	Is_ipv6_enabled        *bool
	Web_acl_id             *string
	Arn                    string
	Tags                   *map[string]interface{}
}

func Aws_cloudfront_distributionMapper(r *Aws_cloudfront_distribution) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Price_class != nil {
		config["price_class"] = *r.Price_class
	}
	config["status"] = r.Status
	if r.Retain_on_delete != nil {
		config["retain_on_delete"] = *r.Retain_on_delete
	}
	if r.Is_ipv6_enabled != nil {
		config["is_ipv6_enabled"] = *r.Is_ipv6_enabled
	}
	if r.Web_acl_id != nil {
		config["web_acl_id"] = *r.Web_acl_id
	}
	config["arn"] = r.Arn
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	if r.Http_version != nil {
		config["http_version"] = *r.Http_version
	}
	config["active_trusted_signers"] = r.Active_trusted_signers
	config["domain_name"] = r.Domain_name
	config["last_modified_time"] = r.Last_modified_time
	if r.Comment != nil {
		config["comment"] = *r.Comment
	}
	config["etag"] = r.Etag
	config["hosted_zone_id"] = r.Hosted_zone_id
	if r.Default_root_object != nil {
		config["default_root_object"] = *r.Default_root_object
	}
	config["enabled"] = r.Enabled
	config["caller_reference"] = r.Caller_reference
	return &terraform.ResourceConfig{
		Config: config,
	}
}

type Aws_route53_health_check struct {
	Invert_healthcheck              *bool
	Measure_latency                 *bool
	Reference_name                  *string
	Enable_sni                      *bool
	Tags                            *map[string]interface{}
	Fqdn                            *string
	Ip_address                      *string
	Resource_type                   string
	Cloudwatch_alarm_name           *string
	Cloudwatch_alarm_region         *string
	Resource_path                   *string
	Search_string                   *string
	Insufficient_data_health_status *string
}

func Aws_route53_health_checkMapper(r *Aws_route53_health_check) *terraform.ResourceConfig {
	config := map[string]interface{}{}
	if r.Ip_address != nil {
		config["ip_address"] = *r.Ip_address
	}
	config["resource_type"] = r.Resource_type
	if r.Cloudwatch_alarm_name != nil {
		config["cloudwatch_alarm_name"] = *r.Cloudwatch_alarm_name
	}
	if r.Cloudwatch_alarm_region != nil {
		config["cloudwatch_alarm_region"] = *r.Cloudwatch_alarm_region
	}
	if r.Resource_path != nil {
		config["resource_path"] = *r.Resource_path
	}
	if r.Search_string != nil {
		config["search_string"] = *r.Search_string
	}
	if r.Insufficient_data_health_status != nil {
		config["insufficient_data_health_status"] = *r.Insufficient_data_health_status
	}
	if r.Invert_healthcheck != nil {
		config["invert_healthcheck"] = *r.Invert_healthcheck
	}
	if r.Measure_latency != nil {
		config["measure_latency"] = *r.Measure_latency
	}
	if r.Reference_name != nil {
		config["reference_name"] = *r.Reference_name
	}
	if r.Enable_sni != nil {
		config["enable_sni"] = *r.Enable_sni
	}
	if r.Tags != nil {
		config["tags"] = *r.Tags
	}
	if r.Fqdn != nil {
		config["fqdn"] = *r.Fqdn
	}
	return &terraform.ResourceConfig{
		Config: config,
	}
}
