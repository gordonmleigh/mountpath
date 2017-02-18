# Mountpath

This library/CLI will return the mount path of a device if it is mounted.


## Library 

### GetMountPath

This function will look in `/proc/mounts` to search for the given device 

    func GetMountPath(dev string) (string, error)

### GetMountPathEx

This function will look in the given file to search for the given device 

    func GetMountPath(dev string, mounts string) (string, error)
