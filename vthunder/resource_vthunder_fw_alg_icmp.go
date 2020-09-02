package vthunder

//vThunder resource FwAlgIcmp

import (
	"github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
	"util"
)

func resourceFwAlgIcmp() *schema.Resource {
	return &schema.Resource{
		Create: resourceFwAlgIcmpCreate,
		Update: resourceFwAlgIcmpUpdate,
		Read:   resourceFwAlgIcmpRead,
		Delete: resourceFwAlgIcmpDelete,
		Schema: map[string]*schema.Schema{
			"disable": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceFwAlgIcmpCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating FwAlgIcmp (Inside resourceFwAlgIcmpCreate) ")

		data := dataToFwAlgIcmp(d)
		logger.Println("[INFO] received formatted data from method data to FwAlgIcmp --")
		d.SetId(strconv.Itoa('1'))
		go_vthunder.PostFwAlgIcmp(client.Token, data, client.Host)

		return resourceFwAlgIcmpRead(d, meta)

	}
	return nil
}

func resourceFwAlgIcmpRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading FwAlgIcmp (Inside resourceFwAlgIcmpRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_vthunder.GetFwAlgIcmp(client.Token, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceFwAlgIcmpUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceFwAlgIcmpRead(d, meta)
}

func resourceFwAlgIcmpDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceFwAlgIcmpRead(d, meta)
}
func dataToFwAlgIcmp(d *schema.ResourceData) go_vthunder.FwAlgIcmp {
	var vc go_vthunder.FwAlgIcmp
	var c go_vthunder.FwAlgIcmpInstance
	c.Disable = d.Get("disable").(string)

	vc.Disable = c
	return vc
}
