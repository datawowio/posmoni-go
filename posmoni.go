/*
Package posmoni provides GO binding for Posmoni REST API.
Full REST API documentation is available at https://datawow.readme.io/v1.0/reference.

Usage

Create a client with posmoni.NewClient, along with supply your project key. After that, use
client.Call with actions object from the https://godoc.org/github.com/datawowio/posmoni/actions
package to perform API calls. The first parameter to client.Call lets you supply a struct
object from index that listed below to unmarshal the result.

Example

	c, err := posmoni.NewClient(ProjectKey)
	if err != nil {
		log.Fatal(err)
	}

	Moderation, getModeration := &posmoni.GetModerations{}, &actions.GetModerations{
		ID: "5a52fb556e11571f570c1530",
	}

	if err := c.Call(Moderation, getModeration); err != nil {
		log.Fatal(err)
	}
	log.Printf("%#v\n", Moderation)

We also provide Get "any type" Image endpoint API. You only supply project key and Image
ID (or Customer ID) for reference.

Example

	c, err := posmoni.NewClient(ProjectKey)
	if err != nil {
		log.Fatal(err)
	}

	var resp map[string]interface{}

	getModerations := &actions.GetModerations{
		ID: "61651d8f3a96703f9768a124",
	}

	if err := c.Call(&resp, getModerations); err != nil {
		log.Fatal(err)
	}

	data := resp["data"].(map[string]interface{})
	attrs := data["attributes"].(map[string]interface{})
	log.Println(attrs["answer"])

*/
package posmoni
