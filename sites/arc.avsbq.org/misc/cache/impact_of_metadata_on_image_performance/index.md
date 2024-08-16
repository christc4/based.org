# Metadata

Saw this blog post on metadata from [dexecure](https://dexecure.com/blog/impact-of-metadata-on-image-performance/)

Let me sift through the information.

- Dexecure used date from the [HTTP Archive](https://httparchive.org/)
and this BigQuery thing

 to generate list of URLs of valid JPEG images.

They used the HTTP Archive craws from '2016/08/01' which contained 4.3 million JPEG images from the top 500,000 sites, totalling up to 195GB. 

They then downloaded and analysed metadata from this dataset.

38.9% of images had some sort of metadata.

On average 15.8% of total image size was occupied by metadata. 

If each of these top sites were just visited once, nearly 13 Gb of data could  have been saved.

## Formats to store metadata

Most common formats are,

- EXIF
	- Info added by camera manufacturers containing settings used when taking the photo

- XMP
	- XML stuff Adobe inserts into pics

- 8BIM
	- inserted by Photoshop

- IPTC
	- user added metadata

- ICC
	- colour profiles embedded in pics

## Do people really store such stuff in images?

1,674,831 images, 39.9% of the images analysed stored some sort of metadata.

![image](average-jpeg-file.webp)

Here is a distribution showing the different formats used to store the metadata and how much size they occupy on an average

![image](formats-used-to-store-metadata.webp)

## The Usual Suspects

These were the fifty most common attributes stored as EXIF data
![image](50-most-common-exif-attributes.webp)

As you can see most of this information is not required by the browser to render the image. So you might be asking "why is sending metadata with images still a thing"

Before we start stripping away all the metadata from our images, there are some use-cases you might have to consider.

### Thumbnails

Some software store thumbnails of the image as EXIF data too. So you can have a JPEG image as a thumbnail within another JPEG image.

Browsers don't make use of this data, so it should be stripped out if it's intended to be part of a webpage, 

### Copyright

One of the primary use cases that we have seen stored as IPTC is the name of the photographer and other copyright information. This might be important depending on your use case though most social sharing sites automatically strip out this information. 

### Orientation 

EXIF data can also contain orientation information of the picture and browsers take cues from this to orient the image correctly. We have seen 9.1% of images store orientation information along as part of the EXIF data. This attribute comes into play only when you visit the image directly - say, by visiting the image directly via its URL. When you are embedding the image directly in the webpage, say via an image tag, this attribute has no effect at all (except in mobile Safari). IE <= 11 doesn’t recognize this attribute in both cases.

Stripping out this attribute seems to be a good idea to be consistent. That way users see the image oriented the same way regardless of whether they visit the image directly or see it as part of a webpage.

### Color Profiles

This seems to be the strongest use case for storing metadata in JPEGs. 10% of the images we analysed had an embedded colour profile and the profiles were as large as 7KB on an average. Images can store these colour profiles as ICC metadata. They help images look consistent depending on the capabilities of the device used when viewing them.

However browser support for these ICC profiles has been very flaky and there aren’t many screens which can take advantage of these embedded profiles.

By default, Dexecure does not strip off ICC profiles from JPEG images since they can lead to different looking images in some cases.

## Other disadvantages

### Longer Time to Layout and Render
In a JPEG file, EXIF metadata appears before other information such as the frame containing the height and width of the image and the scans containing the image data itself. Browsers use the height and width information in the image to help layout the page and the earlier the browser finds this information, the better.

It would also take a longer time for the browser to render the image since the browser would have to download all the metadata in the JPEG file before it starts getting the scans containing the image data. Chrome prioritises the download of the first image that it sees in a webpage and you would be wasting crucial bytes sending metadata, instead of the image content itself. This especially matters during the initial stages of the page load where the TCP window is quite small due to TCP slow start.

### Wider Surface for Security Attacks
There have been quite a few attacks - both server side and on the client side because of improper handling of image metadata. As with any other user provided data, make sure you are sanitising it properly before using them in a security sensitive context.

### Privacy Impact

You can read about how the GPS information stored in one of John McAfee’s pictures almost gave his location away here and how much more recently, Apple inadvertently revealed more than they had probably intended with one of the default desktop images shipped in the recent Mac OS X release here.

Most of the metadata can be stripped for images that are intended to be delivered on the web. Of course, removing metadata from images is one of the techniques Dexecure uses to produce leaner images.
