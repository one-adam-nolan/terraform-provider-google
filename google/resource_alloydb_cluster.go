// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceAlloydbCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceAlloydbClusterCreate,
		Read:   resourceAlloydbClusterRead,
		Update: resourceAlloydbClusterUpdate,
		Delete: resourceAlloydbClusterDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAlloydbClusterImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The ID of the alloydb cluster.`,
			},
			"network": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: projectNumberDiffSuppress,
				Description: `The relative resource name of the VPC network on which the instance can be accessed. It is specified in the following form:

"projects/{projectNumber}/global/networks/{network_id}".`,
			},
			"automated_backup_policy": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `The automated backup policy for this cluster.

If no policy is provided then the default policy will be used. The default policy takes one backup a day, has a backup window of 1 hour, and retains backups for 14 days.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"weekly_schedule": {
							Type:        schema.TypeList,
							Required:    true,
							ForceNew:    true,
							Description: `Weekly schedule for the Backup.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"start_times": {
										Type:        schema.TypeList,
										Required:    true,
										Description: `The times during the day to start a backup. At least one start time must be provided. The start times are assumed to be in UTC and to be an exact hour (e.g., 04:00:00).`,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"hours": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: `Hours of day in 24 hour format. Should be from 0 to 23. An API may choose to allow the value "24:00:00" for scenarios like business closing time.`,
												},
												"minutes": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: `Minutes of hour of day. Must be from 0 to 59.`,
												},
												"nanos": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: `Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.`,
												},
												"seconds": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: `Seconds of minutes of the time. Must normally be from 0 to 59. An API may allow the value 60 if it allows leap-seconds.`,
												},
											},
										},
									},
									"days_of_week": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: `The days of the week to perform a backup. At least one day of the week must be provided. Possible values: ["MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY", "SUNDAY"]`,
										MinItems:    1,
										Elem: &schema.Schema{
											Type:         schema.TypeString,
											ValidateFunc: validateEnum([]string{"MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY", "SUNDAY"}),
										},
									},
								},
							},
						},
						"backup_window": {
							Type:     schema.TypeString,
							Optional: true,
							Description: `The length of the time window during which a backup can be taken. If a backup does not succeed within this time window, it will be canceled and considered failed.

The backup window must be at least 5 minutes long. There is no upper bound on the window. If not set, it will default to 1 hour.

A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".`,
						},
						"enabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: `Whether automated backups are enabled.`,
						},
						"labels": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: `Labels to apply to backups created using this configuration.`,
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"location": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `The location where the backup will be stored. Currently, the only supported option is to store the backup in the same region as the cluster.`,
						},
						"quantity_based_retention": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `Quantity-based Backup retention policy to retain recent backups.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"count": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: `The number of backups to retain.`,
									},
								},
							},
							ConflictsWith: []string{"automated_backup_policy.0.time_based_retention"},
						},
						"time_based_retention": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `Time-based Backup retention policy.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"retention_period": {
										Type:     schema.TypeString,
										Optional: true,
										Description: `The retention period.
A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".`,
									},
								},
							},
							ConflictsWith: []string{"automated_backup_policy.0.quantity_based_retention"},
						},
					},
				},
			},
			"display_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `User-settable and human-readable display name for the Cluster.`,
			},
			"initial_user": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: `Initial user to setup during cluster creation.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"password": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `The initial password for the user.`,
							Sensitive:   true,
						},
						"user": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `The database username.`,
						},
					},
				},
			},
			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: `User-defined labels for the alloydb cluster.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"location": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `The location where the alloydb cluster should reside.`,
			},
			"backup_source": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `Cluster created from backup.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"backup_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `The name of the backup resource.`,
						},
					},
				},
			},
			"database_version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The database engine major version. This is an output-only field and it's populated at the Cluster creation time. This field cannot be changed after cluster creation.`,
			},
			"migration_source": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `Cluster created via DMS migration.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"host_port": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `The host and port of the on-premises instance in host:port format`,
						},
						"reference_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `Place holder for the external source identifier(e.g DMS job name) that created the cluster.`,
						},
						"source_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `Type of migration source.`,
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The name of the cluster resource.`,
			},
			"uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The system-generated UID of the resource.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceAlloydbClusterCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	labelsProp, err := expandAlloydbClusterLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	networkProp, err := expandAlloydbClusterNetwork(d.Get("network"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("network"); !isEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	displayNameProp, err := expandAlloydbClusterDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	initialUserProp, err := expandAlloydbClusterInitialUser(d.Get("initial_user"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("initial_user"); !isEmptyValue(reflect.ValueOf(initialUserProp)) && (ok || !reflect.DeepEqual(v, initialUserProp)) {
		obj["initialUser"] = initialUserProp
	}
	automatedBackupPolicyProp, err := expandAlloydbClusterAutomatedBackupPolicy(d.Get("automated_backup_policy"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("automated_backup_policy"); !isEmptyValue(reflect.ValueOf(automatedBackupPolicyProp)) && (ok || !reflect.DeepEqual(v, automatedBackupPolicyProp)) {
		obj["automatedBackupPolicy"] = automatedBackupPolicyProp
	}

	url, err := replaceVars(d, config, "{{AlloydbBasePath}}projects/{{project}}/locations/{{location}}/clusters?clusterId={{cluster_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Cluster: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Cluster: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Cluster: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/locations/{{location}}/clusters/{{cluster_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = alloydbOperationWaitTime(
		config, res, project, "Creating Cluster", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create Cluster: %s", err)
	}

	log.Printf("[DEBUG] Finished creating Cluster %q: %#v", d.Id(), res)

	return resourceAlloydbClusterRead(d, meta)
}

func resourceAlloydbClusterRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{AlloydbBasePath}}projects/{{project}}/locations/{{location}}/clusters/{{cluster_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Cluster: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("AlloydbCluster %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Cluster: %s", err)
	}

	if err := d.Set("name", flattenAlloydbClusterName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Cluster: %s", err)
	}
	if err := d.Set("uid", flattenAlloydbClusterUid(res["uid"], d, config)); err != nil {
		return fmt.Errorf("Error reading Cluster: %s", err)
	}
	if err := d.Set("labels", flattenAlloydbClusterLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Cluster: %s", err)
	}
	if err := d.Set("network", flattenAlloydbClusterNetwork(res["network"], d, config)); err != nil {
		return fmt.Errorf("Error reading Cluster: %s", err)
	}
	if err := d.Set("display_name", flattenAlloydbClusterDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading Cluster: %s", err)
	}
	if err := d.Set("database_version", flattenAlloydbClusterDatabaseVersion(res["databaseVersion"], d, config)); err != nil {
		return fmt.Errorf("Error reading Cluster: %s", err)
	}
	if err := d.Set("automated_backup_policy", flattenAlloydbClusterAutomatedBackupPolicy(res["automatedBackupPolicy"], d, config)); err != nil {
		return fmt.Errorf("Error reading Cluster: %s", err)
	}
	if err := d.Set("backup_source", flattenAlloydbClusterBackupSource(res["backupSource"], d, config)); err != nil {
		return fmt.Errorf("Error reading Cluster: %s", err)
	}
	if err := d.Set("migration_source", flattenAlloydbClusterMigrationSource(res["migrationSource"], d, config)); err != nil {
		return fmt.Errorf("Error reading Cluster: %s", err)
	}

	return nil
}

func resourceAlloydbClusterUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Cluster: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	labelsProp, err := expandAlloydbClusterLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	networkProp, err := expandAlloydbClusterNetwork(d.Get("network"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("network"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	displayNameProp, err := expandAlloydbClusterDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	automatedBackupPolicyProp, err := expandAlloydbClusterAutomatedBackupPolicy(d.Get("automated_backup_policy"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("automated_backup_policy"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, automatedBackupPolicyProp)) {
		obj["automatedBackupPolicy"] = automatedBackupPolicyProp
	}

	url, err := replaceVars(d, config, "{{AlloydbBasePath}}projects/{{project}}/locations/{{location}}/clusters/{{cluster_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Cluster %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("labels") {
		updateMask = append(updateMask, "labels")
	}

	if d.HasChange("network") {
		updateMask = append(updateMask, "network")
	}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}

	if d.HasChange("automated_backup_policy") {
		updateMask = append(updateMask, "automatedBackupPolicy")
	}
	// updateMask is a URL parameter but not present in the schema, so replaceVars
	// won't set it
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating Cluster %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Cluster %q: %#v", d.Id(), res)
	}

	err = alloydbOperationWaitTime(
		config, res, project, "Updating Cluster", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceAlloydbClusterRead(d, meta)
}

func resourceAlloydbClusterDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Cluster: %s", err)
	}
	billingProject = project

	url, err := replaceVars(d, config, "{{AlloydbBasePath}}projects/{{project}}/locations/{{location}}/clusters/{{cluster_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Cluster %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Cluster")
	}

	err = alloydbOperationWaitTime(
		config, res, project, "Deleting Cluster", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Cluster %q: %#v", d.Id(), res)
	return nil
}

func resourceAlloydbClusterImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/clusters/(?P<cluster_id>[^/]+)",
		"(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<cluster_id>[^/]+)",
		"(?P<location>[^/]+)/(?P<cluster_id>[^/]+)",
		"(?P<cluster_id>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/locations/{{location}}/clusters/{{cluster_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenAlloydbClusterName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAlloydbClusterUid(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAlloydbClusterLabels(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAlloydbClusterNetwork(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAlloydbClusterDisplayName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAlloydbClusterDatabaseVersion(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAlloydbClusterAutomatedBackupPolicy(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["backup_window"] =
		flattenAlloydbClusterAutomatedBackupPolicyBackupWindow(original["backupWindow"], d, config)
	transformed["location"] =
		flattenAlloydbClusterAutomatedBackupPolicyLocation(original["location"], d, config)
	transformed["labels"] =
		flattenAlloydbClusterAutomatedBackupPolicyLabels(original["labels"], d, config)
	transformed["weekly_schedule"] =
		flattenAlloydbClusterAutomatedBackupPolicyWeeklySchedule(original["weeklySchedule"], d, config)
	transformed["time_based_retention"] =
		flattenAlloydbClusterAutomatedBackupPolicyTimeBasedRetention(original["timeBasedRetention"], d, config)
	transformed["quantity_based_retention"] =
		flattenAlloydbClusterAutomatedBackupPolicyQuantityBasedRetention(original["quantityBasedRetention"], d, config)
	transformed["enabled"] =
		flattenAlloydbClusterAutomatedBackupPolicyEnabled(original["enabled"], d, config)
	return []interface{}{transformed}
}
func flattenAlloydbClusterAutomatedBackupPolicyBackupWindow(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAlloydbClusterAutomatedBackupPolicyLocation(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAlloydbClusterAutomatedBackupPolicyLabels(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAlloydbClusterAutomatedBackupPolicyWeeklySchedule(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["days_of_week"] =
		flattenAlloydbClusterAutomatedBackupPolicyWeeklyScheduleDaysOfWeek(original["daysOfWeek"], d, config)
	transformed["start_times"] =
		flattenAlloydbClusterAutomatedBackupPolicyWeeklyScheduleStartTimes(original["startTimes"], d, config)
	return []interface{}{transformed}
}
func flattenAlloydbClusterAutomatedBackupPolicyWeeklyScheduleDaysOfWeek(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAlloydbClusterAutomatedBackupPolicyWeeklyScheduleStartTimes(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"hours":   flattenAlloydbClusterAutomatedBackupPolicyWeeklyScheduleStartTimesHours(original["hours"], d, config),
			"minutes": flattenAlloydbClusterAutomatedBackupPolicyWeeklyScheduleStartTimesMinutes(original["minutes"], d, config),
			"seconds": flattenAlloydbClusterAutomatedBackupPolicyWeeklyScheduleStartTimesSeconds(original["seconds"], d, config),
			"nanos":   flattenAlloydbClusterAutomatedBackupPolicyWeeklyScheduleStartTimesNanos(original["nanos"], d, config),
		})
	}
	return transformed
}
func flattenAlloydbClusterAutomatedBackupPolicyWeeklyScheduleStartTimesHours(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := stringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenAlloydbClusterAutomatedBackupPolicyWeeklyScheduleStartTimesMinutes(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := stringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenAlloydbClusterAutomatedBackupPolicyWeeklyScheduleStartTimesSeconds(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := stringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenAlloydbClusterAutomatedBackupPolicyWeeklyScheduleStartTimesNanos(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := stringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenAlloydbClusterAutomatedBackupPolicyTimeBasedRetention(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["retention_period"] =
		flattenAlloydbClusterAutomatedBackupPolicyTimeBasedRetentionRetentionPeriod(original["retentionPeriod"], d, config)
	return []interface{}{transformed}
}
func flattenAlloydbClusterAutomatedBackupPolicyTimeBasedRetentionRetentionPeriod(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAlloydbClusterAutomatedBackupPolicyQuantityBasedRetention(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["count"] =
		flattenAlloydbClusterAutomatedBackupPolicyQuantityBasedRetentionCount(original["count"], d, config)
	return []interface{}{transformed}
}
func flattenAlloydbClusterAutomatedBackupPolicyQuantityBasedRetentionCount(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := stringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenAlloydbClusterAutomatedBackupPolicyEnabled(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAlloydbClusterBackupSource(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["backup_name"] =
		flattenAlloydbClusterBackupSourceBackupName(original["backupName"], d, config)
	return []interface{}{transformed}
}
func flattenAlloydbClusterBackupSourceBackupName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAlloydbClusterMigrationSource(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["host_port"] =
		flattenAlloydbClusterMigrationSourceHostPort(original["hostPort"], d, config)
	transformed["reference_id"] =
		flattenAlloydbClusterMigrationSourceReferenceId(original["referenceId"], d, config)
	transformed["source_type"] =
		flattenAlloydbClusterMigrationSourceSourceType(original["sourceType"], d, config)
	return []interface{}{transformed}
}
func flattenAlloydbClusterMigrationSourceHostPort(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAlloydbClusterMigrationSourceReferenceId(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAlloydbClusterMigrationSourceSourceType(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandAlloydbClusterLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandAlloydbClusterNetwork(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbClusterDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbClusterInitialUser(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedUser, err := expandAlloydbClusterInitialUserUser(original["user"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUser); val.IsValid() && !isEmptyValue(val) {
		transformed["user"] = transformedUser
	}

	transformedPassword, err := expandAlloydbClusterInitialUserPassword(original["password"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPassword); val.IsValid() && !isEmptyValue(val) {
		transformed["password"] = transformedPassword
	}

	return transformed, nil
}

func expandAlloydbClusterInitialUserUser(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbClusterInitialUserPassword(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbClusterAutomatedBackupPolicy(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedBackupWindow, err := expandAlloydbClusterAutomatedBackupPolicyBackupWindow(original["backup_window"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBackupWindow); val.IsValid() && !isEmptyValue(val) {
		transformed["backupWindow"] = transformedBackupWindow
	}

	transformedLocation, err := expandAlloydbClusterAutomatedBackupPolicyLocation(original["location"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLocation); val.IsValid() && !isEmptyValue(val) {
		transformed["location"] = transformedLocation
	}

	transformedLabels, err := expandAlloydbClusterAutomatedBackupPolicyLabels(original["labels"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLabels); val.IsValid() && !isEmptyValue(val) {
		transformed["labels"] = transformedLabels
	}

	transformedWeeklySchedule, err := expandAlloydbClusterAutomatedBackupPolicyWeeklySchedule(original["weekly_schedule"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedWeeklySchedule); val.IsValid() && !isEmptyValue(val) {
		transformed["weeklySchedule"] = transformedWeeklySchedule
	}

	transformedTimeBasedRetention, err := expandAlloydbClusterAutomatedBackupPolicyTimeBasedRetention(original["time_based_retention"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTimeBasedRetention); val.IsValid() && !isEmptyValue(val) {
		transformed["timeBasedRetention"] = transformedTimeBasedRetention
	}

	transformedQuantityBasedRetention, err := expandAlloydbClusterAutomatedBackupPolicyQuantityBasedRetention(original["quantity_based_retention"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedQuantityBasedRetention); val.IsValid() && !isEmptyValue(val) {
		transformed["quantityBasedRetention"] = transformedQuantityBasedRetention
	}

	transformedEnabled, err := expandAlloydbClusterAutomatedBackupPolicyEnabled(original["enabled"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnabled); val.IsValid() && !isEmptyValue(val) {
		transformed["enabled"] = transformedEnabled
	}

	return transformed, nil
}

func expandAlloydbClusterAutomatedBackupPolicyBackupWindow(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbClusterAutomatedBackupPolicyLocation(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbClusterAutomatedBackupPolicyLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandAlloydbClusterAutomatedBackupPolicyWeeklySchedule(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDaysOfWeek, err := expandAlloydbClusterAutomatedBackupPolicyWeeklyScheduleDaysOfWeek(original["days_of_week"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDaysOfWeek); val.IsValid() && !isEmptyValue(val) {
		transformed["daysOfWeek"] = transformedDaysOfWeek
	}

	transformedStartTimes, err := expandAlloydbClusterAutomatedBackupPolicyWeeklyScheduleStartTimes(original["start_times"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStartTimes); val.IsValid() && !isEmptyValue(val) {
		transformed["startTimes"] = transformedStartTimes
	}

	return transformed, nil
}

func expandAlloydbClusterAutomatedBackupPolicyWeeklyScheduleDaysOfWeek(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbClusterAutomatedBackupPolicyWeeklyScheduleStartTimes(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedHours, err := expandAlloydbClusterAutomatedBackupPolicyWeeklyScheduleStartTimesHours(original["hours"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedHours); val.IsValid() && !isEmptyValue(val) {
			transformed["hours"] = transformedHours
		}

		transformedMinutes, err := expandAlloydbClusterAutomatedBackupPolicyWeeklyScheduleStartTimesMinutes(original["minutes"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedMinutes); val.IsValid() && !isEmptyValue(val) {
			transformed["minutes"] = transformedMinutes
		}

		transformedSeconds, err := expandAlloydbClusterAutomatedBackupPolicyWeeklyScheduleStartTimesSeconds(original["seconds"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedSeconds); val.IsValid() && !isEmptyValue(val) {
			transformed["seconds"] = transformedSeconds
		}

		transformedNanos, err := expandAlloydbClusterAutomatedBackupPolicyWeeklyScheduleStartTimesNanos(original["nanos"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedNanos); val.IsValid() && !isEmptyValue(val) {
			transformed["nanos"] = transformedNanos
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandAlloydbClusterAutomatedBackupPolicyWeeklyScheduleStartTimesHours(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbClusterAutomatedBackupPolicyWeeklyScheduleStartTimesMinutes(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbClusterAutomatedBackupPolicyWeeklyScheduleStartTimesSeconds(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbClusterAutomatedBackupPolicyWeeklyScheduleStartTimesNanos(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbClusterAutomatedBackupPolicyTimeBasedRetention(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRetentionPeriod, err := expandAlloydbClusterAutomatedBackupPolicyTimeBasedRetentionRetentionPeriod(original["retention_period"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRetentionPeriod); val.IsValid() && !isEmptyValue(val) {
		transformed["retentionPeriod"] = transformedRetentionPeriod
	}

	return transformed, nil
}

func expandAlloydbClusterAutomatedBackupPolicyTimeBasedRetentionRetentionPeriod(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbClusterAutomatedBackupPolicyQuantityBasedRetention(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCount, err := expandAlloydbClusterAutomatedBackupPolicyQuantityBasedRetentionCount(original["count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCount); val.IsValid() && !isEmptyValue(val) {
		transformed["count"] = transformedCount
	}

	return transformed, nil
}

func expandAlloydbClusterAutomatedBackupPolicyQuantityBasedRetentionCount(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbClusterAutomatedBackupPolicyEnabled(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
