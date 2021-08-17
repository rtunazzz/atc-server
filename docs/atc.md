# `/ATC`

Adds a variant to your cart on one of the supported websites.

## Required query params

| Key    | Value                            |
| ------ | -------------------------------- |
| `site` | One of the websites listed below |

**Any additional parameters passed in for any of the supported sites will then be transformed into form fields.**

For example the query `?site=solebox&pid=123&key=value&quantity=1` would add the following input fields:

```html
<form method="POST" action="/add_to_cart">
  <input type="hidden" name="pid" value="123" />
  <input type="hidden" name="key" value="value" />
  <input type="hidden" name="quantity" value="1" />
</form>
```

## Supported websites

1. [`solebox`](https://www.solebox.com/)
2. [`snipes`](https://www.snipes.com/)
3. [`onygo`](https://www.onygo.com/)
4. [`NBB`](https://www.notebooksbilliger.de/)
5. [`nprague`](https://www.nprague.cz/)
