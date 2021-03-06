// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: CA certificate.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceCertificateCa() *schema.Resource {
	return &schema.Resource{
		Create: resourceCertificateCaCreate,
		Read:   resourceCertificateCaRead,
		Update: resourceCertificateCaUpdate,
		Delete: resourceCertificateCaDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				ForceNew:     true,
				Required:     true,
			},
			"ca": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"range": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trusted": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scep_url": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"auto_update_days": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"auto_update_days_warning": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"last_updated": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceCertificateCaCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectCertificateCa(d)
	if err != nil {
		return fmt.Errorf("Error creating CertificateCa resource while getting object: %v", err)
	}

	o, err := c.CreateCertificateCa(obj)

	if err != nil {
		return fmt.Errorf("Error creating CertificateCa resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("CertificateCa")
	}

	return resourceCertificateCaRead(d, m)
}

func resourceCertificateCaUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectCertificateCa(d)
	if err != nil {
		return fmt.Errorf("Error updating CertificateCa resource while getting object: %v", err)
	}

	o, err := c.UpdateCertificateCa(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating CertificateCa resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("CertificateCa")
	}

	return resourceCertificateCaRead(d, m)
}

func resourceCertificateCaDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteCertificateCa(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting CertificateCa resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceCertificateCaRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadCertificateCa(mkey)
	if err != nil {
		return fmt.Errorf("Error reading CertificateCa resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectCertificateCa(d, o)
	if err != nil {
		return fmt.Errorf("Error reading CertificateCa resource from API: %v", err)
	}
	return nil
}

func flattenCertificateCaName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCertificateCaCa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCertificateCaRange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCertificateCaSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCertificateCaTrusted(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCertificateCaScepUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCertificateCaAutoUpdateDays(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCertificateCaAutoUpdateDaysWarning(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCertificateCaSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenCertificateCaLastUpdated(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectCertificateCa(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenCertificateCaName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("ca", flattenCertificateCaCa(o["ca"], d, "ca")); err != nil {
		if !fortiAPIPatch(o["ca"]) {
			return fmt.Errorf("Error reading ca: %v", err)
		}
	}

	if err = d.Set("range", flattenCertificateCaRange(o["range"], d, "range")); err != nil {
		if !fortiAPIPatch(o["range"]) {
			return fmt.Errorf("Error reading range: %v", err)
		}
	}

	if err = d.Set("source", flattenCertificateCaSource(o["source"], d, "source")); err != nil {
		if !fortiAPIPatch(o["source"]) {
			return fmt.Errorf("Error reading source: %v", err)
		}
	}

	if err = d.Set("trusted", flattenCertificateCaTrusted(o["trusted"], d, "trusted")); err != nil {
		if !fortiAPIPatch(o["trusted"]) {
			return fmt.Errorf("Error reading trusted: %v", err)
		}
	}

	if err = d.Set("scep_url", flattenCertificateCaScepUrl(o["scep-url"], d, "scep_url")); err != nil {
		if !fortiAPIPatch(o["scep-url"]) {
			return fmt.Errorf("Error reading scep_url: %v", err)
		}
	}

	if err = d.Set("auto_update_days", flattenCertificateCaAutoUpdateDays(o["auto-update-days"], d, "auto_update_days")); err != nil {
		if !fortiAPIPatch(o["auto-update-days"]) {
			return fmt.Errorf("Error reading auto_update_days: %v", err)
		}
	}

	if err = d.Set("auto_update_days_warning", flattenCertificateCaAutoUpdateDaysWarning(o["auto-update-days-warning"], d, "auto_update_days_warning")); err != nil {
		if !fortiAPIPatch(o["auto-update-days-warning"]) {
			return fmt.Errorf("Error reading auto_update_days_warning: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenCertificateCaSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("last_updated", flattenCertificateCaLastUpdated(o["last-updated"], d, "last_updated")); err != nil {
		if !fortiAPIPatch(o["last-updated"]) {
			return fmt.Errorf("Error reading last_updated: %v", err)
		}
	}

	return nil
}

func flattenCertificateCaFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandCertificateCaName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCertificateCaCa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCertificateCaRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCertificateCaSource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCertificateCaTrusted(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCertificateCaScepUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCertificateCaAutoUpdateDays(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCertificateCaAutoUpdateDaysWarning(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCertificateCaSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandCertificateCaLastUpdated(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectCertificateCa(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandCertificateCaName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("ca"); ok {
		t, err := expandCertificateCaCa(d, v, "ca")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ca"] = t
		}
	}

	if v, ok := d.GetOk("range"); ok {
		t, err := expandCertificateCaRange(d, v, "range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["range"] = t
		}
	}

	if v, ok := d.GetOk("source"); ok {
		t, err := expandCertificateCaSource(d, v, "source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source"] = t
		}
	}

	if v, ok := d.GetOk("trusted"); ok {
		t, err := expandCertificateCaTrusted(d, v, "trusted")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusted"] = t
		}
	}

	if v, ok := d.GetOk("scep_url"); ok {
		t, err := expandCertificateCaScepUrl(d, v, "scep_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scep-url"] = t
		}
	}

	if v, ok := d.GetOkExists("auto_update_days"); ok {
		t, err := expandCertificateCaAutoUpdateDays(d, v, "auto_update_days")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-update-days"] = t
		}
	}

	if v, ok := d.GetOkExists("auto_update_days_warning"); ok {
		t, err := expandCertificateCaAutoUpdateDaysWarning(d, v, "auto_update_days_warning")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-update-days-warning"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok {
		t, err := expandCertificateCaSourceIp(d, v, "source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOkExists("last_updated"); ok {
		t, err := expandCertificateCaLastUpdated(d, v, "last_updated")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["last-updated"] = t
		}
	}

	return &obj, nil
}
