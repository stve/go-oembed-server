# go-oembed-server

A small microservice to retrieve oembed data from a url and return as JSON.

The server utilizes the [Oembed provider list](http://oembed.com/#section7) from Oembed.com to support a wide array of services.

# Usage

Pass the URL you want to extract data from by passing the `url` parameter:

```
http://localhost:8000/?url=https://vimeo.com/205090959
```

Data is returned as JSON:

```json
{
  "type": "video",
  "url": "https://vimeo.com/205090959",
  "provider_url": "https://vimeo.com/",
  "provider_name": "Vimeo",
  "title": "Symphony of Light - Kauai Timelapse",
  "description": "From the towering green spires to the pristine beaches, the stunning island of Kauai offers an incredible range of unique landscapes to explore. \"Symphony of Light” aims to take a close look at the shapes and forms of the island, showcased through the relationship between light and shadow.\n\nI created a post showing a bunch of behind the scenes images and talking about some of the shots in the video. Please check it out, if you are interested. \nhttp://www.shainblumphoto.com/project/hawaii-timelapse/\n\nDirected by: \nMichael Shainblum\nhttp://www.shainblumphoto.com/\n\nMusic: \nRyan Taubert - Declaration\n\nSpecial Thanks to\nDynamic Perception\nhttps://www.dynamicperception.com/\nEmotimo\nhttp://emotimo.com/\nAndrew Studer\nMatthew Feeser\nPatrycja Podgorski\nMonika Podgorski\nAndrea Pendergast\nMichael Matti",
  "width": 640,
  "height": 272,
  "thumbnail_url": "https://i.vimeocdn.com/video/619779422_640.jpg",
  "thumbnail_width": 640,
  "thumbnail_height": 272,
  "author_name": "Michael Shainblum",
  "author_url": "https://vimeo.com/shainblum",
  "html": "\u003ciframe src=\"https://player.vimeo.com/video/205090959\" width=\"640\" height=\"272\" frameborder=\"0\" title=\"Symphony of Light - Kauai Timelapse\" webkitallowfullscreen mozallowfullscreen allowfullscreen\u003e\u003c/iframe\u003e"
}
```


## License

MIT (c) Steve Agalloco. See [LICENSE](https://github.com/stve/go-oembed-server/blob/master/LICENSE.md) for details.
