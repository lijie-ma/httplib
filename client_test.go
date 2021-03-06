package gohttp

import (
	"testing"
)

var (
	proxy       = "http://127.0.0.1:8888"
	postUri     = "https://192.168.56.100/post.php"
	loginUri    = "http://192.168.56.100/login.php"
	uploadUri   = "http://192.168.56.100/upload.php"
	session1    = "http://192.168.56.100/session1.php"
	session2Uri = "http://192.168.56.100/session2.php"
	authUri     = `http://192.168.56.100/auth.php`

	jsonUri = "https://192.168.56.100/json.php"
)

func TestHttp2(t *testing.T) {
	v := map[string]interface{}{
		"form_params": map[string]interface{}{
			"key": "ivideo_index",
		},
		"http2": true,
	}
	c := NewClient(v)
	resp := c.Post(postUri, nil)
	if nil == resp {
		t.Log(c.GetErrors())
		return
	}
	t.Log(resp.Body)
}

func TestClient_Post(t *testing.T) {
	v := map[string]interface{}{
		"form_params": map[string]interface{}{
			"key": "ivideo_index",
		},
		//"proxy": proxy,
	}
	c := NewClient(v)
	resp := c.Post(postUri, nil)
	if nil == resp {
		t.Log(c.GetErrors())
		return
	}
	t.Log(resp.Body)

	v2 := map[string]interface{}{
		"json": map[string]interface{}{
			"key": "ivideo_index",
		},
		"proxy": proxy,
	}
	c2 := NewClient(v2)
	resp = c2.Post(jsonUri, nil)
	t.Log(resp.Body)

	v3 := map[string]interface{}{
		"multipart": map[string]interface{}{
			"files": map[string]string{
				"f1": "D:/download/2.sql",
				"f2": "D:/download/1.sql",
			},
			"form_params": map[string]string{
				"key": "ivideo_index",
			},
		},
		"proxy": proxy,
	}
	c3 := NewClient(v3)
	resp = c3.Post(uploadUri, nil)
	t.Log(resp.Body)
}

func TestPost(t *testing.T) {
	v := map[string]interface{}{
		"json": map[string]interface{}{
			"key": "ivideo_index",
		},
	}
	resp := Post(jsonUri, v)
	if nil == resp {
		t.Error("errors", Errors())
	} else {
		t.Log(resp.Body)
	}
}

func TestBaseUri(t *testing.T) {
	v2 := map[string]interface{}{
		"json": map[string]interface{}{
			"key": "ivideo_index",
		},
		"proxy":    proxy,
		"base_uri": "http://192.168.56.100/",
	}
	c2 := NewClient(v2)
	resp := c2.Post("/post.php", nil)
	t.Log("base_uri\t", resp.Body)
}

func TestOption(t *testing.T) {
	v := map[string]interface{}{
		"proxy": proxy,
	}
	option := map[string]interface{}{
		"form_params": map[string]interface{}{
			"key": "ivideo_index",
		},
	}
	c := NewClient(v)
	resp := c.Post(postUri, option)
	t.Log("option\t", resp.Body)

	v2 := map[string]interface{}{
		"json": map[string]interface{}{
			"key": "ivideo_index",
		},
		"proxy": proxy,
	}
	option2 := map[string]interface{}{
		"json": `{"key":"value"}`,
	}
	c2 := NewClient(v2)
	resp = c2.Post(jsonUri, option2)
	t.Log("option\t", resp.Body)

	v3 := map[string]interface{}{
		"multipart": map[string]interface{}{
			"files": map[string]string{
				"f1": "D:/download/2.sql",
				"f2": "D:/download/1.sql",
			},
			"form_params": map[string]string{
				"key": "ivideo_index",
			},
		},
		"proxy": proxy,
	}
	option3 := map[string]interface{}{
		"uploads": map[string]interface{}{
			"files": map[string]string{
				"f1": "D:/download/2.sql",
			},
			"form_params": map[string]string{
				"key": "ivideo_index",
			},
		},
		"proxy": proxy,
	}
	c3 := NewClient(v3)
	resp = c3.Post(uploadUri, option3)
	t.Log("option\t", resp.Body)
}

func TestCookie(t *testing.T) {
	v2 := map[string]interface{}{
		"json": map[string]interface{}{
			"key": "ivideo_index",
		},
		"proxy":   proxy,
		"cookies": true,
	}
	option2 := map[string]interface{}{
		"json": `{"key":"value"}`,
	}
	c2 := NewClient(v2)
	resp := c2.Post(session1, nil)

	option2["cookies"] = resp.Cookies()
	resp = c2.Post(session2Uri, option2)
	t.Log("Cookie\t", resp.Body)
	//close cookies
	option2["cookies"] = false
	c2.CloseCookies()
	resp = c2.Post(session2Uri, option2)
	t.Log("Cookie colse\t", resp.Body)
}

func TestRedirect(t *testing.T) {
	v := map[string]interface{}{
		"form_params": map[string]interface{}{
			"name": "aa",
		},
		"proxy":   proxy,
		"cookies": true,
	}
	option2 := map[string]interface{}{
		"json": `{"key":"value"}`,
	}
	c2 := NewClient(v)
	resp := c2.Post(loginUri, nil)
	option2["cookies"] = c2.GetCookies()
	resp = c2.Post(session2Uri, option2)
	t.Log("Cookie\t", resp.Body)
}

func TestAuth(t *testing.T) {
	v := map[string]interface{}{
		AUTH: []string{"malijie", "pwd"},
	}
	c2 := NewClient(v)
	resp := c2.Post(authUri, nil)
	if nil == resp {
		t.Log(c2.GetErrors())
		return
	}

	t.Log(resp.Body)
}
