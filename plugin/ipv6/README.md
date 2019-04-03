# ipv6

## Name

*ipv6* - disable ipv6 parse if needed.

## Description

Use ipv6 plugin if don't want to parse dns record, add plugin in front of parse/forward plugin, the ipv6 plugin handler will skip all next handlers if dns query type is AAAA.

## Syntax
~~~ txt
ipv6 true/false
~~~
*true* stand for supporting ipv6 parse, false stand for forbiden