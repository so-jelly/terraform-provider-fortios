// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure how log messages are displayed on the GUI.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceLogGuiDisplay() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogGuiDisplayUpdate,
		Read:   resourceLogGuiDisplayRead,
		Update: resourceLogGuiDisplayUpdate,
		Delete: resourceLogGuiDisplayDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"resolve_hosts": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resolve_apps": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortiview_unscanned_apps": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceLogGuiDisplayUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectLogGuiDisplay(d)
	if err != nil {
		return fmt.Errorf("Error updating LogGuiDisplay resource while getting object: %v", err)
	}

	o, err := c.UpdateLogGuiDisplay(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating LogGuiDisplay resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("LogGuiDisplay")
	}

	return resourceLogGuiDisplayRead(d, m)
}

func resourceLogGuiDisplayDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteLogGuiDisplay(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting LogGuiDisplay resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogGuiDisplayRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadLogGuiDisplay(mkey)
	if err != nil {
		return fmt.Errorf("Error reading LogGuiDisplay resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogGuiDisplay(d, o)
	if err != nil {
		return fmt.Errorf("Error reading LogGuiDisplay resource from API: %v", err)
	}
	return nil
}

func flattenLogGuiDisplayResolveHosts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogGuiDisplayResolveApps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogGuiDisplayFortiviewUnscannedApps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectLogGuiDisplay(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("resolve_hosts", flattenLogGuiDisplayResolveHosts(o["resolve-hosts"], d, "resolve_hosts")); err != nil {
		if !fortiAPIPatch(o["resolve-hosts"]) {
			return fmt.Errorf("Error reading resolve_hosts: %v", err)
		}
	}

	if err = d.Set("resolve_apps", flattenLogGuiDisplayResolveApps(o["resolve-apps"], d, "resolve_apps")); err != nil {
		if !fortiAPIPatch(o["resolve-apps"]) {
			return fmt.Errorf("Error reading resolve_apps: %v", err)
		}
	}

	if err = d.Set("fortiview_unscanned_apps", flattenLogGuiDisplayFortiviewUnscannedApps(o["fortiview-unscanned-apps"], d, "fortiview_unscanned_apps")); err != nil {
		if !fortiAPIPatch(o["fortiview-unscanned-apps"]) {
			return fmt.Errorf("Error reading fortiview_unscanned_apps: %v", err)
		}
	}

	return nil
}

func flattenLogGuiDisplayFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandLogGuiDisplayResolveHosts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogGuiDisplayResolveApps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogGuiDisplayFortiviewUnscannedApps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectLogGuiDisplay(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("resolve_hosts"); ok {
		t, err := expandLogGuiDisplayResolveHosts(d, v, "resolve_hosts")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["resolve-hosts"] = t
		}
	}

	if v, ok := d.GetOk("resolve_apps"); ok {
		t, err := expandLogGuiDisplayResolveApps(d, v, "resolve_apps")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["resolve-apps"] = t
		}
	}

	if v, ok := d.GetOk("fortiview_unscanned_apps"); ok {
		t, err := expandLogGuiDisplayFortiviewUnscannedApps(d, v, "fortiview_unscanned_apps")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiview-unscanned-apps"] = t
		}
	}

	return &obj, nil
}
