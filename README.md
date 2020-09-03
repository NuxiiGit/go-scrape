# go-scrape

A web scraper written in Go that offers methods of displaying site information.

## Example

Displaying site information as JSON

```
~$ go-scrape https://nuxiigit.github.io/404.html json
{"name":"html","attrs":[],"content":"\n\t\n\t\n","children":[{"name":"head","attrs":[],"content":"\n\t\t\n\t\t\n\t\t\n\t\t\n\t\t\n\t","children":[{"name":"meta","attrs":[{"name":"charset",
"value":"utf-8"},],"content":"","children":[],},{"name":"link","attrs":[{"name":"rel","value":"stylesheet"},{"name":"href","value":"/style/root.css"},],"content":"","children":[],},{"name"
:"link","attrs":[{"name":"rel","value":"stylesheet"},{"name":"href","value":"/style/tabular.css"},],"content":"","children":[],},{"name":"link","attrs":[{"name":"rel","value":"stylesheet"}
,{"name":"href","value":"/style/animations.css"},],"content":"","children":[],},{"name":"style","attrs":[],"content":"\n\t\t\tbody {\n\t\t\t\theight : 100vh;\n\t\t\t}\n\n\t\t\timg:hover {\
n\t\t\t\tanimation : anim-wobble 0.25s linear infinite;\n\t\t\t}\n\t\t","children":[],},],},{"name":"body","attrs":[],"content":"\n\t\t\n\t\t\n\t","children":[{"name":"table","attrs":[{"na
me":"class","value":"flex"},],"content":"\n\t\t\n\t\t","children":[{"name":"tr","attrs":[],"content":"\n\t\t\t\n\t\t","children":[{"name":"td","attrs":[{"name":"class","value":"centring"},
],"content":"\n\t\t\t\t\n\t\t\t\t\n\t\t\t\t\n\t\t\t","children":[{"name":"img","attrs":[{"name":"width","value":"150em"},{"name":"src","value":"/img/avatar-5.png"},],"content":"","children
":[],},{"name":"h1","attrs":[],"content":"404","children":[],},{"name":"p","attrs":[],"content":"The page does not exist.","children":[{"name":"span","attrs":[{"name":"id","value":"url"},]
,"content":"","children":[],},],},],},],},],},{"name":"script","attrs":[],"content":"\n\t\t\tvar href = window.location.href;\n\t\t\tvar a = document.createElement('a');\n\t\t\ta.appendChi
ld(document.createTextNode(`${href}`));\n\t\t\ta.href = href;\n\t\t\tvar span = document.getElementById(\"url\");\n\t\t\tspan.appendChild(document.createTextNode(\"(\"));\n\t\t\tspan.appen
dChild(a);\n\t\t\tspan.appendChild(document.createTextNode(\") \"));\n\t\t","children":[],},],},],}
```
