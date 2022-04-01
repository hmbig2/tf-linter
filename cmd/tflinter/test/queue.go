package test

import (
	"regexp"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var regexp4Name = regexp.MustCompile(`^[a-z0-9_]{1,128}$`)

const CU_16, CU_64, CU_256 = 16, 64, 256
const RESOURCE_MODE_SHARED, RESOURCE_MODE_EXCLUSIVE = 0, 1
const QUEUE_TYPE_SQL, QUEUE_TYPE_GENERAL = "sql", "general"
const QUEUE_FEATURE_BASIC, QUEUE_FEATURE_AI = "basic", "ai"
const QUEUE_PLATFORM_X86, QUEUE_platform_AARCH64 = "x86_64", "aarch64"

const (
	actionRestart  = "restart"
	actionScaleOut = "scale_out"
	actionScaleIn  = "scale_in"
)

func ResourceDliQueue() *schema.Resource {
	return &schema.Resource{
		Create: resourceDliQueueCreate,
		Read:   resourceDliQueueRead,
		Update: resourceDliQueueUpdate,
		Delete: resourceDliQueueDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"regionA": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringMatch(regexp4Name, " only contain digits, lower letters, and underscores (_)"),
			},

			"queue_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				Default:      "sql",
				ValidateFunc: validation.StringInSlice([]string{QUEUE_TYPE_SQL, QUEUE_TYPE_GENERAL}, false),
			},

			"description": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				ForceNew: true,
			},

			"cu_count": {
				Type:     schema.TypeInt,
				Required: true,
			},

			"enterprise_project_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},

			"platform": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				Default:      QUEUE_PLATFORM_X86,
				ValidateFunc: validation.StringInSlice([]string{QUEUE_PLATFORM_X86, QUEUE_platform_AARCH64}, false),
			},

			"resource_mode": {
				Type:         schema.TypeInt,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.IntInSlice([]int{RESOURCE_MODE_SHARED, RESOURCE_MODE_EXCLUSIVE}),
			},

			"feature": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{QUEUE_FEATURE_BASIC, QUEUE_FEATURE_AI}, false),
			},

			"tags": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				ForceNew: true,
			},

			"vpc_cidr": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"create_time": {
				Type:     schema.TypeInt,
				Computed: true,
			},

			"management_subnet_cidr": {
				Type:       schema.TypeString,
				Optional:   true,
				ForceNew:   true,
				Deprecated: "management_subnet_cidr is Deprecated",
			},

			"subnet_cidr": {
				Type:       schema.TypeString,
				Optional:   true,
				ForceNew:   true,
				Deprecated: "subnet_cidr is Deprecated",
			},

			"conf": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"spark_sql_max_records_per_file": {
							Type:     schema.TypeInt,
							Optional: true,
							ForceNew: true,
						},
						"spark_sql_auto_broadcast_join_threshold": {
							Type:     schema.TypeInt,
							Optional: true,
							ForceNew: true,
						},
						"spark_sql_shuffle_partitions": {
							Type:     schema.TypeInt,
							Optional: true,
							ForceNew: true,
						},
						"spark_sql_dynamic_partition_overwrite_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							ForceNew: true,
						},
						"spark_sql_files_max_partition_bytes": {
							Type:     schema.TypeInt,
							Optional: true,
							ForceNew: true,
						},
						"spark_sql_bad_records_path": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"dli_sql_sqlasync_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							ForceNew: true,
						},
						"dli_sql_job_timeout": {
							Type:     schema.TypeInt,
							Optional: true,
							ForceNew: true,
						},
					},
				},
			},
		},

		Timeouts: &schema.ResourceTimeout{
			Update: schema.DefaultTimeout(45 * time.Minute),
		},
	}
}

func resourceDliQueueCreate(d *schema.ResourceData, meta interface{}) error {

	return resourceDliQueueRead(d, meta)
}

func resourceDliQueueRead(d *schema.ResourceData, meta interface{}) error {

	return nil
}

func resourceDliQueueDelete(d *schema.ResourceData, meta interface{}) error {

	return nil
}

/*
  support cu_count scaling
*/
func resourceDliQueueUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceDliQueueRead(d, meta)
}
