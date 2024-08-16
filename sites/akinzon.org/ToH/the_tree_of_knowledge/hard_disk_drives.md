# Hard Disk Drives

Up: [External Memory](external_memory)
Brother(s): [Solid State Drives](solid_state_drives)
TARGET DECK
Leaves of the Tree of Ohara::Branch I::Semester I::1203 CompSys::Computer Systems

![Pasted image 20231129123953.png](pasted_image_20231129123953.png)

What are Hard disks and Solid State drives?
Hard disks are a form of [Magnetic Storage](magnetic_storage) where the disk is covered in a magnetisable material, typically iron oxide.

They trace their origins back to the Winchester Hard Disk. Just like disks today, they were made of platters with heads that would be situated a little above the disk.

**Some Hard disk key terms:**

**Track** - Tracks are the concentric rings in which data is stored
	> There are gaps between tracks
	> These are to reduce errors due to the [Read and Write Head|read/write head ](read_and_write_head|read/write_head_) interfering with the wrong track
	> There are the same number of bits per track to keep read/write speeds consistent where constant angular velocity is used
	
**Sector** - Sectors are what tracks are divided into; data is transferred to/from the disk in sectors.
	> Usually of a fixed length
	> Also includes error correction and other data
	
**Cylinder** - The same track on the same sector of different disks
	> Used so that the disk heads do not need to move, which increases access speed
	
What are tracks, sectors and cylinders #flashcard 
Tracks are the concentric rings in which data is stored
Sectors are what tracks are divided into; data is transferred to/from the disk in sectors.
Cylinders are the same track on the same sector of different disks
<!--ID: 1701369224408-->



#### Features & Benefits of HDDs
 - Cheap
 - Decently fast (especially with [RAID](raid))
 - Only getting larger
 - Come with an onboard processor/cache


#### What determines HDD access times and speeds?

These are determined by:
	> **Seek time** - Time it takes for the read/write head to move to the correct track
	> **Rotational latency** - Time it takes for the correct track to move to the read/write head
	> Access time =  seek time + rotational latency
	> Must also factor in the time it takes to transfer data.


#### Typical HDD speeds

Off disk - 150MB/s, but 6Gbps using [SATA](sata)
All this means is that there are different speeds and optimisations that vary from disk to disk

#### On-Disk [Cache](cache)

Modern disks have ~ 8-128 MB cache to store whole tracks by caching what was just read/written

This acts as a buffer between the disk and external I/O

Frequently used data can be retrieved from cache rather than just continuously reading from the disk.

What are the benefits of on-disk [Cache](cache) #flashcard 
It acts as a buffer between the disk and external I/O, allowing for data to be retrieved from cache instead of the disk which is faster.
#### Current disks in the market
<!--ID: 1701369224414-->


 - The largest are around 22TB
 - They can have a spin of up to 10-15k rpm
 - They can also have a data rate of 150-250 MiB/s
 - Average access time of 3.6 - 8 ms
 - Use around 10W power

MTBF - Mean Time Between Failures


























#### Why:
#### How:









