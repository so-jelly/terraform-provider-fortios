// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure Link Health Monitor.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemLinkMonitor() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemLinkMonitorCreate,
		Read:   resourceSystemLinkMonitorRead,
		Update: resourceSystemLinkMonitorUpdate,
		Delete: resourceSystemLinkMonitorDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional: true,
				Computed: true,
			},
			"addr_mode": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcintf": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional: true,
				Computed: true,
			},
			"server": &schema.Schema{
				Type: schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": &schema.Schema{
							Type: schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"protocol": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port": &schema.Schema{
				Type: schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional: true,
				Computed: true,
			},
			"gateway_ip": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gateway_ip6": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip6": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http_get": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1024),
				Required: true,
			},
			"http_agent": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1024),
				Optional: true,
				Computed: true,
			},
			"http_match": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1024),
				Optional: true,
				Computed: true,
			},
			"interval": &schema.Schema{
				Type: schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 3600),
				Optional: true,
				Computed: true,
			},
			"failtime": &schema.Schema{
				Type: schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 3600),
				Optional: true,
				Computed: true,
			},
			"recoverytime": &schema.Schema{
				Type: schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 3600),
				Optional: true,
				Computed: true,
			},
			"security_mode": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"password": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional: true,
			},
			"packet_size": &schema.Schema{
				Type: schema.TypeInt,
				ValidateFunc: validation.IntBetween(64, 1024),
				Optional: true,
				Computed: true,
			},
			"ha_priority": &schema.Schema{
				Type: schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 50),
				Optional: true,
				Computed: true,
			},
			"update_cascade_interface": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"update_static_route": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemLinkMonitorCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemLinkMonitor(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemLinkMonitor resource while getting object: %v", err)
	}

	o, err := c.CreateSystemLinkMonitor(obj)

	if err != nil {
		return fmt.Errorf("Error creating SystemLinkMonitor resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemLinkMonitor")
	}

	return resourceSystemLinkMonitorRead(d, m)
}

func resourceSystemLinkMonitorUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemLinkMonitor(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemLinkMonitor resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemLinkMonitor(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemLinkMonitor resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemLinkMonitor")
	}

	return resourceSystemLinkMonitorRead(d, m)
}

func resourceSystemLinkMonitorDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemLinkMonitor(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemLinkMonitor resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemLinkMonitorRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemLinkMonitor(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemLinkMonitor resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemLinkMonitor(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemLinkMonitor resource from API: %v", err)
	}
	return nil
}


func flattenSystemLinkMonitorName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLinkMonitorAddrMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLinkMonitorSrcintf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLinkMonitorServer(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := i["address"]; ok {
			tmp["address"] = flattenSystemLinkMonitorServerAddress(i["address"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemLinkMonitorServerAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLinkMonitorProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLinkMonitorPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLinkMonitorGatewayIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLinkMonitorGatewayIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLinkMonitorSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLinkMonitorSourceIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLinkMonitorHttpGet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLinkMonitorHttpAgent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLinkMonitorHttpMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLinkMonitorInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLinkMonitorFailtime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLinkMonitorRecoverytime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLinkMonitorSecurityMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLinkMonitorPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLinkMonitorPacketSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLinkMonitorHaPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLinkMonitorUpdateCascadeInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLinkMonitorUpdateStaticRoute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLinkMonitorStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}


func refreshObjectSystemLinkMonitor(d *schema.ResourceData, o map[string]interface{}) error {
	var err error


	if err = d.Set("name", flattenSystemLinkMonitorName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("addr_mode", flattenSystemLinkMonitorAddrMode(o["addr-mode"], d, "addr_mode")); err != nil {
		if !fortiAPIPatch(o["addr-mode"]) {
			return fmt.Errorf("Error reading addr_mode: %v", err)
		}
	}

	if err = d.Set("srcintf", flattenSystemLinkMonitorSrcintf(o["srcintf"], d, "srcintf")); err != nil {
		if !fortiAPIPatch(o["srcintf"]) {
			return fmt.Errorf("Error reading srcintf: %v", err)
		}
	}

    if isImportTable() {
        if err = d.Set("server", flattenSystemLinkMonitorServer(o["server"], d, "server")); err != nil {
            if !fortiAPIPatch(o["server"]) {
                return fmt.Errorf("Error reading server: %v", err)
            }
        }
    } else {
        if _, ok := d.GetOk("server"); ok {
            if err = d.Set("server", flattenSystemLinkMonitorServer(o["server"], d, "server")); err != nil {
                if !fortiAPIPatch(o["server"]) {
                    return fmt.Errorf("Error reading server: %v", err)
                }
            }
        }
    }

	if err = d.Set("protocol", flattenSystemLinkMonitorProtocol(o["protocol"], d, "protocol")); err != nil {
		if !fortiAPIPatch(o["protocol"]) {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("port", flattenSystemLinkMonitorPort(o["port"], d, "port")); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("gateway_ip", flattenSystemLinkMonitorGatewayIp(o["gateway-ip"], d, "gateway_ip")); err != nil {
		if !fortiAPIPatch(o["gateway-ip"]) {
			return fmt.Errorf("Error reading gateway_ip: %v", err)
		}
	}

	if err = d.Set("gateway_ip6", flattenSystemLinkMonitorGatewayIp6(o["gateway-ip6"], d, "gateway_ip6")); err != nil {
		if !fortiAPIPatch(o["gateway-ip6"]) {
			return fmt.Errorf("Error reading gateway_ip6: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenSystemLinkMonitorSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("source_ip6", flattenSystemLinkMonitorSourceIp6(o["source-ip6"], d, "source_ip6")); err != nil {
		if !fortiAPIPatch(o["source-ip6"]) {
			return fmt.Errorf("Error reading source_ip6: %v", err)
		}
	}

	if err = d.Set("http_get", flattenSystemLinkMonitorHttpGet(o["http-get"], d, "http_get")); err != nil {
		if !fortiAPIPatch(o["http-get"]) {
			return fmt.Errorf("Error reading http_get: %v", err)
		}
	}

	if err = d.Set("http_agent", flattenSystemLinkMonitorHttpAgent(o["http-agent"], d, "http_agent")); err != nil {
		if !fortiAPIPatch(o["http-agent"]) {
			return fmt.Errorf("Error reading http_agent: %v", err)
		}
	}

	if err = d.Set("http_match", flattenSystemLinkMonitorHttpMatch(o["http-match"], d, "http_match")); err != nil {
		if !fortiAPIPatch(o["http-match"]) {
			return fmt.Errorf("Error reading http_match: %v", err)
		}
	}

	if err = d.Set("interval", flattenSystemLinkMonitorInterval(o["interval"], d, "interval")); err != nil {
		if !fortiAPIPatch(o["interval"]) {
			return fmt.Errorf("Error reading interval: %v", err)
		}
	}

	if err = d.Set("failtime", flattenSystemLinkMonitorFailtime(o["failtime"], d, "failtime")); err != nil {
		if !fortiAPIPatch(o["failtime"]) {
			return fmt.Errorf("Error reading failtime: %v", err)
		}
	}

	if err = d.Set("recoverytime", flattenSystemLinkMonitorRecoverytime(o["recoverytime"], d, "recoverytime")); err != nil {
		if !fortiAPIPatch(o["recoverytime"]) {
			return fmt.Errorf("Error reading recoverytime: %v", err)
		}
	}

	if err = d.Set("security_mode", flattenSystemLinkMonitorSecurityMode(o["security-mode"], d, "security_mode")); err != nil {
		if !fortiAPIPatch(o["security-mode"]) {
			return fmt.Errorf("Error reading security_mode: %v", err)
		}
	}

	if err = d.Set("password", flattenSystemLinkMonitorPassword(o["password"], d, "password")); err != nil {
		if !fortiAPIPatch(o["password"]) {
			return fmt.Errorf("Error reading password: %v", err)
		}
	}

	if err = d.Set("packet_size", flattenSystemLinkMonitorPacketSize(o["packet-size"], d, "packet_size")); err != nil {
		if !fortiAPIPatch(o["packet-size"]) {
			return fmt.Errorf("Error reading packet_size: %v", err)
		}
	}

	if err = d.Set("ha_priority", flattenSystemLinkMonitorHaPriority(o["ha-priority"], d, "ha_priority")); err != nil {
		if !fortiAPIPatch(o["ha-priority"]) {
			return fmt.Errorf("Error reading ha_priority: %v", err)
		}
	}

	if err = d.Set("update_cascade_interface", flattenSystemLinkMonitorUpdateCascadeInterface(o["update-cascade-interface"], d, "update_cascade_interface")); err != nil {
		if !fortiAPIPatch(o["update-cascade-interface"]) {
			return fmt.Errorf("Error reading update_cascade_interface: %v", err)
		}
	}

	if err = d.Set("update_static_route", flattenSystemLinkMonitorUpdateStaticRoute(o["update-static-route"], d, "update_static_route")); err != nil {
		if !fortiAPIPatch(o["update-static-route"]) {
			return fmt.Errorf("Error reading update_static_route: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemLinkMonitorStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}


	return nil
}

func flattenSystemLinkMonitorFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}


func expandSystemLinkMonitorName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLinkMonitorAddrMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLinkMonitorSrcintf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLinkMonitorServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := ""  // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["address"], _ = expandSystemLinkMonitorServerAddress(d, i["address"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemLinkMonitorServerAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLinkMonitorProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLinkMonitorPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLinkMonitorGatewayIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLinkMonitorGatewayIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLinkMonitorSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLinkMonitorSourceIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLinkMonitorHttpGet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLinkMonitorHttpAgent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLinkMonitorHttpMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLinkMonitorInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLinkMonitorFailtime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLinkMonitorRecoverytime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLinkMonitorSecurityMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLinkMonitorPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLinkMonitorPacketSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLinkMonitorHaPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLinkMonitorUpdateCascadeInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLinkMonitorUpdateStaticRoute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLinkMonitorStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}


func getObjectSystemLinkMonitor(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})


	if v, ok := d.GetOk("name"); ok {
		t, err := expandSystemLinkMonitorName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("addr_mode"); ok {
		t, err := expandSystemLinkMonitorAddrMode(d, v, "addr_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addr-mode"] = t
		}
	}

	if v, ok := d.GetOk("srcintf"); ok {
		t, err := expandSystemLinkMonitorSrcintf(d, v, "srcintf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcintf"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok {
		t, err := expandSystemLinkMonitorServer(d, v, "server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok {
		t, err := expandSystemLinkMonitorProtocol(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok {
		t, err := expandSystemLinkMonitorPort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("gateway_ip"); ok {
		t, err := expandSystemLinkMonitorGatewayIp(d, v, "gateway_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gateway-ip"] = t
		}
	}

	if v, ok := d.GetOk("gateway_ip6"); ok {
		t, err := expandSystemLinkMonitorGatewayIp6(d, v, "gateway_ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gateway-ip6"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok {
		t, err := expandSystemLinkMonitorSourceIp(d, v, "source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("source_ip6"); ok {
		t, err := expandSystemLinkMonitorSourceIp6(d, v, "source_ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip6"] = t
		}
	}

	if v, ok := d.GetOk("http_get"); ok {
		t, err := expandSystemLinkMonitorHttpGet(d, v, "http_get")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-get"] = t
		}
	}

	if v, ok := d.GetOk("http_agent"); ok {
		t, err := expandSystemLinkMonitorHttpAgent(d, v, "http_agent")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-agent"] = t
		}
	}

	if v, ok := d.GetOk("http_match"); ok {
		t, err := expandSystemLinkMonitorHttpMatch(d, v, "http_match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-match"] = t
		}
	}

	if v, ok := d.GetOk("interval"); ok {
		t, err := expandSystemLinkMonitorInterval(d, v, "interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interval"] = t
		}
	}

	if v, ok := d.GetOk("failtime"); ok {
		t, err := expandSystemLinkMonitorFailtime(d, v, "failtime")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["failtime"] = t
		}
	}

	if v, ok := d.GetOk("recoverytime"); ok {
		t, err := expandSystemLinkMonitorRecoverytime(d, v, "recoverytime")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["recoverytime"] = t
		}
	}

	if v, ok := d.GetOk("security_mode"); ok {
		t, err := expandSystemLinkMonitorSecurityMode(d, v, "security_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-mode"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok {
		t, err := expandSystemLinkMonitorPassword(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("packet_size"); ok {
		t, err := expandSystemLinkMonitorPacketSize(d, v, "packet_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["packet-size"] = t
		}
	}

	if v, ok := d.GetOk("ha_priority"); ok {
		t, err := expandSystemLinkMonitorHaPriority(d, v, "ha_priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha-priority"] = t
		}
	}

	if v, ok := d.GetOk("update_cascade_interface"); ok {
		t, err := expandSystemLinkMonitorUpdateCascadeInterface(d, v, "update_cascade_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-cascade-interface"] = t
		}
	}

	if v, ok := d.GetOk("update_static_route"); ok {
		t, err := expandSystemLinkMonitorUpdateStaticRoute(d, v, "update_static_route")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-static-route"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandSystemLinkMonitorStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}


	return &obj, nil
}
