# ffmpeg

`ffmpeg -i a.mp4 -itsoffset 2 -i a.m4a -c:v copy -c:a aac -strict experimental b.mp4`

`ffmpeg -i a.mp4 -i a.m4a -c:v copy -c:a aac -strict experimental b.mp4`

## converting videos

`ffmpeg -i a.mkv -codec copy a.mp4`
