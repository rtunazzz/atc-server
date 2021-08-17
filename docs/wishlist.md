# `/Wishlist`

Adds a variant to your wishlist on one of the supported websites.

## Required query params

| Key    | Value                            |
| ------ | -------------------------------- |
| `site` | One of the websites listed below |

**Any additional parameters passed in for any of the supported sites will then be transformed into form fields.**

For example the query `?site=solebox&pid=123&key=value&quantity=1` would add the following input fields:

```html
<form method="POST" action="/add_to_wishlist">
  <input type="hidden" name="pid" value="123" />
  <input type="hidden" name="key" value="value" />
  <input type="hidden" name="quantity" value="1" />
</form>
```

## Overwriting the default form handler
You can also specify the `run_script` and `script_name` properties in the [config.json](../config.json) file. This will then overwrite the default form submit script and instead execute the `script_name`, which is expected to be located in the [scripts](../web/scripts) folder. 

## Supported websites

1. [`snipes`](https://www.snipes.com/)
2. [`onygo`](https://www.onygo.com/)
3. [`nike`](https://www.nike.com/)
