# resource_converter
The resource converter is to convert the resoure with its units to the lowest the unit in the Kubernetes standpoint.

The details about the resource units is available at:

[Resource Management for Pods and Containers](https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/)

The following example is to use resource_converter to convert the memory "1Mi":

```
# echo "1Mi" | ./resource_converter
1048576
```
