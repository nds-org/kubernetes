package volumes

import (
	"strings"
	"github.com/rackspace/gophercloud"
)

func createURL(c *gophercloud.ServiceClient) string {
	url := c.ServiceURL("volumes")
	url = strings.Replace(url, "/v2/", "/v1/", 1)
	return url
}

func listURL(c *gophercloud.ServiceClient) string {
	return createURL(c)
}

func deleteURL(c *gophercloud.ServiceClient, id string) string {
	url := c.ServiceURL("volumes", id)
	url = strings.Replace(url, "/v2/", "/v1/", 1)
	return url
}

func getURL(c *gophercloud.ServiceClient, id string) string {
	return deleteURL(c, id)
}

func updateURL(c *gophercloud.ServiceClient, id string) string {
	return deleteURL(c, id)
}
