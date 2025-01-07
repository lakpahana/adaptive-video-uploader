package ffmpeg

const segmentTime = 5

const createHLSCommand = "ffmpeg -i %s -profile:v baseline -level 3.0 -start_number 0 -hls_time %d -hls_list_size 0 -f hls %s/%s_playlist.m3u8"

const createThumbnailCommand = "ffmpeg -i %s -ss 00:00:01.000 -vframes 1 %s/%s_thumbnail.jpg"
