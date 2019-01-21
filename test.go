package awstypes

// Goodnight sweet linter, I'm sorry

type aws_vpn_connection_route struct {
	destination_cidr_block string
	vpn_connection_id      string
}

type aws_iam_service_linked_role struct {
	arn              string
	create_date      string
	unique_id        string
	custom_suffix    *string
	description      *string
	aws_service_name string
	name             string
	path             string
}

type aws_network_acl struct {
	tags      *map[string]interface{}
	owner_id  string
	vpc_id    string
	subnet_id *string
}

type aws_s3_bucket_public_access_block struct {
	bucket                  string
	block_public_acls       *bool
	block_public_policy     *bool
	ignore_public_acls      *bool
	restrict_public_buckets *bool
}

type aws_network_interface_sg_attachment struct {
	security_group_id    string
	network_interface_id string
}

type aws_cloudfront_public_key struct {
	etag             string
	name             *string
	name_prefix      *string
	caller_reference string
	comment          *string
	encoded_key      string
}

type aws_config_delivery_channel struct {
	sns_topic_arn  *string
	name           *string
	s3_bucket_name string
	s3_key_prefix  *string
}

type aws_batch_job_queue struct {
	name  string
	state string
	arn   string
}

type aws_glue_catalog_database struct {
	name         string
	description  *string
	location_uri *string
	parameters   *map[string]interface{}
	catalog_id   *string
}

type aws_inspector_resource_group struct {
	tags map[string]interface{}
	arn  string
}

type aws_cloudhsm_v2_cluster struct {
	cluster_state            string
	hsm_type                 string
	cluster_id               string
	vpc_id                   string
	source_backup_identifier *string
	security_group_id        string
	tags                     *map[string]interface{}
}

type aws_ssm_maintenance_window_target struct {
	window_id         string
	resource_type     string
	owner_information *string
}

type aws_lb struct {
	arn_suffix                       string
	name_prefix                      *string
	enable_deletion_protection       *bool
	name                             *string
	load_balancer_type               *string
	enable_http2                     *bool
	zone_id                          string
	dns_name                         string
	internal                         *bool
	enable_cross_zone_load_balancing *bool
	ip_address_type                  *string
	vpc_id                           string
	tags                             *map[string]interface{}
	arn                              string
}

type aws_dynamodb_global_table struct {
	name string
	arn  string
}

type aws_opsworks_instance struct {
	infrastructure_class         *string
	reported_os_version          *string
	public_dns                   *string
	reported_os_family           *string
	ssh_host_dsa_key_fingerprint *string
	ssh_key_name                 *string
	ami_id                       *string
	architecture                 *string
	private_dns                  *string
	ecs_cluster_arn              *string
	elastic_ip                   *string
	os                           *string
	private_ip                   *string
	root_device_volume_id        *string
	auto_scaling_type            *string
	delete_ebs                   *bool
	delete_eip                   *bool
	tenancy                      *string
	virtualization_type          *string
	ebs_optimized                *bool
	last_service_error_id        *string
	subnet_id                    *string
	agent_version                *string
	hostname                     *string
	instance_type                *string
	state                        *string
	ec2_instance_id              string
	instance_profile_arn         *string
	public_ip                    *string
	reported_agent_version       *string
	reported_os_name             *string
	ssh_host_rsa_key_fingerprint *string
	status                       *string
	availability_zone            *string
	created_at                   *string
	platform                     *string
	stack_id                     string
	install_updates_on_boot      *bool
	registered_by                *string
	root_device_type             *string
}

type aws_cognito_resource_server struct {
	name         string
	user_pool_id string
	identifier   string
}

type aws_ec2_transit_gateway_route_table_association struct {
	resource_id                    string
	resource_type                  string
	transit_gateway_attachment_id  string
	transit_gateway_route_table_id string
}

type aws_ses_domain_mail_from struct {
	domain                 string
	mail_from_domain       string
	behavior_on_mx_failure *string
}

type aws_sns_sms_preferences struct {
	default_sender_id                     *string
	default_sms_type                      *string
	usage_report_s3_bucket                *string
	monthly_spend_limit                   *string
	delivery_status_iam_role_arn          *string
	delivery_status_success_sampling_rate *string
}

type aws_alb_target_group_attachment struct {
	target_group_arn  string
	target_id         string
	availability_zone *string
}

type aws_ami_copy struct {
	ena_support          bool
	encrypted            *bool
	image_location       string
	root_device_name     string
	architecture         string
	description          *string
	kernel_id            string
	kms_key_id           *string
	manage_ebs_snapshots bool
	name                 string
	ramdisk_id           string
	source_ami_id        string
	virtualization_type  string
	source_ami_region    string
	sriov_net_support    string
	tags                 *map[string]interface{}
	root_snapshot_id     string
}

type aws_cloudwatch_event_target struct {
	target_id  *string
	input      *string
	role_arn   *string
	rule       string
	arn        string
	input_path *string
}

type aws_elasticache_subnet_group struct {
	description *string
	name        string
}

type aws_ses_active_receipt_rule_set struct {
	rule_set_name string
}

type aws_codedeploy_app struct {
	name             string
	compute_platform *string
	unique_id        *string
}

type aws_dax_parameter_group struct {
	name        string
	description *string
}

type aws_lightsail_domain struct {
	domain_name string
	arn         string
}

type aws_redshift_cluster struct {
	final_snapshot_identifier    *string
	cluster_public_key           *string
	cluster_type                 *string
	node_type                    string
	allow_version_upgrade        *bool
	enhanced_vpc_routing         *bool
	elastic_ip                   *string
	dns_name                     string
	tags                         *map[string]interface{}
	cluster_identifier           string
	encrypted                    *bool
	database_name                *string
	skip_final_snapshot          *bool
	enable_logging               *bool
	s3_key_prefix                *string
	owner_account                *string
	publicly_accessible          *bool
	kms_key_id                   *string
	bucket_name                  *string
	snapshot_identifier          *string
	cluster_subnet_group_name    *string
	endpoint                     *string
	cluster_revision_number      *string
	availability_zone            *string
	cluster_parameter_group_name *string
	preferred_maintenance_window *string
	cluster_version              *string
	snapshot_cluster_identifier  *string
	master_username              *string
	master_password              *string
}

type aws_load_balancer_policy struct {
	load_balancer_name string
	policy_name        string
	policy_type_name   string
}

type aws_autoscaling_attachment struct {
	autoscaling_group_name string
	elb                    *string
	alb_target_group_arn   *string
}

type aws_cloudfront_origin_access_identity struct {
	etag                            string
	iam_arn                         string
	s3_canonical_user_id            string
	comment                         *string
	caller_reference                string
	cloudfront_access_identity_path string
}

type aws_cloudwatch_log_metric_filter struct {
	log_group_name string
	name           string
	pattern        string
}

type aws_eks_cluster struct {
	endpoint         string
	name             string
	platform_version string
	role_arn         string
	version          *string
	arn              string
	created_at       string
}

type aws_default_subnet struct {
	vpc_id                          string
	availability_zone               string
	availability_zone_id            string
	arn                             string
	owner_id                        string
	tags                            *map[string]interface{}
	cidr_block                      string
	ipv6_cidr_block                 string
	map_public_ip_on_launch         *bool
	assign_ipv6_address_on_creation bool
	ipv6_cidr_block_association_id  string
}

type aws_wafregional_web_acl struct {
	name        string
	metric_name string
}

type aws_ec2_transit_gateway_route_table struct {
	tags                            *map[string]interface{}
	transit_gateway_id              string
	default_association_route_table bool
	default_propagation_route_table bool
}

type aws_iam_user struct {
	unique_id            string
	name                 string
	path                 *string
	permissions_boundary *string
	force_destroy        *bool
	tags                 *map[string]interface{}
	arn                  string
}

type aws_ses_receipt_rule struct {
	scan_enabled  *bool
	enabled       *bool
	tls_policy    *string
	name          string
	after         *string
	rule_set_name string
}

type aws_simpledb_domain struct {
	name string
}

type aws_cognito_user_pool struct {
	arn                        string
	name                       string
	tags                       *map[string]interface{}
	creation_date              string
	email_verification_message *string
	last_modified_date         string
	sms_authentication_message *string
	email_verification_subject *string
	endpoint                   string
	mfa_configuration          *string
	sms_verification_message   *string
}

type aws_neptune_cluster struct {
	snapshot_identifier                  *string
	cluster_identifier                   *string
	cluster_resource_id                  string
	preferred_maintenance_window         *string
	tags                                 *map[string]interface{}
	endpoint                             string
	iam_database_authentication_enabled  *bool
	replication_source_identifier        *string
	skip_final_snapshot                  *bool
	apply_immediately                    *bool
	engine_version                       *string
	preferred_backup_window              *string
	storage_encrypted                    *bool
	neptune_cluster_parameter_group_name *string
	final_snapshot_identifier            *string
	kms_key_arn                          *string
	hosted_zone_id                       string
	reader_endpoint                      string
	engine                               *string
	neptune_subnet_group_name            *string
	arn                                  string
	cluster_identifier_prefix            *string
}

type aws_placement_group struct {
	name     string
	strategy string
}

type aws_wafregional_geo_match_set struct {
	name string
}

type aws_wafregional_size_constraint_set struct {
	name string
}

type aws_iam_group_membership struct {
	group string
	name  string
}

type aws_opsworks_application struct {
	name                      string
	stack_id                  string
	document_root             *string
	short_name                *string
	data_source_database_name *string
	description               *string
	rails_env                 *string
	aws_flow_ruby_settings    *string
	data_source_type          *string
	resource_type             string
	auto_bundle_on_deploy     *string
	data_source_arn           *string
	enable_ssl                *bool
}

type aws_organizations_policy struct {
	name          string
	resource_type *string
	arn           string
	content       string
	description   *string
}

type aws_snapshot_create_volume_permission struct {
	account_id  string
	snapshot_id string
}

type aws_api_gateway_client_certificate struct {
	pem_encoded_certificate string
	description             *string
	created_date            string
	expiration_date         string
}

type aws_dx_connection struct {
	bandwidth           string
	location            string
	jumbo_frame_capable bool
	tags                *map[string]interface{}
	arn                 string
	name                string
}

type aws_dynamodb_table_item struct {
	table_name string
	hash_key   string
	range_key  *string
	item       string
}

type aws_globalaccelerator_accelerator struct {
	name            string
	ip_address_type *string
	enabled         *bool
}

type aws_waf_web_acl struct {
	name        string
	metric_name string
}

type aws_api_gateway_usage_plan struct {
	product_code *string
	name         string
	description  *string
}

type aws_iam_account_alias struct {
	account_alias string
}

type aws_batch_compute_environment struct {
	service_role             string
	arn                      string
	ecs_cluster_arn          string
	status                   string
	compute_environment_name string
	state                    *string
	resource_type            string
	ecc_cluster_arn          string
	status_reason            string
}

type aws_cloudfront_distribution struct {
	web_acl_id             *string
	active_trusted_signers map[string]interface{}
	last_modified_time     string
	arn                    string
	hosted_zone_id         string
	domain_name            string
	price_class            *string
	is_ipv6_enabled        *bool
	comment                *string
	status                 string
	etag                   string
	enabled                bool
	caller_reference       string
	retain_on_delete       *bool
	default_root_object    *string
	http_version           *string
	tags                   *map[string]interface{}
}

type aws_dynamodb_table struct {
	name             string
	arn              string
	billing_mode     *string
	stream_enabled   *bool
	stream_view_type *string
	hash_key         string
	stream_label     string
	tags             *map[string]interface{}
	range_key        *string
	stream_arn       string
}

type aws_rds_cluster_instance struct {
	db_subnet_group_name            *string
	publicly_accessible             *bool
	apply_immediately               *bool
	performance_insights_kms_key_id *string
	identifier                      *string
	cluster_identifier              string
	engine                          *string
	engine_version                  *string
	availability_zone               *string
	writer                          bool
	storage_encrypted               bool
	instance_class                  string
	preferred_maintenance_window    *string
	performance_insights_enabled    *bool
	dbi_resource_id                 string
	auto_minor_version_upgrade      *bool
	preferred_backup_window         *string
	tags                            *map[string]interface{}
	endpoint                        string
	db_parameter_group_name         *string
	kms_key_id                      string
	monitoring_role_arn             *string
	arn                             string
	identifier_prefix               *string
	copy_tags_to_snapshot           *bool
}

type aws_redshift_security_group struct {
	name        string
	description *string
}

type aws_api_gateway_deployment struct {
	created_date      string
	invoke_url        string
	execution_arn     string
	rest_api_id       string
	stage_name        string
	description       *string
	stage_description *string
	variables         *map[string]interface{}
}

type aws_app_cookie_stickiness_policy struct {
	name          string
	load_balancer string
	cookie_name   string
}

type aws_resourcegroups_group struct {
	description *string
	arn         string
	name        string
}

type aws_redshift_subnet_group struct {
	name        string
	description *string
	tags        *map[string]interface{}
}

type aws_securityhub_standards_subscription struct {
	standards_arn string
}

type aws_waf_byte_match_set struct {
	name string
}

type aws_api_gateway_documentation_part struct {
	properties  string
	rest_api_id string
}

type aws_config_configuration_recorder struct {
	name     *string
	role_arn string
}

type aws_dx_private_virtual_interface struct {
	vpn_gateway_id      *string
	amazon_address      *string
	jumbo_frame_capable bool
	connection_id       string
	bgp_auth_key        *string
	address_family      string
	customer_address    *string
	arn                 string
	name                string
	dx_gateway_id       *string
	tags                *map[string]interface{}
}

type aws_ecr_lifecycle_policy struct {
	repository  string
	policy      string
	registry_id string
}

type aws_db_cluster_snapshot struct {
	engine                         string
	db_cluster_snapshot_identifier string
	source_db_cluster_snapshot_arn string
	db_cluster_identifier          string
	db_cluster_snapshot_arn        string
	license_model                  string
	snapshot_type                  string
	vpc_id                         string
	storage_encrypted              bool
	engine_version                 string
	kms_key_id                     string
	status                         string
}

type aws_iot_thing struct {
	thing_type_name   *string
	default_client_id string
	arn               string
	name              string
	attributes        *map[string]interface{}
}

type aws_ses_receipt_rule_set struct {
	rule_set_name string
}

type aws_vpc_endpoint_route_table_association struct {
	route_table_id  string
	vpc_endpoint_id string
}

type aws_appmesh_virtual_node struct {
	arn               string
	created_date      string
	last_updated_date string
	name              string
	mesh_name         string
}

type aws_efs_file_system struct {
	creation_token   *string
	encrypted        *bool
	kms_key_id       *string
	dns_name         string
	arn              string
	reference_name   *string
	performance_mode *string
	tags             *map[string]interface{}
	throughput_mode  *string
}

type aws_kinesis_stream struct {
	arn             *string
	tags            *map[string]interface{}
	name            string
	encryption_type *string
	kms_key_id      *string
}

type aws_vpc_dhcp_options struct {
	netbios_node_type *string
	tags              *map[string]interface{}
	owner_id          string
	domain_name       *string
}

type aws_transfer_server struct {
	arn                    string
	endpoint               string
	invocation_role        *string
	url                    *string
	identity_provider_type *string
	logging_role           *string
	force_destroy          *bool
	tags                   *map[string]interface{}
}

type aws_batch_job_definition struct {
	arn                  string
	name                 string
	container_properties *string
	parameters           *map[string]interface{}
	resource_type        string
}

type aws_cloud9_environment_ec2 struct {
	subnet_id     *string
	arn           string
	resource_type string
	name          string
	instance_type string
	description   *string
	owner_arn     *string
}

type aws_ebs_volume struct {
	resource_type     *string
	arn               string
	availability_zone string
	encrypted         *bool
	kms_key_id        *string
	snapshot_id       *string
	tags              *map[string]interface{}
}

type aws_kinesis_analytics_application struct {
	create_timestamp      string
	status                string
	arn                   string
	code                  *string
	description           *string
	last_update_timestamp string
	name                  string
}

type aws_spot_fleet_request struct {
	excess_capacity_termination_policy  *string
	spot_price                          *string
	wait_for_fulfillment                *bool
	replace_unhealthy_instances         *bool
	instance_interruption_behaviour     *string
	terminate_instances_with_expiration *bool
	valid_from                          *string
	valid_until                         *string
	iam_fleet_role                      string
	allocation_strategy                 *string
	fleet_type                          *string
	spot_request_state                  string
	client_token                        string
}

type aws_cognito_identity_provider struct {
	user_pool_id      string
	attribute_mapping *map[string]interface{}
	provider_details  map[string]interface{}
	provider_name     string
	provider_type     string
}

type aws_ses_template struct {
	name    string
	html    *string
	subject *string
	text    *string
}

type aws_api_gateway_account struct {
	cloudwatch_role_arn *string
}

type aws_ec2_transit_gateway_route struct {
	destination_cidr_block         string
	transit_gateway_attachment_id  string
	transit_gateway_route_table_id string
}

type aws_cloudformation_stack struct {
	on_failure       *string
	parameters       *map[string]interface{}
	policy_url       *string
	tags             *map[string]interface{}
	iam_role_arn     *string
	name             string
	template_body    *string
	disable_rollback *bool
	outputs          map[string]interface{}
	policy_body      *string
	template_url     *string
}

type aws_iam_user_login_profile struct {
	password_reset_required *bool
	key_fingerprint         string
	encrypted_password      string
	user                    string
	pgp_key                 string
}

type aws_opsworks_ganglia_layer struct {
	password                    string
	drain_elb_on_shutdown       *bool
	custom_json                 *string
	auto_healing                *bool
	install_updates_on_boot     *bool
	use_ebs_optimized_instances *bool
	auto_assign_public_ips      *bool
	username                    *string
	stack_id                    string
	name                        *string
	auto_assign_elastic_ips     *bool
	url                         *string
	custom_instance_profile_arn *string
	elastic_load_balancer       *string
}

type aws_route53_zone struct {
	name              string
	zone_id           string
	delegation_set_id *string
	tags              *map[string]interface{}
	comment           *string
	vpc_id            *string
	vpc_region        *string
	force_destroy     *bool
}

type aws_api_gateway_rest_api struct {
	root_resource_id string
	created_date     string
	description      *string
	api_key_source   *string
	policy           *string
	body             *string
	name             string
	execution_arn    string
}

type aws_appautoscaling_scheduled_action struct {
	service_namespace  string
	resource_id        string
	scalable_dimension *string
	schedule           *string
	start_time         *string
	end_time           *string
	arn                string
	name               string
}

type aws_route_table struct {
	vpc_id   string
	tags     *map[string]interface{}
	owner_id string
}

type aws_appmesh_virtual_router struct {
	created_date      string
	last_updated_date string
	name              string
	mesh_name         string
	arn               string
}

type aws_key_pair struct {
	key_name        *string
	key_name_prefix *string
	public_key      string
	fingerprint     string
}

type aws_lb_listener_certificate struct {
	listener_arn    string
	certificate_arn string
}

type aws_api_gateway_integration struct {
	http_method                string
	resource_type              string
	connection_id              *string
	request_templates          *map[string]interface{}
	content_handling           *string
	resource_id                string
	credentials                *string
	request_parameters_in_json *string
	rest_api_id                string
	integration_http_method    *string
	connection_type            *string
	uri                        *string
	request_parameters         *map[string]interface{}
	passthrough_behavior       *string
	cache_namespace            *string
}

type aws_vpc_endpoint_connection_notification struct {
	vpc_endpoint_service_id     *string
	vpc_endpoint_id             *string
	connection_notification_arn string
	state                       string
	notification_type           string
}

type aws_emr_cluster struct {
	release_label                     string
	core_instance_type                *string
	configurations_json               *string
	service_role                      string
	configurations                    *string
	security_configuration            *string
	additional_info                   *string
	cluster_state                     string
	visible_to_all_users              *bool
	termination_protection            *bool
	autoscaling_role                  *string
	master_public_dns                 string
	keep_job_flow_alive_when_no_steps *bool
	tags                              *map[string]interface{}
	scale_down_behavior               *string
	name                              string
	master_instance_type              *string
	log_uri                           *string
	custom_ami_id                     *string
}

type aws_organizations_account struct {
	joined_timestamp           string
	status                     string
	name                       string
	email                      string
	iam_user_access_to_billing *string
	role_name                  *string
	arn                        string
	joined_method              string
}

type aws_waf_ipset struct {
	name string
	arn  string
}

type aws_waf_regex_pattern_set struct {
	name string
}

type aws_appautoscaling_policy struct {
	scalable_dimension      string
	adjustment_type         *string
	resource_id             string
	service_namespace       string
	arn                     string
	metric_aggregation_type *string
	name                    string
	policy_type             *string
}

type aws_guardduty_threatintelset struct {
	detector_id string
	name        string
	format      string
	location    string
	activate    bool
}

type aws_redshift_parameter_group struct {
	name        string
	family      string
	description *string
}

type aws_vpc_endpoint_subnet_association struct {
	vpc_endpoint_id string
	subnet_id       string
}

type aws_appsync_datasource struct {
	api_id           string
	description      *string
	service_role_arn *string
	name             string
	resource_type    string
	arn              string
}

type aws_iot_policy_attachment struct {
	policy string
	target string
}

type aws_lb_listener_rule struct {
	arn          string
	listener_arn string
}

type aws_ses_domain_identity_verification struct {
	arn    string
	domain string
}

type aws_ssm_association struct {
	name                string
	schedule_expression *string
	association_name    *string
	association_id      string
	instance_id         *string
	document_version    *string
	parameters          *map[string]interface{}
}

type aws_sqs_queue struct {
	name                        *string
	content_based_deduplication *bool
	redrive_policy              *string
	name_prefix                 *string
	fifo_queue                  *bool
	kms_master_key_id           *string
	policy                      *string
	arn                         string
	tags                        *map[string]interface{}
}

type aws_athena_named_query struct {
	name        string
	query       string
	database    string
	description *string
}

type aws_iam_saml_provider struct {
	arn                    string
	valid_until            string
	name                   string
	saml_metadata_document string
}

type aws_mq_broker struct {
	engine_type                string
	engine_version             string
	arn                        string
	auto_minor_version_upgrade *bool
	deployment_mode            *string
	host_instance_type         string
	broker_name                string
	publicly_accessible        *bool
	apply_immediately          *bool
}

type aws_neptune_cluster_snapshot struct {
	kms_key_id                     string
	license_model                  string
	db_cluster_identifier          string
	db_cluster_snapshot_arn        string
	engine_version                 string
	storage_encrypted              bool
	vpc_id                         string
	source_db_cluster_snapshot_arn string
	db_cluster_snapshot_identifier string
	engine                         string
	snapshot_type                  string
	status                         string
}

type aws_pinpoint_apns_voip_channel struct {
	default_authentication_method *string
	private_key                   *string
	token_key                     *string
	token_key_id                  *string
	application_id                string
	bundle_id                     *string
	certificate                   *string
	enabled                       *bool
	team_id                       *string
}

type aws_ebs_snapshot_copy struct {
	data_encryption_key_id string
	source_region          string
	source_snapshot_id     string
	description            *string
	owner_id               string
	owner_alias            string
	encrypted              *bool
	kms_key_id             *string
	volume_id              string
	tags                   *map[string]interface{}
}

type aws_ec2_transit_gateway struct {
	association_default_route_table_id string
	auto_accept_shared_attachments     *string
	default_route_table_propagation    *string
	description                        *string
	owner_id                           string
	propagation_default_route_table_id string
	vpn_ecmp_support                   *string
	arn                                string
	default_route_table_association    *string
	dns_support                        *string
	tags                               *map[string]interface{}
}

type aws_storagegateway_working_storage struct {
	disk_id     string
	gateway_arn string
}

type aws_sqs_queue_policy struct {
	queue_url string
	policy    string
}

type aws_docdb_subnet_group struct {
	name        *string
	name_prefix *string
	description *string
	tags        *map[string]interface{}
	arn         string
}

type aws_dx_connection_association struct {
	connection_id string
	lag_id        string
}

type aws_sns_topic_policy struct {
	arn    string
	policy string
}

type aws_cloudhsm_v2_hsm struct {
	hsm_state         string
	hsm_eni_id        string
	cluster_id        string
	subnet_id         *string
	availability_zone *string
	ip_address        *string
	hsm_id            string
}

type aws_gamelift_alias struct {
	name        string
	description *string
	arn         string
}

type aws_iot_policy struct {
	name               string
	policy             string
	arn                string
	default_version_id string
}

type aws_opsworks_stack struct {
	default_instance_profile_arn  string
	default_os                    *string
	hostname_theme                *string
	color                         *string
	configuration_manager_name    *string
	default_availability_zone     *string
	default_ssh_key_name          *string
	agent_version                 *string
	arn                           string
	service_role_arn              string
	tags                          *map[string]interface{}
	use_opsworks_security_groups  *bool
	custom_json                   *string
	default_root_device_type      *string
	use_custom_cookbooks          *bool
	name                          string
	configuration_manager_version *string
	berkshelf_version             *string
	vpc_id                        *string
	stack_endpoint                string
	region                        string
	manage_berkshelf              *bool
	default_subnet_id             *string
}

type aws_config_configuration_aggregator struct {
	arn  string
	name string
}

type aws_lightsail_key_pair struct {
	name_prefix           *string
	arn                   string
	name                  *string
	pgp_key               *string
	fingerprint           string
	public_key            *string
	private_key           string
	encrypted_fingerprint string
	encrypted_private_key string
}

type aws_s3_bucket_inventory struct {
	included_object_versions string
	bucket                   string
	name                     string
	enabled                  *bool
}

type aws_storagegateway_nfs_file_share struct {
	fileshare_id            string
	kms_encrypted           *bool
	read_only               *bool
	role_arn                string
	default_storage_class   *string
	gateway_arn             string
	guess_mime_type_enabled *bool
	object_acl              *string
	requester_pays          *bool
	arn                     string
	kms_key_arn             *string
	location_arn            string
	squash                  *string
}

type aws_config_config_rule struct {
	name                        string
	rule_id                     string
	arn                         string
	description                 *string
	input_parameters            *string
	maximum_execution_frequency *string
}

type aws_media_package_channel struct {
	arn         string
	channel_id  string
	description *string
}

type aws_redshift_event_subscription struct {
	status          string
	source_type     *string
	severity        *string
	tags            *map[string]interface{}
	name            string
	sns_topic_arn   string
	enabled         *bool
	customer_aws_id string
}

type aws_spot_datafeed_subscription struct {
	bucket string
	prefix *string
}

type aws_waf_geo_match_set struct {
	name string
}

type aws_wafregional_sql_injection_match_set struct {
	name string
}

type aws_pinpoint_gcm_channel struct {
	api_key        string
	enabled        *bool
	application_id string
}

type aws_api_gateway_base_path_mapping struct {
	base_path   *string
	stage_name  *string
	domain_name string
	api_id      string
}

type aws_cognito_user_group struct {
	description  *string
	name         string
	role_arn     *string
	user_pool_id string
}

type aws_route struct {
	gateway_id                  *string
	instance_id                 *string
	origin                      string
	state                       string
	destination_ipv6_cidr_block *string
	egress_only_gateway_id      *string
	nat_gateway_id              *string
	network_interface_id        *string
	route_table_id              string
	transit_gateway_id          *string
	vpc_peering_connection_id   *string
	destination_cidr_block      *string
	destination_prefix_list_id  string
	instance_owner_id           string
}

type aws_service_discovery_http_namespace struct {
	name        string
	description *string
	arn         string
}

type aws_iam_account_password_policy struct {
	expire_passwords               bool
	hard_expiry                    *bool
	require_numbers                *bool
	require_symbols                *bool
	require_uppercase_characters   *bool
	allow_users_to_change_password *bool
	require_lowercase_characters   *bool
}

type aws_rds_cluster_parameter_group struct {
	description *string
	tags        *map[string]interface{}
	arn         string
	name        *string
	name_prefix *string
	family      string
}

type aws_acm_certificate_validation struct {
	certificate_arn string
}

type aws_config_aggregate_authorization struct {
	arn        string
	account_id string
	region     string
}

type aws_codedeploy_deployment_group struct {
	app_name               string
	deployment_group_name  string
	deployment_config_name *string
	service_role_arn       string
}

type aws_glue_crawler struct {
	security_configuration *string
	role                   string
	configuration          *string
	name                   string
	database_name          string
	description            *string
	schedule               *string
	table_prefix           *string
}

type aws_elb_attachment struct {
	elb      string
	instance string
}

type aws_licensemanager_association struct {
	resource_arn              string
	license_configuration_arn string
}

type aws_acmpca_certificate_authority struct {
	certificate_chain           string
	enabled                     *bool
	tags                        *map[string]interface{}
	resource_type               *string
	arn                         string
	certificate_signing_request string
	not_after                   string
	not_before                  string
	serial                      string
	status                      string
	certificate                 string
}

type aws_iam_policy struct {
	description *string
	path        *string
	policy      string
	name        *string
	name_prefix *string
	arn         string
}

type aws_ram_resource_share struct {
	tags                      *map[string]interface{}
	name                      string
	allow_external_principals *bool
}

type aws_db_parameter_group struct {
	family      string
	description *string
	tags        *map[string]interface{}
	arn         string
	name        *string
	name_prefix *string
}

type aws_lambda_permission struct {
	qualifier           *string
	source_account      *string
	source_arn          *string
	statement_id        *string
	statement_id_prefix *string
	action              string
	event_source_token  *string
	function_name       string
	principal           string
}

type aws_ses_configuration_set struct {
	name string
}

type aws_dx_gateway struct {
	name            string
	amazon_side_asn string
}

type aws_elastictranscoder_preset struct {
	container           string
	description         *string
	arn                 string
	name                *string
	resource_type       *string
	video_codec_options *map[string]interface{}
}

type aws_neptune_event_subscription struct {
	name_prefix     *string
	sns_topic_arn   string
	source_type     *string
	enabled         *bool
	tags            *map[string]interface{}
	arn             string
	name            *string
	customer_aws_id string
}

type aws_opsworks_rds_db_instance struct {
	stack_id            string
	rds_db_instance_arn string
	db_password         string
	db_user             string
}

type aws_default_route_table struct {
	tags                   *map[string]interface{}
	owner_id               string
	default_route_table_id string
	vpc_id                 string
}

type aws_ssm_patch_baseline struct {
	description                       *string
	operating_system                  *string
	approved_patches_compliance_level *string
	name                              string
}

type aws_api_gateway_vpc_link struct {
	name        string
	description *string
}

type aws_cloudwatch_event_permission struct {
	action       *string
	principal    string
	statement_id string
}

type aws_codecommit_repository struct {
	arn             string
	repository_id   string
	clone_url_http  string
	clone_url_ssh   string
	default_branch  *string
	repository_name string
	description     *string
}

type aws_elasticache_replication_group struct {
	configuration_endpoint_address string
	auth_token                     *string
	engine_version                 *string
	replication_group_id           string
	subnet_group_name              *string
	maintenance_window             *string
	parameter_group_name           *string
	replication_group_description  string
	tags                           *map[string]interface{}
	auto_minor_version_upgrade     *bool
	primary_endpoint_address       string
	snapshot_window                *string
	apply_immediately              *bool
	automatic_failover_enabled     *bool
	notification_topic_arn         *string
	engine                         *string
	snapshot_name                  *string
	transit_encryption_enabled     *bool
	at_rest_encryption_enabled     *bool
	node_type                      *string
}

type aws_api_gateway_api_key struct {
	created_date      string
	last_updated_date string
	value             *string
	name              string
	description       *string
	enabled           *bool
}

type aws_kms_grant struct {
	name               *string
	retiring_principal *string
	retire_on_delete   *bool
	key_id             string
	grantee_principal  string
	grant_id           string
	grant_token        string
}

type aws_docdb_cluster_parameter_group struct {
	tags        *map[string]interface{}
	arn         string
	name        *string
	name_prefix *string
	family      string
	description *string
}

type aws_network_interface struct {
	subnet_id         string
	private_ip        *string
	source_dest_check *bool
	private_dns_name  string
	description       *string
	tags              *map[string]interface{}
}

type aws_cloudwatch_metric_alarm struct {
	alarm_description                     *string
	unit                                  *string
	metric_name                           string
	namespace                             string
	statistic                             *string
	extended_statistic                    *string
	comparison_operator                   string
	treat_missing_data                    *string
	alarm_name                            string
	dimensions                            *map[string]interface{}
	arn                                   string
	actions_enabled                       *bool
	evaluate_low_sample_count_percentiles *string
}

type aws_codepipeline struct {
	arn      string
	name     string
	role_arn string
}

type aws_db_security_group struct {
	tags        *map[string]interface{}
	arn         string
	name        string
	description *string
}

type aws_directory_service_conditional_forwarder struct {
	directory_id       string
	remote_domain_name string
}

type aws_dax_cluster struct {
	node_type              string
	tags                   *map[string]interface{}
	configuration_endpoint string
	description            *string
	parameter_group_name   *string
	cluster_address        string
	cluster_name           string
	notification_topic_arn *string
	maintenance_window     *string
	subnet_group_name      *string
	arn                    string
	iam_role_arn           string
}

type aws_opsworks_rails_app_layer struct {
	drain_elb_on_shutdown       *bool
	name                        *string
	app_server                  *string
	rubygems_version            *string
	elastic_load_balancer       *string
	auto_healing                *bool
	use_ebs_optimized_instances *bool
	ruby_version                *string
	auto_assign_public_ips      *bool
	stack_id                    string
	manage_bundler              *bool
	bundler_version             *string
	passenger_version           *string
	auto_assign_elastic_ips     *bool
	custom_instance_profile_arn *string
	custom_json                 *string
	install_updates_on_boot     *bool
}

type aws_inspector_assessment_target struct {
	name               string
	arn                string
	resource_group_arn *string
}

type aws_ssm_maintenance_window_task struct {
	description      *string
	max_concurrency  string
	task_type        string
	task_arn         string
	service_role_arn string
	window_id        string
	max_errors       string
	name             *string
}

type aws_wafregional_rule_group struct {
	name        string
	metric_name string
}

type aws_appmesh_mesh struct {
	name              string
	arn               string
	created_date      string
	last_updated_date string
}

type aws_cognito_user_pool_client struct {
	client_secret                        string
	generate_secret                      *bool
	name                                 string
	default_redirect_uri                 *string
	allowed_oauth_flows_user_pool_client *bool
	user_pool_id                         string
}

type aws_iam_policy_attachment struct {
	name       string
	policy_arn string
}

type aws_iam_role_policy_attachment struct {
	role       string
	policy_arn string
}

type aws_glacier_vault struct {
	tags          *map[string]interface{}
	name          string
	location      string
	arn           string
	access_policy *string
}

type aws_iot_thing_principal_attachment struct {
	principal string
	thing     string
}

type aws_secretsmanager_secret struct {
	kms_key_id          *string
	name                *string
	policy              *string
	rotation_enabled    bool
	tags                *map[string]interface{}
	arn                 string
	description         *string
	rotation_lambda_arn *string
	name_prefix         *string
}

type aws_alb struct {
	name                             *string
	enable_deletion_protection       *bool
	enable_cross_zone_load_balancing *bool
	enable_http2                     *bool
	ip_address_type                  *string
	arn                              string
	arn_suffix                       string
	zone_id                          string
	dns_name                         string
	internal                         *bool
	load_balancer_type               *string
	vpc_id                           string
	tags                             *map[string]interface{}
	name_prefix                      *string
}

type aws_cloudwatch_log_group struct {
	name_prefix *string
	kms_key_id  *string
	arn         string
	tags        *map[string]interface{}
	name        *string
}

type aws_ses_domain_identity struct {
	arn                string
	domain             string
	verification_token string
}

type aws_pinpoint_apns_channel struct {
	bundle_id                     *string
	default_authentication_method *string
	enabled                       *bool
	team_id                       *string
	token_key                     *string
	application_id                string
	certificate                   *string
	private_key                   *string
	token_key_id                  *string
}

type aws_api_gateway_authorizer struct {
	identity_validation_expression *string
	authorizer_uri                 *string
	identity_source                *string
	name                           string
	resource_type                  *string
	authorizer_credentials         *string
	rest_api_id                    string
}

type aws_ecr_repository struct {
	registry_id    string
	repository_url string
	name           string
	tags           *map[string]interface{}
	arn            string
}

type aws_neptune_parameter_group struct {
	name        string
	family      string
	description *string
	tags        *map[string]interface{}
	arn         string
}

type aws_redshift_snapshot_copy_grant struct {
	snapshot_copy_grant_name string
	kms_key_id               *string
	tags                     *map[string]interface{}
}

type aws_dms_replication_instance struct {
	publicly_accessible          *bool
	tags                         *map[string]interface{}
	replication_subnet_group_id  *string
	preferred_maintenance_window *string
	replication_instance_arn     string
	replication_instance_class   string
	replication_instance_id      string
	apply_immediately            *bool
	auto_minor_version_upgrade   *bool
	multi_az                     *bool
	availability_zone            *string
	engine_version               *string
	kms_key_arn                  *string
}

type aws_iam_group_policy struct {
	policy      string
	name        *string
	name_prefix *string
	group       string
}

type aws_pinpoint_event_stream struct {
	role_arn               string
	application_id         string
	destination_stream_arn string
}

type aws_s3_bucket_policy struct {
	bucket string
	policy string
}

type aws_servicecatalog_portfolio struct {
	tags          *map[string]interface{}
	arn           string
	created_time  string
	name          string
	description   *string
	provider_name *string
}

type aws_vpc_endpoint_service_allowed_principal struct {
	vpc_endpoint_service_id string
	principal_arn           string
}

type aws_cloudwatch_log_subscription_filter struct {
	distribution    *string
	name            string
	destination_arn string
	filter_pattern  string
	log_group_name  string
	role_arn        *string
}

type aws_datasync_location_efs struct {
	tags                *map[string]interface{}
	uri                 string
	arn                 string
	efs_file_system_arn string
	subdirectory        *string
}

type aws_ec2_transit_gateway_vpc_attachment struct {
	dns_support                                     *string
	tags                                            *map[string]interface{}
	transit_gateway_default_route_table_association *bool
	transit_gateway_default_route_table_propagation *bool
	ipv6_support                                    *string
	transit_gateway_id                              string
	vpc_id                                          string
	vpc_owner_id                                    string
}

type aws_glue_job struct {
	default_arguments      *map[string]interface{}
	description            *string
	name                   string
	role_arn               string
	security_configuration *string
}

type aws_wafregional_xss_match_set struct {
	name string
}

type aws_pinpoint_adm_channel struct {
	enabled        *bool
	application_id string
	client_id      string
	client_secret  string
}

type aws_autoscaling_notification struct {
	topic_arn string
}

type aws_egress_only_internet_gateway struct {
	vpc_id string
}

type aws_sns_platform_application struct {
	name                             string
	platform                         string
	platform_credential              string
	event_endpoint_deleted_topic_arn *string
	success_feedback_role_arn        *string
	success_feedback_sample_rate     *string
	arn                              string
	event_delivery_failure_topic_arn *string
	event_endpoint_created_topic_arn *string
	event_endpoint_updated_topic_arn *string
	failure_feedback_role_arn        *string
	platform_principal               *string
}

type aws_sns_topic struct {
	name                                  *string
	display_name                          *string
	name_prefix                           *string
	policy                                *string
	delivery_policy                       *string
	application_success_feedback_role_arn *string
	lambda_success_feedback_role_arn      *string
	sqs_failure_feedback_role_arn         *string
	application_failure_feedback_role_arn *string
	http_success_feedback_role_arn        *string
	kms_master_key_id                     *string
	lambda_failure_feedback_role_arn      *string
	sqs_success_feedback_role_arn         *string
	arn                                   string
	http_failure_feedback_role_arn        *string
}

type aws_waf_regex_match_set struct {
	name string
}

type aws_api_gateway_model struct {
	rest_api_id  string
	name         string
	description  *string
	schema       *string
	content_type string
}

type aws_lambda_alias struct {
	description      *string
	function_name    string
	function_version string
	name             string
	arn              string
	invoke_arn       string
}

type aws_neptune_subnet_group struct {
	arn         string
	name        *string
	name_prefix *string
	description *string
	tags        *map[string]interface{}
}

type aws_service_discovery_public_dns_namespace struct {
	name        string
	description *string
	arn         string
	hosted_zone string
}

type aws_codepipeline_webhook struct {
	authentication  string
	name            string
	url             string
	target_action   string
	target_pipeline string
}

type aws_internet_gateway struct {
	owner_id string
	vpc_id   *string
	tags     *map[string]interface{}
}

type aws_opsworks_custom_layer struct {
	auto_assign_elastic_ips     *bool
	auto_assign_public_ips      *bool
	elastic_load_balancer       *string
	use_ebs_optimized_instances *bool
	short_name                  string
	custom_json                 *string
	auto_healing                *bool
	drain_elb_on_shutdown       *bool
	name                        string
	custom_instance_profile_arn *string
	install_updates_on_boot     *bool
	stack_id                    string
}

type aws_cloudwatch_dashboard struct {
	dashboard_arn  string
	dashboard_body string
	dashboard_name string
}

type aws_db_option_group struct {
	engine_name              string
	major_engine_version     string
	option_group_description *string
	tags                     *map[string]interface{}
	arn                      string
	name                     *string
	name_prefix              *string
}

type aws_ec2_transit_gateway_route_table_propagation struct {
	resource_id                    string
	resource_type                  string
	transit_gateway_attachment_id  string
	transit_gateway_route_table_id string
}

type aws_mq_configuration struct {
	description    *string
	engine_type    string
	engine_version string
	name           string
	arn            string
	data           string
}

type aws_glue_classifier struct {
	name string
}

type aws_sfn_state_machine struct {
	definition    string
	name          string
	role_arn      string
	creation_date string
	status        string
	tags          *map[string]interface{}
}

type aws_elb struct {
	name_prefix               *string
	source_security_group     *string
	source_security_group_id  string
	zone_id                   string
	cross_zone_load_balancing *bool
	dns_name                  string
	arn                       string
	tags                      *map[string]interface{}
	name                      *string
	internal                  *bool
	connection_draining       *bool
}

type aws_lb_cookie_stickiness_policy struct {
	name          string
	load_balancer string
}

type aws_vpn_gateway_route_propagation struct {
	vpn_gateway_id string
	route_table_id string
}

type aws_wafregional_rule struct {
	name        string
	metric_name string
}

type aws_lightsail_static_ip_attachment struct {
	static_ip_name string
	instance_name  string
}

type aws_load_balancer_listener_policy struct {
	load_balancer_name string
}

type aws_main_route_table_association struct {
	vpc_id                  string
	route_table_id          string
	original_route_table_id string
}

type aws_route53_health_check struct {
	insufficient_data_health_status *string
	resource_type                   string
	fqdn                            *string
	invert_healthcheck              *bool
	resource_path                   *string
	enable_sni                      *bool
	tags                            *map[string]interface{}
	cloudwatch_alarm_region         *string
	reference_name                  *string
	search_string                   *string
	measure_latency                 *bool
	ip_address                      *string
	cloudwatch_alarm_name           *string
}

type aws_vpc_peering_connection_accepter struct {
	accept_status             string
	auto_accept               *bool
	vpc_id                    string
	peer_vpc_id               string
	peer_owner_id             string
	peer_region               string
	tags                      *map[string]interface{}
	vpc_peering_connection_id string
}

type aws_wafregional_rate_based_rule struct {
	name        string
	metric_name string
	rate_key    string
}

type aws_datasync_task struct {
	arn                      string
	cloudwatch_log_group_arn *string
	destination_location_arn string
	name                     *string
	source_location_arn      string
	tags                     *map[string]interface{}
}

type aws_route53_zone_association struct {
	zone_id    string
	vpc_id     string
	vpc_region *string
}

type aws_ssm_parameter struct {
	resource_type   string
	value           string
	overwrite       *bool
	allowed_pattern *string
	tags            *map[string]interface{}
	name            string
	description     *string
	arn             *string
	key_id          *string
}

type aws_storagegateway_gateway struct {
	gateway_timezone    string
	gateway_type        *string
	medium_changer_type *string
	arn                 string
	activation_key      *string
	gateway_id          string
	gateway_ip_address  *string
	gateway_name        string
	smb_guest_password  *string
	tape_drive_type     *string
}

type aws_pinpoint_baidu_channel struct {
	application_id string
	enabled        *bool
	api_key        string
	secret_key     string
}

type aws_pinpoint_sms_channel struct {
	application_id string
	enabled        *bool
	sender_id      *string
	short_code     *string
}

type aws_db_instance struct {
	auto_minor_version_upgrade          *bool
	name                                *string
	password                            *string
	resource_id                         string
	iam_database_authentication_enabled *bool
	backup_window                       *string
	maintenance_window                  *string
	parameter_group_name                *string
	address                             string
	snapshot_identifier                 *string
	kms_key_id                          *string
	identifier_prefix                   *string
	publicly_accessible                 *bool
	replicate_source_db                 *string
	license_model                       *string
	monitoring_role_arn                 *string
	deletion_protection                 *bool
	engine_version                      *string
	storage_type                        *string
	instance_class                      string
	multi_az                            *bool
	domain_iam_role_name                *string
	availability_zone                   *string
	status                              string
	allow_major_version_upgrade         *bool
	engine                              *string
	identifier                          *string
	copy_tags_to_snapshot               *bool
	apply_immediately                   *bool
	tags                                *map[string]interface{}
	endpoint                            string
	ca_cert_identifier                  string
	domain                              *string
	storage_encrypted                   *bool
	skip_final_snapshot                 *bool
	timezone                            *string
	username                            *string
	db_subnet_group_name                *string
	hosted_zone_id                      string
	character_set_name                  *string
	arn                                 string
	final_snapshot_identifier           *string
	option_group_name                   *string
}

type aws_devicefarm_project struct {
	arn  string
	name string
}

type aws_eip struct {
	allocation_id             string
	domain                    string
	public_ip                 string
	private_ip                string
	associate_with_private_ip *string
	public_ipv4_pool          *string
	vpc                       *bool
	network_interface         *string
	tags                      *map[string]interface{}
	instance                  *string
	association_id            string
}

type aws_glue_catalog_table struct {
	catalog_id         *string
	database_name      string
	description        *string
	name               string
	owner              *string
	parameters         *map[string]interface{}
	table_type         *string
	view_original_text *string
	view_expanded_text *string
}

type aws_ecs_task_definition struct {
	task_role_arn         *string
	ipc_mode              *string
	arn                   string
	execution_role_arn    *string
	tags                  *map[string]interface{}
	container_definitions string
	pid_mode              *string
	cpu                   *string
	family                string
	memory                *string
	network_mode          *string
}

type aws_elasticache_cluster struct {
	configuration_endpoint string
	snapshot_name          *string
	subnet_group_name      *string
	parameter_group_name   *string
	apply_immediately      *bool
	maintenance_window     *string
	node_type              *string
	notification_topic_arn *string
	replication_group_id   *string
	tags                   *map[string]interface{}
	availability_zone      *string
	cluster_id             string
	engine                 *string
	snapshot_window        *string
	az_mode                *string
	cluster_address        string
	engine_version         *string
}

type aws_rds_cluster struct {
	cluster_identifier                  *string
	cluster_identifier_prefix           *string
	endpoint                            string
	storage_encrypted                   *bool
	iam_database_authentication_enabled *bool
	database_name                       *string
	apply_immediately                   *bool
	replication_source_identifier       *string
	db_subnet_group_name                *string
	global_cluster_identifier           *string
	hosted_zone_id                      string
	source_region                       *string
	engine_mode                         *string
	engine_version                      *string
	snapshot_identifier                 *string
	arn                                 string
	engine                              *string
	skip_final_snapshot                 *bool
	master_username                     *string
	kms_key_id                          *string
	cluster_resource_id                 string
	deletion_protection                 *bool
	reader_endpoint                     string
	master_password                     *string
	preferred_backup_window             *string
	preferred_maintenance_window        *string
	tags                                *map[string]interface{}
	db_cluster_parameter_group_name     *string
	final_snapshot_identifier           *string
}

type aws_sfn_activity struct {
	creation_date string
	tags          *map[string]interface{}
	name          string
}

type aws_dms_endpoint struct {
	endpoint_id                 string
	service_access_role         *string
	endpoint_type               string
	certificate_arn             *string
	extra_connection_attributes *string
	ssl_mode                    *string
	username                    *string
	engine_name                 string
	kms_key_arn                 *string
	server_name                 *string
	database_name               *string
	endpoint_arn                string
	password                    *string
	tags                        *map[string]interface{}
}

type aws_iam_server_certificate struct {
	certificate_body  string
	certificate_chain *string
	path              *string
	private_key       string
	name              *string
	name_prefix       *string
	arn               *string
}

type aws_iam_user_group_membership struct {
	user string
}

type aws_inspector_assessment_template struct {
	name       string
	target_arn string
	arn        string
}

type aws_api_gateway_resource struct {
	path        string
	rest_api_id string
	parent_id   string
	path_part   string
}

type aws_elastictranscoder_pipeline struct {
	aws_kms_key_arn *string
	name            *string
	role            string
	output_bucket   *string
	arn             string
	input_bucket    string
}

type aws_subnet struct {
	assign_ipv6_address_on_creation *bool
	ipv6_cidr_block_association_id  string
	tags                            *map[string]interface{}
	vpc_id                          string
	availability_zone_id            *string
	map_public_ip_on_launch         *bool
	arn                             string
	owner_id                        string
	cidr_block                      string
	ipv6_cidr_block                 *string
	availability_zone               *string
}

type aws_waf_size_constraint_set struct {
	name string
}

type aws_opsworks_nodejs_app_layer struct {
	auto_assign_elastic_ips     *bool
	use_ebs_optimized_instances *bool
	stack_id                    string
	name                        *string
	auto_assign_public_ips      *bool
	elastic_load_balancer       *string
	auto_healing                *bool
	install_updates_on_boot     *bool
	drain_elb_on_shutdown       *bool
	nodejs_version              *string
	custom_instance_profile_arn *string
	custom_json                 *string
}

type aws_ses_event_destination struct {
	name                   string
	configuration_set_name string
	enabled                *bool
}

type aws_ssm_patch_group struct {
	baseline_id string
	patch_group string
}

type aws_api_gateway_documentation_version struct {
	version     string
	rest_api_id string
	description *string
}

type aws_api_gateway_stage struct {
	cache_cluster_size    *string
	deployment_id         string
	description           *string
	documentation_version *string
	rest_api_id           string
	stage_name            string
	variables             *map[string]interface{}
	client_certificate_id *string
	execution_arn         string
	xray_tracing_enabled  *bool
	cache_cluster_enabled *bool
	invoke_url            string
	tags                  *map[string]interface{}
}

type aws_cloudwatch_event_rule struct {
	name_prefix         *string
	schedule_expression *string
	event_pattern       *string
	description         *string
	role_arn            *string
	is_enabled          *bool
	arn                 string
	name                *string
}

type aws_organizations_policy_attachment struct {
	policy_id string
	target_id string
}

type aws_wafregional_web_acl_association struct {
	web_acl_id   string
	resource_arn string
}

type aws_api_gateway_method_settings struct {
	rest_api_id string
	stage_name  string
	method_path string
}

type aws_iot_certificate struct {
	csr    string
	active bool
	arn    string
}

type aws_lambda_function struct {
	s3_object_version *string
	role              string
	function_name     string
	arn               string
	handler           string
	publish           *bool
	kms_key_arn       *string
	filename          *string
	s3_key            *string
	version           string
	last_modified     string
	s3_bucket         *string
	runtime           string
	invoke_arn        string
	tags              *map[string]interface{}
	description       *string
	source_code_hash  *string
	qualified_arn     string
}

type aws_waf_rule struct {
	name        string
	metric_name string
}

type aws_gamelift_game_session_queue struct {
	name string
	arn  string
}

type aws_iam_access_key struct {
	user              string
	status            string
	secret            string
	ses_smtp_password string
	pgp_key           *string
	key_fingerprint   string
	encrypted_secret  string
}

type aws_iam_openid_connect_provider struct {
	arn string
	url string
}

type aws_iam_role struct {
	description           *string
	assume_role_policy    string
	force_detach_policies *bool
	create_date           string
	arn                   string
	unique_id             string
	name                  *string
	name_prefix           *string
	path                  *string
	permissions_boundary  *string
	tags                  *map[string]interface{}
}

type aws_cognito_user_pool_domain struct {
	domain                      string
	certificate_arn             *string
	user_pool_id                string
	aws_account_id              string
	cloudfront_distribution_arn string
	s3_bucket                   string
	version                     string
}

type aws_efs_mount_target struct {
	file_system_id       string
	ip_address           *string
	subnet_id            string
	network_interface_id string
	dns_name             string
	file_system_arn      string
}

type aws_ssm_document struct {
	tags            *map[string]interface{}
	default_version string
	description     string
	hash_type       string
	latest_version  string
	name            string
	schema_version  string
	permissions     *map[string]interface{}
	hash            string
	owner           string
	document_format *string
	document_type   string
	created_date    string
	status          string
	arn             string
	content         string
}

type aws_dx_gateway_association struct {
	dx_gateway_id  string
	vpn_gateway_id string
}

type aws_secretsmanager_secret_version struct {
	secret_string *string
	secret_binary *string
	version_id    string
	arn           string
	secret_id     string
}

type aws_autoscaling_schedule struct {
	end_time               *string
	recurrence             *string
	arn                    string
	scheduled_action_name  string
	autoscaling_group_name string
	start_time             *string
}

type aws_budgets_budget struct {
	account_id        *string
	name_prefix       *string
	budget_type       string
	limit_unit        string
	time_period_start string
	time_unit         string
	cost_filters      *map[string]interface{}
	name              *string
	limit_amount      string
	time_period_end   *string
}

type aws_guardduty_member struct {
	detector_id                string
	email                      string
	relationship_status        string
	invite                     *bool
	disable_email_notification *bool
	invitation_message         *string
	account_id                 string
}

type aws_opsworks_mysql_layer struct {
	root_password                  *string
	custom_instance_profile_arn    *string
	use_ebs_optimized_instances    *bool
	name                           *string
	auto_assign_elastic_ips        *bool
	elastic_load_balancer          *string
	custom_json                    *string
	root_password_on_all_instances *bool
	install_updates_on_boot        *bool
	drain_elb_on_shutdown          *bool
	stack_id                       string
	auto_assign_public_ips         *bool
	auto_healing                   *bool
}

type aws_glacier_vault_lock struct {
	complete_lock         bool
	ignore_deletion_error *bool
	policy                string
	vault_name            string
}

type aws_lambda_layer_version struct {
	version           string
	layer_name        string
	s3_object_version *string
	created_date      string
	s3_bucket         *string
	description       *string
	layer_arn         string
	license_info      *string
	source_code_hash  *string
	filename          *string
	s3_key            *string
	arn               string
}

type aws_rds_cluster_endpoint struct {
	endpoint                    string
	arn                         string
	cluster_endpoint_identifier string
	cluster_identifier          string
	custom_endpoint_type        string
}

type aws_pinpoint_apns_sandbox_channel struct {
	default_authentication_method *string
	enabled                       *bool
	team_id                       *string
	token_key_id                  *string
	application_id                string
	bundle_id                     *string
	certificate                   *string
	private_key                   *string
	token_key                     *string
}

type aws_ec2_fleet struct {
	excess_capacity_termination_policy  *string
	terminate_instances                 *bool
	resource_type                       *string
	replace_unhealthy_instances         *bool
	tags                                *map[string]interface{}
	terminate_instances_with_expiration *bool
}

type aws_waf_rule_group struct {
	name        string
	metric_name string
}

type aws_dx_hosted_public_virtual_interface struct {
	customer_address *string
	owner_account_id string
	arn              string
	connection_id    string
	address_family   string
	amazon_address   *string
	name             string
	bgp_auth_key     *string
}

type aws_load_balancer_backend_server_policy struct {
	load_balancer_name string
}

type aws_licensemanager_license_configuration struct {
	description              *string
	license_count_hard_limit *bool
	license_counting_type    string
	name                     string
	tags                     *map[string]interface{}
}

type aws_default_vpc_dhcp_options struct {
	domain_name_servers string
	ntp_servers         string
	netbios_node_type   *string
	tags                *map[string]interface{}
	owner_id            string
	domain_name         string
}

type aws_codedeploy_deployment_config struct {
	compute_platform       *string
	deployment_config_id   string
	deployment_config_name string
}

type aws_dlm_lifecycle_policy struct {
	description        string
	execution_role_arn string
	state              *string
}

type aws_dx_lag struct {
	arn                   string
	name                  string
	connections_bandwidth string
	location              string
	force_destroy         *bool
	tags                  *map[string]interface{}
}

type aws_sagemaker_notebook_instance struct {
	tags          *map[string]interface{}
	arn           string
	name          string
	role_arn      string
	instance_type string
	subnet_id     *string
	kms_key_id    *string
}

type aws_cloudwatch_log_destination_policy struct {
	destination_name string
	access_policy    string
}

type aws_db_subnet_group struct {
	tags        *map[string]interface{}
	arn         string
	name        *string
	name_prefix *string
	description *string
}

type aws_gamelift_fleet struct {
	ec2_instance_type                  string
	name                               string
	operating_system                   string
	arn                                string
	description                        *string
	new_game_session_protection_policy *string
	build_id                           string
}

type aws_cloudwatch_log_resource_policy struct {
	policy_name     string
	policy_document string
}

type aws_elastic_beanstalk_environment struct {
	poll_interval          *string
	tags                   *map[string]interface{}
	tier                   *string
	platform_arn           *string
	cname_prefix           *string
	solution_stack_name    *string
	arn                    string
	version_label          *string
	cname                  string
	application            string
	wait_for_ready_timeout *string
	name                   string
	description            *string
	template_name          *string
}

type aws_ses_identity_notification_topic struct {
	topic_arn         *string
	notification_type string
	identity          string
}

type aws_appmesh_route struct {
	name                string
	mesh_name           string
	virtual_router_name string
	arn                 string
	created_date        string
	last_updated_date   string
}

type aws_cognito_identity_pool struct {
	allow_unauthenticated_identities *bool
	supported_login_providers        *map[string]interface{}
	identity_pool_name               string
	arn                              string
	developer_provider_name          *string
}

type aws_guardduty_detector struct {
	enable                       *bool
	account_id                   string
	finding_publishing_frequency *string
}

type aws_service_discovery_service struct {
	description *string
	arn         string
	name        string
}

type aws_storagegateway_upload_buffer struct {
	gateway_arn string
	disk_id     string
}

type aws_sns_topic_subscription struct {
	protocol               string
	topic_arn              string
	delivery_policy        *string
	filter_policy          *string
	endpoint               string
	endpoint_auto_confirms *bool
	raw_message_delivery   *bool
	arn                    string
}

type aws_ami_launch_permission struct {
	image_id   string
	account_id string
}

type aws_athena_database struct {
	name          string
	bucket        string
	force_destroy *bool
}

type aws_elasticsearch_domain_policy struct {
	access_policies string
	domain_name     string
}

type aws_iot_thing_type struct {
	deprecated *bool
	arn        string
	name       string
}

type aws_elastic_beanstalk_configuration_template struct {
	solution_stack_name *string
	name                string
	application         string
	description         *string
	environment_id      *string
}

type aws_iam_role_policy struct {
	policy      string
	name        *string
	name_prefix *string
	role        string
}

type aws_network_acl_rule struct {
	network_acl_id  string
	egress          *bool
	protocol        string
	ipv6_cidr_block *string
	icmp_code       *string
	rule_action     string
	cidr_block      *string
	icmp_type       *string
}

type aws_kinesis_firehose_delivery_stream struct {
	arn            *string
	version_id     *string
	destination_id *string
	tags           *map[string]interface{}
	destination    string
	name           string
}

type aws_s3_account_public_access_block struct {
	ignore_public_acls      *bool
	restrict_public_buckets *bool
	account_id              *string
	block_public_acls       *bool
	block_public_policy     *bool
}

type aws_rds_global_cluster struct {
	arn                        string
	database_name              *string
	deletion_protection        *bool
	engine                     *string
	engine_version             *string
	global_cluster_identifier  string
	global_cluster_resource_id string
	storage_encrypted          *bool
}

type aws_api_gateway_request_validator struct {
	rest_api_id                 string
	name                        string
	validate_request_body       *bool
	validate_request_parameters *bool
}

type aws_dx_public_virtual_interface struct {
	name             string
	arn              string
	bgp_auth_key     *string
	address_family   string
	customer_address *string
	amazon_address   *string
	tags             *map[string]interface{}
	connection_id    string
}

type aws_elastic_beanstalk_application_version struct {
	key          string
	name         string
	force_delete *bool
	application  string
	description  *string
	bucket       string
}

type aws_opsworks_php_app_layer struct {
	auto_healing                *bool
	drain_elb_on_shutdown       *bool
	stack_id                    string
	custom_instance_profile_arn *string
	elastic_load_balancer       *string
	use_ebs_optimized_instances *bool
	name                        *string
	auto_assign_elastic_ips     *bool
	auto_assign_public_ips      *bool
	custom_json                 *string
	install_updates_on_boot     *bool
}

type aws_appsync_graphql_api struct {
	arn                 string
	uris                map[string]interface{}
	authentication_type string
	name                string
}

type aws_autoscaling_group struct {
	placement_group           *string
	force_delete              *bool
	metrics_granularity       *string
	service_linked_role_arn   *string
	arn                       string
	health_check_type         *string
	wait_for_capacity_timeout *string
	name                      *string
	launch_configuration      *string
	protect_from_scale_in     *bool
	name_prefix               *string
}

type aws_directory_service_directory struct {
	enable_sso        *bool
	password          string
	alias             *string
	short_name        *string
	name              string
	tags              *map[string]interface{}
	access_url        string
	resource_type     *string
	edition           *string
	size              *string
	security_group_id string
	description       *string
}

type aws_iam_user_policy_attachment struct {
	user       string
	policy_arn string
}

type aws_s3_bucket_metric struct {
	bucket string
	name   string
}

type aws_alb_listener struct {
	ssl_policy        *string
	certificate_arn   *string
	arn               string
	load_balancer_arn string
	protocol          *string
}

type aws_ami_from_instance struct {
	tags                    *map[string]interface{}
	description             *string
	kernel_id               string
	source_instance_id      string
	sriov_net_support       string
	virtualization_type     string
	architecture            string
	manage_ebs_snapshots    bool
	name                    string
	image_location          string
	ramdisk_id              string
	root_device_name        string
	root_snapshot_id        string
	snapshot_without_reboot *bool
	ena_support             bool
}

type aws_db_event_subscription struct {
	customer_aws_id string
	name            *string
	sns_topic       string
	source_type     *string
	enabled         *bool
	tags            *map[string]interface{}
	arn             string
	name_prefix     *string
}

type aws_lb_listener struct {
	load_balancer_arn string
	protocol          *string
	ssl_policy        *string
	certificate_arn   *string
	arn               string
}

type aws_config_configuration_recorder_status struct {
	is_enabled bool
	name       string
}

type aws_db_snapshot struct {
	vpc_id                        string
	db_instance_identifier        string
	engine                        string
	source_region                 string
	tags                          *map[string]interface{}
	db_snapshot_identifier        string
	availability_zone             string
	encrypted                     bool
	option_group_name             string
	source_db_snapshot_identifier string
	snapshot_type                 string
	status                        string
	storage_type                  string
	db_snapshot_arn               string
	engine_version                string
	kms_key_id                    string
	license_model                 string
}

type aws_vpc_peering_connection_options struct {
	vpc_peering_connection_id string
}

type aws_wafregional_byte_match_set struct {
	name string
}

type aws_route53_record struct {
	health_check_id                  *string
	fqdn                             string
	multivalue_answer_routing_policy *bool
	allow_overwrite                  *bool
	zone_id                          string
	resource_type                    string
	set_identifier                   *string
	name                             string
	failover                         *string
}

type aws_default_security_group struct {
	arn                    string
	owner_id               string
	tags                   *map[string]interface{}
	revoke_rules_on_delete *bool
	name                   string
	vpc_id                 *string
}

type aws_transfer_ssh_key struct {
	body      string
	server_id string
	user_name string
}

type aws_lb_target_group_attachment struct {
	target_group_arn  string
	target_id         string
	availability_zone *string
}

type aws_storagegateway_cached_iscsi_volume struct {
	network_interface_id string
	target_name          string
	source_volume_arn    *string
	target_arn           string
	volume_arn           string
	arn                  string
	chap_enabled         bool
	gateway_arn          string
	snapshot_id          *string
	volume_id            string
}

type aws_api_gateway_usage_plan_key struct {
	key_id        string
	key_type      string
	usage_plan_id string
	name          string
	value         string
}

type aws_flow_log struct {
	log_destination      *string
	log_destination_type *string
	log_group_name       *string
	vpc_id               *string
	subnet_id            *string
	eni_id               *string
	traffic_type         string
	iam_role_arn         *string
}

type aws_iot_topic_rule struct {
	name        string
	description *string
	enabled     bool
	sql         string
	sql_version string
	arn         string
}

type aws_opsworks_memcached_layer struct {
	elastic_load_balancer       *string
	auto_healing                *bool
	install_updates_on_boot     *bool
	stack_id                    string
	use_ebs_optimized_instances *bool
	auto_assign_public_ips      *bool
	custom_instance_profile_arn *string
	drain_elb_on_shutdown       *bool
	auto_assign_elastic_ips     *bool
	custom_json                 *string
	name                        *string
}

type aws_elasticache_security_group struct {
	description *string
	name        string
}

type aws_neptune_cluster_parameter_group struct {
	arn         string
	name        *string
	name_prefix *string
	family      string
	description *string
	tags        *map[string]interface{}
}

type aws_transfer_user struct {
	arn            string
	home_directory *string
	policy         *string
	role           string
	server_id      string
	tags           *map[string]interface{}
	user_name      string
}

type aws_wafregional_regex_match_set struct {
	name string
}

type aws_pinpoint_app struct {
	application_id string
	name           *string
	name_prefix    *string
}

type aws_swf_domain struct {
	name                                        *string
	name_prefix                                 *string
	description                                 *string
	workflow_execution_retention_period_in_days string
}

type aws_ami struct {
	description          *string
	kernel_id            *string
	sriov_net_support    *string
	virtualization_type  *string
	tags                 *map[string]interface{}
	manage_ebs_snapshots bool
	root_device_name     *string
	architecture         *string
	name                 string
	root_snapshot_id     string
	image_location       *string
	ena_support          *bool
	ramdisk_id           *string
}

type aws_codecommit_trigger struct {
	repository_name  string
	configuration_id string
}

type aws_ecs_service struct {
	name                    string
	launch_type             *string
	propagate_tags          *string
	iam_role                *string
	task_definition         string
	enable_ecs_managed_tags *bool
	platform_version        *string
	tags                    *map[string]interface{}
	cluster                 *string
	scheduling_strategy     *string
}

type aws_glue_security_configuration struct {
	name string
}

type aws_default_vpc struct {
	main_route_table_id              string
	default_network_acl_id           string
	dhcp_options_id                  string
	arn                              string
	instance_tenancy                 string
	enable_dns_hostnames             *bool
	enable_classiclink               *bool
	default_security_group_id        string
	ipv6_cidr_block                  string
	tags                             *map[string]interface{}
	enable_dns_support               *bool
	assign_generated_ipv6_cidr_block bool
	default_route_table_id           string
	cidr_block                       string
	enable_classiclink_dns_support   *bool
	ipv6_association_id              string
	owner_id                         string
}

type aws_ec2_capacity_reservation struct {
	availability_zone       string
	ebs_optimized           *bool
	end_date_type           *string
	instance_match_criteria *string
	tags                    *map[string]interface{}
	tenancy                 *string
	end_date                *string
	ephemeral_storage       *bool
	instance_platform       string
	instance_type           string
}

type aws_macie_member_account_association struct {
	member_account_id string
}

type aws_proxy_protocol_policy struct {
	load_balancer string
}

type aws_ssm_activation struct {
	expiration_date *string
	iam_role        string
	activation_code string
	name            *string
	description     *string
	expired         string
}

type aws_s3_bucket struct {
	tags                        *map[string]interface{}
	policy                      *string
	arn                         *string
	hosted_zone_id              *string
	region                      *string
	website_endpoint            *string
	bucket_prefix               *string
	bucket_domain_name          string
	request_payer               *string
	acl                         *string
	website_domain              *string
	force_destroy               *bool
	acceleration_status         *string
	bucket                      *string
	bucket_regional_domain_name string
}

type aws_launch_template struct {
	image_id                             *string
	key_name                             *string
	user_data                            *string
	arn                                  string
	tags                                 *map[string]interface{}
	ebs_optimized                        *string
	description                          *string
	instance_type                        *string
	kernel_id                            *string
	name                                 *string
	name_prefix                          *string
	disable_api_termination              *bool
	ram_disk_id                          *string
	instance_initiated_shutdown_behavior *string
}

type aws_opsworks_haproxy_layer struct {
	auto_assign_elastic_ips     *bool
	healthcheck_url             *string
	stats_user                  *string
	custom_instance_profile_arn *string
	elastic_load_balancer       *string
	auto_healing                *bool
	install_updates_on_boot     *bool
	use_ebs_optimized_instances *bool
	auto_assign_public_ips      *bool
	custom_json                 *string
	stack_id                    string
	name                        *string
	stats_password              string
	healthcheck_method          *string
	drain_elb_on_shutdown       *bool
	stats_enabled               *bool
	stats_url                   *string
}

type aws_acm_certificate struct {
	certificate_body  *string
	domain_name       *string
	arn               string
	validation_method *string
	tags              *map[string]interface{}
	certificate_chain *string
	private_key       *string
}

type aws_dms_replication_task struct {
	replication_task_settings *string
	source_endpoint_arn       string
	tags                      *map[string]interface{}
	cdc_start_time            *string
	replication_instance_arn  string
	replication_task_arn      string
	target_endpoint_arn       string
	migration_type            string
	replication_task_id       string
	table_mappings            string
}

type aws_emr_instance_group struct {
	status        string
	name          *string
	ebs_optimized *bool
	cluster_id    string
	instance_type string
}

type aws_glue_trigger struct {
	description   *string
	enabled       *bool
	name          string
	schedule      *string
	resource_type string
}

type aws_volume_attachment struct {
	device_name  string
	instance_id  string
	volume_id    string
	force_detach *bool
	skip_destroy *bool
}

type aws_waf_rate_based_rule struct {
	rate_key    string
	name        string
	metric_name string
}

type aws_datasync_location_s3 struct {
	subdirectory  string
	tags          *map[string]interface{}
	uri           string
	arn           string
	s3_bucket_arn string
}

type aws_elasticache_parameter_group struct {
	name        string
	family      string
	description *string
}

type aws_kms_alias struct {
	name_prefix    *string
	target_key_id  string
	target_key_arn string
	arn            string
	name           *string
}

type aws_service_discovery_private_dns_namespace struct {
	name        string
	description *string
	vpc         string
	arn         string
	hosted_zone string
}

type aws_cognito_identity_pool_roles_attachment struct {
	identity_pool_id string
	roles            map[string]interface{}
}

type aws_ssm_resource_data_sync struct {
	name string
}

type aws_storagegateway_smb_file_share struct {
	arn                     string
	kms_key_arn             *string
	object_acl              *string
	role_arn                string
	default_storage_class   *string
	location_arn            string
	read_only               *bool
	authentication          *string
	fileshare_id            string
	gateway_arn             string
	guess_mime_type_enabled *bool
	kms_encrypted           *bool
	requester_pays          *bool
}

type aws_ses_domain_dkim struct {
	domain string
}

type aws_alb_target_group struct {
	name_prefix       *string
	vpc_id            *string
	tags              *map[string]interface{}
	arn_suffix        string
	name              *string
	protocol          *string
	arn               string
	proxy_protocol_v2 *bool
	target_type       *string
}

type aws_autoscaling_lifecycle_hook struct {
	lifecycle_transition    string
	notification_metadata   *string
	notification_target_arn *string
	role_arn                *string
	name                    string
	autoscaling_group_name  string
	default_result          *string
}

type aws_guardduty_ipset struct {
	detector_id string
	name        string
	format      string
	location    string
	activate    bool
}

type aws_launch_configuration struct {
	image_id                    string
	instance_type               string
	ebs_optimized               *bool
	vpc_classic_link_id         *string
	spot_price                  *string
	enable_monitoring           *bool
	name                        *string
	name_prefix                 *string
	iam_instance_profile        *string
	associate_public_ip_address *bool
	placement_tenancy           *string
	key_name                    *string
	user_data                   *string
	user_data_base64            *string
}

type aws_route53_delegation_set struct {
	reference_name *string
}

type aws_media_store_container struct {
	arn      string
	endpoint string
	name     string
}

type aws_vpn_gateway struct {
	amazon_side_asn   *string
	vpc_id            *string
	tags              *map[string]interface{}
	availability_zone *string
}

type aws_datasync_location_nfs struct {
	arn             string
	server_hostname string
	subdirectory    string
	tags            *map[string]interface{}
	uri             string
}

type aws_neptune_cluster_instance struct {
	publicly_accessible          *bool
	cluster_identifier           string
	engine                       *string
	identifier_prefix            *string
	neptune_subnet_group_name    *string
	address                      string
	apply_immediately            *bool
	availability_zone            *string
	kms_key_arn                  string
	neptune_parameter_group_name *string
	tags                         *map[string]interface{}
	auto_minor_version_upgrade   *bool
	dbi_resource_id              string
	identifier                   *string
	preferred_backup_window      *string
	storage_encrypted            bool
	writer                       bool
	arn                          string
	endpoint                     string
	engine_version               *string
	instance_class               string
	preferred_maintenance_window *string
}

type aws_network_interface_attachment struct {
	attachment_id        string
	status               string
	instance_id          string
	network_interface_id string
}

type aws_opsworks_user_profile struct {
	user_arn              string
	allow_self_management *bool
	ssh_username          string
	ssh_public_key        *string
}

type aws_vpc_endpoint struct {
	policy              *string
	state               string
	vpc_endpoint_type   *string
	service_name        string
	private_dns_enabled *bool
	vpc_id              string
	prefix_list_id      string
	auto_accept         *bool
}

type aws_appsync_api_key struct {
	description *string
	api_id      string
	expires     *string
	key         string
}

type aws_autoscaling_policy struct {
	name                    string
	arn                     string
	adjustment_type         *string
	autoscaling_group_name  string
	policy_type             *string
	metric_aggregation_type *string
}

type aws_dax_subnet_group struct {
	name        string
	description *string
	vpc_id      string
}

type aws_storagegateway_cache struct {
	disk_id     string
	gateway_arn string
}

type aws_vpc_ipv4_cidr_block_association struct {
	vpc_id     string
	cidr_block string
}

type aws_wafregional_ipset struct {
	name string
	arn  string
}

type aws_organizations_organization struct {
	master_account_arn   string
	master_account_email string
	master_account_id    string
	feature_set          *string
	arn                  string
}

type aws_appautoscaling_target struct {
	resource_id        string
	role_arn           *string
	scalable_dimension string
	service_namespace  string
}

type aws_datasync_agent struct {
	arn            string
	activation_key *string
	ip_address     *string
	name           *string
	tags           *map[string]interface{}
}

type aws_ebs_snapshot struct {
	data_encryption_key_id string
	owner_alias            string
	encrypted              bool
	owner_id               string
	kms_key_id             string
	tags                   *map[string]interface{}
	volume_id              string
	description            *string
}

type aws_macie_s3_bucket_association struct {
	prefix            *string
	member_account_id *string
	bucket_name       string
}

type aws_nat_gateway struct {
	private_ip           string
	public_ip            string
	tags                 *map[string]interface{}
	allocation_id        string
	subnet_id            string
	network_interface_id string
}

type aws_s3_bucket_object struct {
	key                    string
	storage_class          *string
	server_side_encryption *string
	version_id             string
	cache_control          *string
	content_type           *string
	etag                   *string
	bucket                 string
	acl                    *string
	content_disposition    *string
	content_encoding       *string
	content_language       *string
	source                 *string
	content                *string
	content_base64         *string
	website_redirect       *string
	kms_key_id             *string
	tags                   *map[string]interface{}
}

type aws_s3_bucket_notification struct {
	bucket string
}

type aws_pinpoint_email_channel struct {
	enabled        *bool
	from_address   string
	identity       string
	role_arn       string
	application_id string
}

type aws_gamelift_build struct {
	version          *string
	name             string
	operating_system string
}

type aws_kms_key struct {
	policy              *string
	is_enabled          *bool
	enable_key_rotation *bool
	key_usage           *string
	tags                *map[string]interface{}
	arn                 string
	key_id              string
	description         *string
}

type aws_lb_ssl_negotiation_policy struct {
	name          string
	load_balancer string
}

type aws_media_store_container_policy struct {
	policy         string
	container_name string
}

type aws_lb_target_group struct {
	name_prefix       *string
	vpc_id            *string
	arn               string
	name              *string
	protocol          *string
	target_type       *string
	tags              *map[string]interface{}
	arn_suffix        string
	proxy_protocol_v2 *bool
}

type aws_cloudwatch_log_destination struct {
	name       string
	role_arn   string
	target_arn string
	arn        string
}

type aws_codebuild_project struct {
	service_role   string
	tags           *map[string]interface{}
	arn            string
	name           string
	badge_enabled  *bool
	badge_url      string
	encryption_key *string
	description    *string
}

type aws_dms_replication_subnet_group struct {
	vpc_id                               string
	replication_subnet_group_arn         string
	replication_subnet_group_description string
	replication_subnet_group_id          string
	tags                                 *map[string]interface{}
}

type aws_api_gateway_gateway_response struct {
	rest_api_id         string
	response_type       string
	status_code         *string
	response_templates  *map[string]interface{}
	response_parameters *map[string]interface{}
}

type aws_opsworks_java_app_layer struct {
	auto_assign_public_ips      *bool
	name                        *string
	install_updates_on_boot     *bool
	jvm_type                    *string
	elastic_load_balancer       *string
	auto_healing                *bool
	jvm_version                 *string
	jvm_options                 *string
	auto_assign_elastic_ips     *bool
	custom_instance_profile_arn *string
	app_server                  *string
	app_server_version          *string
	custom_json                 *string
	drain_elb_on_shutdown       *bool
	stack_id                    string
	use_ebs_optimized_instances *bool
}

type aws_ses_receipt_filter struct {
	name   string
	cidr   string
	policy string
}

type aws_api_gateway_integration_response struct {
	status_code                 string
	selection_pattern           *string
	response_parameters         *map[string]interface{}
	content_handling            *string
	rest_api_id                 string
	resource_id                 string
	http_method                 string
	response_templates          *map[string]interface{}
	response_parameters_in_json *string
}

type aws_dms_certificate struct {
	certificate_wallet *string
	certificate_arn    string
	certificate_id     string
	certificate_pem    *string
}

type aws_iam_user_policy struct {
	user        string
	policy      string
	name        *string
	name_prefix *string
}

type aws_emr_security_configuration struct {
	name          *string
	name_prefix   *string
	configuration string
	creation_date string
}

type aws_securityhub_account struct {
}

type aws_vpc_peering_connection struct {
	accept_status string
	peer_owner_id *string
	peer_vpc_id   string
	vpc_id        string
	tags          *map[string]interface{}
	auto_accept   *bool
	peer_region   *string
}

type aws_customer_gateway struct {
	ip_address    string
	resource_type string
	tags          *map[string]interface{}
}

type aws_ssm_maintenance_window struct {
	name                       string
	schedule                   string
	allow_unassociated_targets *bool
	end_date                   *string
	start_date                 *string
	enabled                    *bool
	schedule_timezone          *string
}

type aws_spot_instance_request struct {
	iam_instance_profile                 *string
	block_device                         *map[string]interface{}
	password_data                        string
	private_ip                           *string
	instance_state                       string
	private_dns                          string
	disable_api_termination              *bool
	monitoring                           *bool
	tags                                 *map[string]interface{}
	ami                                  string
	arn                                  string
	availability_zone                    *string
	instance_initiated_shutdown_behavior *string
	get_password_data                    *bool
	host_id                              *string
	spot_instance_id                     string
	public_ip                            string
	spot_request_state                   string
	user_data_base64                     *string
	primary_network_interface_id         string
	tenancy                              *string
	wait_for_fulfillment                 *bool
	launch_group                         *string
	spot_bid_status                      string
	instance_type                        string
	subnet_id                            *string
	volume_tags                          *map[string]interface{}
	spot_type                            *string
	instance_interruption_behaviour      *string
	associate_public_ip_address          *bool
	placement_group                      *string
	key_name                             *string
	public_dns                           string
	network_interface_id                 string
	spot_price                           *string
	source_dest_check                    *bool
	user_data                            *string
	ebs_optimized                        *bool
	valid_from                           *string
	valid_until                          *string
}

type aws_vpn_connection struct {
	transit_gateway_id             *string
	static_routes_only             *bool
	tunnel1_inside_cidr            *string
	tunnel2_inside_cidr            *string
	tags                           *map[string]interface{}
	tunnel1_cgw_inside_address     string
	tunnel1_vgw_inside_address     string
	tunnel2_address                string
	tunnel2_vgw_inside_address     string
	resource_type                  string
	tunnel2_preshared_key          *string
	tunnel2_bgp_asn                string
	vpn_gateway_id                 *string
	tunnel1_preshared_key          *string
	tunnel1_address                string
	customer_gateway_id            string
	customer_gateway_configuration *string
	tunnel1_bgp_asn                string
	tunnel2_cgw_inside_address     string
}

type aws_alb_listener_rule struct {
	arn          string
	listener_arn string
}

type aws_api_gateway_domain_name struct {
	certificate_body          *string
	certificate_chain         *string
	certificate_arn           *string
	certificate_upload_date   string
	regional_certificate_name *string
	regional_zone_id          string
	regional_domain_name      string
	regional_certificate_arn  *string
	certificate_name          *string
	certificate_private_key   *string
	domain_name               string
	cloudfront_domain_name    string
	cloudfront_zone_id        string
}

type aws_elasticsearch_domain struct {
	access_policies       *string
	advanced_options      *map[string]interface{}
	kibana_endpoint       string
	domain_id             string
	endpoint              string
	domain_name           string
	arn                   string
	elasticsearch_version *string
	tags                  *map[string]interface{}
}

type aws_glue_connection struct {
	connection_properties map[string]interface{}
	connection_type       *string
	description           *string
	name                  string
	catalog_id            *string
}

type aws_opsworks_permission struct {
	allow_ssh  *bool
	allow_sudo *bool
	user_arn   string
	level      *string
	stack_id   *string
}

type aws_route_table_association struct {
	subnet_id      string
	route_table_id string
}

type aws_cloudtrail struct {
	s3_bucket_name                string
	kms_key_id                    *string
	arn                           string
	is_multi_region_trail         *bool
	tags                          *map[string]interface{}
	name                          string
	s3_key_prefix                 *string
	cloud_watch_logs_role_arn     *string
	cloud_watch_logs_group_arn    *string
	include_global_service_events *bool
	home_region                   string
	enable_logging                *bool
	is_organization_trail         *bool
	sns_topic_name                *string
	enable_log_file_validation    *bool
}

type aws_ecr_repository_policy struct {
	repository  string
	policy      string
	registry_id string
}

type aws_elastic_beanstalk_application struct {
	name        string
	description *string
}

type aws_iam_group struct {
	arn       string
	unique_id string
	name      string
	path      *string
}

type aws_ecs_cluster struct {
	name string
	tags *map[string]interface{}
	arn  string
}

type aws_api_gateway_method_response struct {
	rest_api_id                 string
	resource_id                 string
	http_method                 string
	status_code                 string
	response_models             *map[string]interface{}
	response_parameters         *map[string]interface{}
	response_parameters_in_json *string
}

type aws_dx_hosted_private_virtual_interface struct {
	amazon_address      *string
	jumbo_frame_capable bool
	arn                 string
	connection_id       string
	name                string
	address_family      string
	owner_account_id    string
	bgp_auth_key        *string
	customer_address    *string
}

type aws_lambda_event_source_mapping struct {
	enabled                     *bool
	last_modified               string
	state_transition_reason     string
	event_source_arn            string
	starting_position_timestamp *string
	function_arn                string
	last_processing_result      string
	state                       string
	uuid                        string
	function_name               string
	starting_position           *string
}

type aws_opsworks_static_web_layer struct {
	install_updates_on_boot     *bool
	stack_id                    string
	auto_assign_public_ips      *bool
	custom_json                 *string
	auto_assign_elastic_ips     *bool
	auto_healing                *bool
	use_ebs_optimized_instances *bool
	name                        *string
	drain_elb_on_shutdown       *bool
	custom_instance_profile_arn *string
	elastic_load_balancer       *string
}

type aws_cloudwatch_log_stream struct {
	name           string
	log_group_name string
	arn            string
}

type aws_route53_query_log struct {
	cloudwatch_log_group_arn string
	zone_id                  string
}

type aws_security_group_rule struct {
	security_group_id        string
	source_security_group_id *string
	protocol                 string
	self                     *bool
	description              *string
	resource_type            string
}

type aws_vpc_endpoint_service struct {
	state               string
	service_name        string
	service_type        string
	private_dns_name    string
	acceptance_required bool
}

type aws_pinpoint_apns_voip_sandbox_channel struct {
	bundle_id                     *string
	default_authentication_method *string
	token_key                     *string
	token_key_id                  *string
	application_id                string
	certificate                   *string
	enabled                       *bool
	private_key                   *string
	team_id                       *string
}

type aws_waf_xss_match_set struct {
	name string
}

type aws_waf_sql_injection_match_set struct {
	name string
}

type aws_api_gateway_method struct {
	resource_id                string
	http_method                string
	authorization              string
	authorizer_id              *string
	request_models             *map[string]interface{}
	request_parameters_in_json *string
	request_validator_id       *string
	rest_api_id                string
	api_key_required           *bool
	request_parameters         *map[string]interface{}
}

type aws_dx_hosted_public_virtual_interface_accepter struct {
	tags                 *map[string]interface{}
	arn                  string
	virtual_interface_id string
}

type aws_iam_group_policy_attachment struct {
	group      string
	policy_arn string
}

type aws_security_group struct {
	arn                    string
	owner_id               string
	name                   *string
	description            *string
	vpc_id                 *string
	tags                   *map[string]interface{}
	revoke_rules_on_delete *bool
	name_prefix            *string
}

type aws_codebuild_webhook struct {
	project_name  string
	branch_filter *string
	payload_url   string
	secret        string
	url           string
}

type aws_dx_bgp_peer struct {
	address_family       string
	virtual_interface_id string
	amazon_address       *string
	bgp_auth_key         *string
	customer_address     *string
	bgp_status           string
}

type aws_vpn_gateway_attachment struct {
	vpc_id         string
	vpn_gateway_id string
}

type aws_wafregional_regex_pattern_set struct {
	name string
}

type aws_default_network_acl struct {
	owner_id               string
	vpc_id                 string
	default_network_acl_id string
	tags                   *map[string]interface{}
}

type aws_vpc struct {
	default_route_table_id           string
	owner_id                         string
	instance_tenancy                 *string
	main_route_table_id              string
	default_network_acl_id           string
	enable_dns_support               *bool
	ipv6_association_id              string
	arn                              string
	enable_classiclink_dns_support   *bool
	assign_generated_ipv6_cidr_block *bool
	dhcp_options_id                  string
	default_security_group_id        string
	ipv6_cidr_block                  string
	cidr_block                       string
	enable_dns_hostnames             *bool
	enable_classiclink               *bool
	tags                             *map[string]interface{}
}

type aws_eip_association struct {
	public_ip            *string
	allocation_id        *string
	allow_reassociation  *bool
	instance_id          *string
	network_interface_id *string
	private_ip_address   *string
}

type aws_iam_instance_profile struct {
	arn         string
	create_date string
	unique_id   string
	name        *string
	name_prefix *string
	path        *string
	role        *string
}

type aws_lightsail_static_ip struct {
	name         string
	ip_address   string
	arn          string
	support_code string
}

type aws_securityhub_product_subscription struct {
	product_arn string
	arn         string
}

type aws_vpc_dhcp_options_association struct {
	dhcp_options_id string
	vpc_id          string
}

type aws_alb_listener_certificate struct {
	listener_arn    string
	certificate_arn string
}

type aws_dx_hosted_private_virtual_interface_accepter struct {
	tags                 *map[string]interface{}
	arn                  string
	virtual_interface_id string
	vpn_gateway_id       *string
	dx_gateway_id        *string
}

type aws_iam_user_ssh_key struct {
	ssh_public_key_id string
	fingerprint       string
	username          string
	public_key        string
	encoding          string
	status            *string
}

type aws_instance struct {
	placement_group                      *string
	subnet_id                            *string
	public_ip                            string
	iam_instance_profile                 *string
	source_dest_check                    *bool
	primary_network_interface_id         string
	network_interface_id                 string
	volume_tags                          *map[string]interface{}
	block_device                         *map[string]interface{}
	ami                                  string
	instance_type                        string
	password_data                        string
	associate_public_ip_address          *bool
	key_name                             *string
	public_dns                           string
	instance_state                       string
	monitoring                           *bool
	arn                                  string
	private_ip                           *string
	user_data_base64                     *string
	tenancy                              *string
	tags                                 *map[string]interface{}
	availability_zone                    *string
	private_dns                          string
	disable_api_termination              *bool
	instance_initiated_shutdown_behavior *string
	host_id                              *string
	get_password_data                    *bool
	user_data                            *string
	ebs_optimized                        *bool
}

type aws_lightsail_instance struct {
	arn                string
	username           string
	bundle_id          string
	key_pair_name      *string
	user_data          *string
	public_ip_address  string
	blueprint_id       string
	is_static_ip       bool
	private_ip_address string
	ipv6_address       string
	name               string
	availability_zone  string
	created_at         string
}
