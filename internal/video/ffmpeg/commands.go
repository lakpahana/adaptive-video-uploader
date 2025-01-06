package ffmpeg

const segmentTime = 5

const createHLSCommand = "ffmpeg -i %s -profile:v baseline -level 3.0 -start_number 0 -hls_time %d -hls_list_size 0 -f hls %s/%s_playlist.m3u8"

const createThumbnailCommand = "ffmpeg -i %s -ss 00:00:01.000 -vframes 1 %s/thumbnail.jpg"

const createDASHCommand = "ffmpeg -i %s -map 0:v:0 -s:v:0 1920x1080 -c:v:0 libx264 -b:v:0 5000k -maxrate:v:0 5350k -bufsize:v:0 7500k -map 0:v:0 -s:v:1 1280x720 -c:v:1 libx264 -b:v:1 3000k -maxrate:v:1 3210k -bufsize:v:1 4500k -map 0:v:0 -s:v:2 854x480 -c:v:2 libx264 -b:v:2 1000k -maxrate:v:2 1070k -bufsize:v:2 1500k -map 0:a:0 -c:a aac -b:a 128k -profile:v:0 high -profile:v:1 high -profile:v:2 high -bf 1 -keyint_min 120 -g 120 -sc_threshold 0 -b_strategy 0 -use_timeline 1 -use_template 1 -seg_duration 4 -init_seg_name init_$RepresentationID$.m4s -media_seg_name chunk_$RepresentationID$_$Number%s$.m4s -f dash %s/manifest.mpd"
