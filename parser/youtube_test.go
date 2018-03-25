package parser

import (
	"encoding/json"
	"testing"
)

type mockExtractorFormatSelector struct{}

func (m mockExtractorFormatSelector) Extract(string) (resp *youtubeResponse) {
	resp = &youtubeResponse{
		Data:  make(chan *youtubeVideo),
		Error: make(chan error),
	}

	go func() {
		s := &youtubeVideo{}
		if err := json.Unmarshal(mock, s); err != nil {
			resp.Error <- err
			return
		}

		resp.Data <- s
	}()

	return
}

func (mockExtractorFormatSelector) Select(v *youtubeVideo, q YoutubeQuality) {
	youtubeParser{}.Select(v, q)
}

func init() {
	es = &mockExtractorFormatSelector{}
}

func TestErrorOnEmptyUrl(t *testing.T) {
	if err, _ := Parse("", QualityBest); err != nil {
		t.Error(err)
	}
}

func TestParseWorst(t *testing.T) {
	data, err := Parse("mock", QualityWorst)
	if err != nil {
		t.Error(err)
	}
	if data.FormatId != "249" {
		t.Error("Wrong format")
	}
}

func TestParseBest(t *testing.T)  {
	data, err := Parse("mock", QualityBest)
	if err != nil {
		t.Error(err)
	}
	if data.FormatId != "171" {
		t.Error("Wrong format")
	}
}

var mock = []byte("{\"id\": \"NE71LIUJ4gQ\", \"uploader\": \"YouTube HD Brasil\", \"uploader_id\": \"diegomarlon1\", \"uploader_url\": \"http://www.youtube.com/user/diegomarlon1\", \"upload_date\": \"20120815\", \"license\": \"Standard YouTube License\", \"creator\": null, \"title\": \"Bragaboys - Para dan\\u00e7ar isso aqui \\u00e9 (Bomba)\", \"alt_title\": null, \"thumbnail\": \"https://i.ytimg.com/vi/NE71LIUJ4gQ/maxresdefault.jpg\", \"description\": \"* PLEASE NOTE: SEE THIS LINK TO NEW VERSION\\nhttps://www.youtube.com/watch?v=6Axe32Q5B1k\", \"categories\": [\"Music\"], \"tags\": [\"bragaboys\", \"ax\\u00e9 anos 90\", \"anos 90\", \"ax\\u00e9\", \"Ax\\u00e9 Bahia\", \"Ax\\u00e9 antigas\", \"ax\\u00e9 music\", \"o movimento e sexy\", \"um movimento sexy\", \"movimento \\u00e9 sexy\", \"o movimento \\u00e9 bem sexy\", \"carnaval\", \"em destaque\", \"Tchakabum\", \"King Africa\", \"axe\", \"em baixo\", \"em cima\", \"Para dan\\u00e7ar isso aqui \\u00e9\", \"suavecito para abajo\", \"para abajo\", \"Un movimento muy sexy\", \"flashback\", \"musicas antigas mais tocadas\", \"festa\", \"ache\", \"axe download\", \"la bomba\", \"bomba\", \"devagarinho\", \"at\\u00e9\", \"Brazilian music\", \"m\\u00fasica de Brasil\", \"m\\u00fasica brasile\\u00f1a\", \"brasil\", \"brazil\", \"music\"], \"subtitles\": {}, \"automatic_captions\": {}, \"duration\": 129, \"age_limit\": 0, \"annotations\": null, \"chapters\": null, \"webpage_url\": \"https://www.youtube.com/watch?v=NE71LIUJ4gQ\", \"view_count\": 10329037, \"like_count\": 42391, \"dislike_count\": 2087, \"average_rating\": 4.81231164932, \"formats\": [{\"format_id\": \"249\", \"url\": \"https://r1---sn-oxunxg8pjvn-nw2e.googlevideo.com/videoplayback?clen=823009&gir=yes&keepalive=yes&mime=audio%2Fwebm&key=yt6&signature=5A54AC13A4DB3D90D74CCA1483C1FDC5F959A931.CC89469B9CC844B4247C02C0F50826ACC72BAD76&ei=WqK1WouIH4arxgS3wp34Bg&c=WEB&source=youtube&fvip=2&itag=249&id=o-AN4EhYYtTRVcbehrgICK0l6t3isae2fVTJxuy7Mzeg2Y&dur=128.741&requiressl=yes&mm=31%2C29&pl=20&mn=sn-oxunxg8pjvn-nw2e%2Csn-bg07dnsl&ms=au%2Crdu&mt=1521852891&mv=m&ip=189.32.180.245&initcwndbps=422500&lmt=1518268117695507&ipbits=0&expire=1521874618&sparams=clen%2Cdur%2Cei%2Cgir%2Cid%2Cinitcwndbps%2Cip%2Cipbits%2Citag%2Ckeepalive%2Clmt%2Cmime%2Cmm%2Cmn%2Cms%2Cmv%2Cpl%2Crequiressl%2Csource%2Cexpire&ratebypass=yes\", \"player_url\": \"/yts/jsbin/player-vflMfSEyN/en_US/base.js\", \"ext\": \"webm\", \"format_note\": \"DASH audio\", \"acodec\": \"opus\", \"abr\": 50, \"filesize\": 823009, \"tbr\": 54.155, \"vcodec\": \"none\", \"downloader_options\": {\"http_chunk_size\": 10485760}, \"format\": \"249 - audio only (DASH audio)\", \"protocol\": \"https\", \"http_headers\": {\"User-Agent\": \"Mozilla/5.0 (X11; Linux x86_64; rv:59.0) Gecko/20100101 Firefox/59.0 (Chrome)\", \"Accept-Charset\": \"ISO-8859-1,utf-8;q=0.7,*;q=0.7\", \"Accept\": \"text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\", \"Accept-Encoding\": \"gzip, deflate\", \"Accept-Language\": \"en-us,en;q=0.5\"}}, {\"format_id\": \"250\", \"url\": \"https://r1---sn-oxunxg8pjvn-nw2e.googlevideo.com/videoplayback?clen=1087599&gir=yes&keepalive=yes&mime=audio%2Fwebm&key=yt6&signature=2B74CEF937C7F55FB7B4FC91819DD3082595B8A4.836FC9DE1C147DD3AC9C9EF7E2206E9F45FFA66E&ei=WqK1WouIH4arxgS3wp34Bg&c=WEB&source=youtube&fvip=2&itag=250&id=o-AN4EhYYtTRVcbehrgICK0l6t3isae2fVTJxuy7Mzeg2Y&dur=128.741&requiressl=yes&mm=31%2C29&pl=20&mn=sn-oxunxg8pjvn-nw2e%2Csn-bg07dnsl&ms=au%2Crdu&mt=1521852891&mv=m&ip=189.32.180.245&initcwndbps=422500&lmt=1518268111312813&ipbits=0&expire=1521874618&sparams=clen%2Cdur%2Cei%2Cgir%2Cid%2Cinitcwndbps%2Cip%2Cipbits%2Citag%2Ckeepalive%2Clmt%2Cmime%2Cmm%2Cmn%2Cms%2Cmv%2Cpl%2Crequiressl%2Csource%2Cexpire&ratebypass=yes\", \"player_url\": \"/yts/jsbin/player-vflMfSEyN/en_US/base.js\", \"ext\": \"webm\", \"format_note\": \"DASH audio\", \"acodec\": \"opus\", \"abr\": 70, \"filesize\": 1087599, \"tbr\": 71.972, \"vcodec\": \"none\", \"downloader_options\": {\"http_chunk_size\": 10485760}, \"format\": \"250 - audio only (DASH audio)\", \"protocol\": \"https\", \"http_headers\": {\"User-Agent\": \"Mozilla/5.0 (X11; Linux x86_64; rv:59.0) Gecko/20100101 Firefox/59.0 (Chrome)\", \"Accept-Charset\": \"ISO-8859-1,utf-8;q=0.7,*;q=0.7\", \"Accept\": \"text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\", \"Accept-Encoding\": \"gzip, deflate\", \"Accept-Language\": \"en-us,en;q=0.5\"}}, {\"format_id\": \"140\", \"url\": \"https://r1---sn-oxunxg8pjvn-nw2e.googlevideo.com/videoplayback?clen=2045939&gir=yes&keepalive=yes&mime=audio%2Fmp4&key=yt6&signature=BF18E65588607240C1E8CA8294CA4EA17BE608C6.980E52603BD343DC3FD3D9FFA53E7EC137881730&ei=WqK1WouIH4arxgS3wp34Bg&c=WEB&source=youtube&fvip=2&itag=140&id=o-AN4EhYYtTRVcbehrgICK0l6t3isae2fVTJxuy7Mzeg2Y&dur=128.777&requiressl=yes&mm=31%2C29&pl=20&mn=sn-oxunxg8pjvn-nw2e%2Csn-bg07dnsl&ms=au%2Crdu&mt=1521852891&mv=m&ip=189.32.180.245&initcwndbps=422500&lmt=1518266363468878&ipbits=0&expire=1521874618&sparams=clen%2Cdur%2Cei%2Cgir%2Cid%2Cinitcwndbps%2Cip%2Cipbits%2Citag%2Ckeepalive%2Clmt%2Cmime%2Cmm%2Cmn%2Cms%2Cmv%2Cpl%2Crequiressl%2Csource%2Cexpire&ratebypass=yes\", \"player_url\": \"/yts/jsbin/player-vflMfSEyN/en_US/base.js\", \"ext\": \"m4a\", \"format_note\": \"DASH audio\", \"acodec\": \"mp4a.40.2\", \"abr\": 128, \"container\": \"m4a_dash\", \"filesize\": 2045939, \"tbr\": 128.064, \"vcodec\": \"none\", \"downloader_options\": {\"http_chunk_size\": 10485760}, \"format\": \"140 - audio only (DASH audio)\", \"protocol\": \"https\", \"http_headers\": {\"User-Agent\": \"Mozilla/5.0 (X11; Linux x86_64; rv:59.0) Gecko/20100101 Firefox/59.0 (Chrome)\", \"Accept-Charset\": \"ISO-8859-1,utf-8;q=0.7,*;q=0.7\", \"Accept\": \"text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\", \"Accept-Encoding\": \"gzip, deflate\", \"Accept-Language\": \"en-us,en;q=0.5\"}}, {\"format_id\": \"251\", \"url\": \"https://r1---sn-oxunxg8pjvn-nw2e.googlevideo.com/videoplayback?clen=2164611&gir=yes&keepalive=yes&mime=audio%2Fwebm&key=yt6&signature=2F152E3E5567A3858A1032C43B4CEA2037687E1D.B75B14067BC1FFAB0959433F3880E39345EC2477&ei=WqK1WouIH4arxgS3wp34Bg&c=WEB&source=youtube&fvip=2&itag=251&id=o-AN4EhYYtTRVcbehrgICK0l6t3isae2fVTJxuy7Mzeg2Y&dur=128.741&requiressl=yes&mm=31%2C29&pl=20&mn=sn-oxunxg8pjvn-nw2e%2Csn-bg07dnsl&ms=au%2Crdu&mt=1521852891&mv=m&ip=189.32.180.245&initcwndbps=422500&lmt=1518268107127571&ipbits=0&expire=1521874618&sparams=clen%2Cdur%2Cei%2Cgir%2Cid%2Cinitcwndbps%2Cip%2Cipbits%2Citag%2Ckeepalive%2Clmt%2Cmime%2Cmm%2Cmn%2Cms%2Cmv%2Cpl%2Crequiressl%2Csource%2Cexpire&ratebypass=yes\", \"player_url\": \"/yts/jsbin/player-vflMfSEyN/en_US/base.js\", \"ext\": \"webm\", \"format_note\": \"DASH audio\", \"acodec\": \"opus\", \"abr\": 160, \"filesize\": 2164611, \"tbr\": 144.238, \"vcodec\": \"none\", \"downloader_options\": {\"http_chunk_size\": 10485760}, \"format\": \"251 - audio only (DASH audio)\", \"protocol\": \"https\", \"http_headers\": {\"User-Agent\": \"Mozilla/5.0 (X11; Linux x86_64; rv:59.0) Gecko/20100101 Firefox/59.0 (Chrome)\", \"Accept-Charset\": \"ISO-8859-1,utf-8;q=0.7,*;q=0.7\", \"Accept\": \"text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\", \"Accept-Encoding\": \"gzip, deflate\", \"Accept-Language\": \"en-us,en;q=0.5\"}}, {\"format_id\": \"171\", \"url\": \"https://r1---sn-oxunxg8pjvn-nw2e.googlevideo.com/videoplayback?clen=2294642&gir=yes&keepalive=yes&mime=audio%2Fwebm&key=yt6&signature=0169F87DB2F53BE1A2439441B463FD7B31086B40.653300D24CA3181E3222325EB0087741648C77A0&ei=WqK1WouIH4arxgS3wp34Bg&c=WEB&source=youtube&fvip=2&itag=171&id=o-AN4EhYYtTRVcbehrgICK0l6t3isae2fVTJxuy7Mzeg2Y&dur=128.732&requiressl=yes&mm=31%2C29&pl=20&mn=sn-oxunxg8pjvn-nw2e%2Csn-bg07dnsl&ms=au%2Crdu&mt=1521852891&mv=m&ip=189.32.180.245&initcwndbps=422500&lmt=1518268125409651&ipbits=0&expire=1521874618&sparams=clen%2Cdur%2Cei%2Cgir%2Cid%2Cinitcwndbps%2Cip%2Cipbits%2Citag%2Ckeepalive%2Clmt%2Cmime%2Cmm%2Cmn%2Cms%2Cmv%2Cpl%2Crequiressl%2Csource%2Cexpire&ratebypass=yes\", \"player_url\": \"/yts/jsbin/player-vflMfSEyN/en_US/base.js\", \"ext\": \"webm\", \"acodec\": \"vorbis\", \"format_note\": \"DASH audio\", \"abr\": 128, \"filesize\": 2294642, \"tbr\": 159.289, \"vcodec\": \"none\", \"downloader_options\": {\"http_chunk_size\": 10485760}, \"format\": \"171 - audio only (DASH audio)\", \"protocol\": \"https\", \"http_headers\": {\"User-Agent\": \"Mozilla/5.0 (X11; Linux x86_64; rv:59.0) Gecko/20100101 Firefox/59.0 (Chrome)\", \"Accept-Charset\": \"ISO-8859-1,utf-8;q=0.7,*;q=0.7\", \"Accept\": \"text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\", \"Accept-Encoding\": \"gzip, deflate\", \"Accept-Language\": \"en-us,en;q=0.5\"}}, {\"format_id\": \"278\", \"url\": \"https://r1---sn-oxunxg8pjvn-nw2e.googlevideo.com/videoplayback?clen=1515568&gir=yes&keepalive=yes&aitags=133%2C134%2C135%2C136%2C137%2C160%2C242%2C243%2C244%2C247%2C248%2C278&mime=video%2Fwebm&key=yt6&itag=278&ei=WqK1WouIH4arxgS3wp34Bg&c=WEB&source=youtube&fvip=2&signature=50172A46F2B021B08F582C0A7ED85FB6D52BA589.B206F98D10BD94CD33097B72850A6E41E440D269&id=o-AN4EhYYtTRVcbehrgICK0l6t3isae2fVTJxuy7Mzeg2Y&dur=128.728&requiressl=yes&mm=31%2C29&pl=20&mn=sn-oxunxg8pjvn-nw2e%2Csn-bg07dnsl&ms=au%2Crdu&mt=1521852891&mv=m&ip=189.32.180.245&initcwndbps=422500&lmt=1518268836802700&ipbits=0&expire=1521874618&sparams=aitags%2Cclen%2Cdur%2Cei%2Cgir%2Cid%2Cinitcwndbps%2Cip%2Cipbits%2Citag%2Ckeepalive%2Clmt%2Cmime%2Cmm%2Cmn%2Cms%2Cmv%2Cpl%2Crequiressl%2Csource%2Cexpire&ratebypass=yes\", \"player_url\": \"/yts/jsbin/player-vflMfSEyN/en_US/base.js\", \"ext\": \"webm\", \"height\": 144, \"format_note\": \"144p\", \"container\": \"webm\", \"vcodec\": \"vp9\", \"filesize\": 1515568, \"tbr\": 98.249, \"width\": 256, \"fps\": 30, \"acodec\": \"none\", \"downloader_options\": {\"http_chunk_size\": 10485760}, \"format\": \"278 - 256x144 (144p)\", \"protocol\": \"https\", \"http_headers\": {\"User-Agent\": \"Mozilla/5.0 (X11; Linux x86_64; rv:59.0) Gecko/20100101 Firefox/59.0 (Chrome)\", \"Accept-Charset\": \"ISO-8859-1,utf-8;q=0.7,*;q=0.7\", \"Accept\": \"text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\", \"Accept-Encoding\": \"gzip, deflate\", \"Accept-Language\": \"en-us,en;q=0.5\"}}, {\"format_id\": \"160\", \"url\": \"https://r1---sn-oxunxg8pjvn-nw2e.googlevideo.com/videoplayback?clen=1847824&gir=yes&keepalive=yes&aitags=133%2C134%2C135%2C136%2C137%2C160%2C242%2C243%2C244%2C247%2C248%2C278&mime=video%2Fmp4&key=yt6&itag=160&ei=WqK1WouIH4arxgS3wp34Bg&c=WEB&source=youtube&fvip=2&signature=726C6F278914D0339CA478EF1D44AC0867FE1F1E.0E88C4B82D86DCE02FD4170B6DEE2BA14D69050E&id=o-AN4EhYYtTRVcbehrgICK0l6t3isae2fVTJxuy7Mzeg2Y&dur=128.728&requiressl=yes&mm=31%2C29&pl=20&mn=sn-oxunxg8pjvn-nw2e%2Csn-bg07dnsl&ms=au%2Crdu&mt=1521852891&mv=m&ip=189.32.180.245&initcwndbps=422500&lmt=1518266664482013&ipbits=0&expire=1521874618&sparams=aitags%2Cclen%2Cdur%2Cei%2Cgir%2Cid%2Cinitcwndbps%2Cip%2Cipbits%2Citag%2Ckeepalive%2Clmt%2Cmime%2Cmm%2Cmn%2Cms%2Cmv%2Cpl%2Crequiressl%2Csource%2Cexpire&ratebypass=yes\", \"player_url\": \"/yts/jsbin/player-vflMfSEyN/en_US/base.js\", \"ext\": \"mp4\", \"height\": 144, \"format_note\": \"144p\", \"vcodec\": \"avc1.4d400c\", \"filesize\": 1847824, \"tbr\": 130.265, \"width\": 256, \"fps\": 30, \"acodec\": \"none\", \"downloader_options\": {\"http_chunk_size\": 10485760}, \"format\": \"160 - 256x144 (144p)\", \"protocol\": \"https\", \"http_headers\": {\"User-Agent\": \"Mozilla/5.0 (X11; Linux x86_64; rv:59.0) Gecko/20100101 Firefox/59.0 (Chrome)\", \"Accept-Charset\": \"ISO-8859-1,utf-8;q=0.7,*;q=0.7\", \"Accept\": \"text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\", \"Accept-Encoding\": \"gzip, deflate\", \"Accept-Language\": \"en-us,en;q=0.5\"}}, {\"format_id\": \"242\", \"url\": \"https://r1---sn-oxunxg8pjvn-nw2e.googlevideo.com/videoplayback?clen=3036005&gir=yes&keepalive=yes&aitags=133%2C134%2C135%2C136%2C137%2C160%2C242%2C243%2C244%2C247%2C248%2C278&mime=video%2Fwebm&key=yt6&itag=242&ei=WqK1WouIH4arxgS3wp34Bg&c=WEB&source=youtube&fvip=2&signature=B45D6C2BC74BFB1C5CF3526AE1A6436B43C03CB4.0F6DD1DADC1480581AD9417F6A3BC58EADBB5303&id=o-AN4EhYYtTRVcbehrgICK0l6t3isae2fVTJxuy7Mzeg2Y&dur=128.728&requiressl=yes&mm=31%2C29&pl=20&mn=sn-oxunxg8pjvn-nw2e%2Csn-bg07dnsl&ms=au%2Crdu&mt=1521852891&mv=m&ip=189.32.180.245&initcwndbps=422500&lmt=1518268840231597&ipbits=0&expire=1521874618&sparams=aitags%2Cclen%2Cdur%2Cei%2Cgir%2Cid%2Cinitcwndbps%2Cip%2Cipbits%2Citag%2Ckeepalive%2Clmt%2Cmime%2Cmm%2Cmn%2Cms%2Cmv%2Cpl%2Crequiressl%2Csource%2Cexpire&ratebypass=yes\", \"player_url\": \"/yts/jsbin/player-vflMfSEyN/en_US/base.js\", \"ext\": \"webm\", \"height\": 240, \"format_note\": \"240p\", \"vcodec\": \"vp9\", \"filesize\": 3036005, \"tbr\": 221.508, \"width\": 426, \"fps\": 30, \"acodec\": \"none\", \"downloader_options\": {\"http_chunk_size\": 10485760}, \"format\": \"242 - 426x240 (240p)\", \"protocol\": \"https\", \"http_headers\": {\"User-Agent\": \"Mozilla/5.0 (X11; Linux x86_64; rv:59.0) Gecko/20100101 Firefox/59.0 (Chrome)\", \"Accept-Charset\": \"ISO-8859-1,utf-8;q=0.7,*;q=0.7\", \"Accept\": \"text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\", \"Accept-Encoding\": \"gzip, deflate\", \"Accept-Language\": \"en-us,en;q=0.5\"}}, {\"format_id\": \"133\", \"url\": \"https://r1---sn-oxunxg8pjvn-nw2e.googlevideo.com/videoplayback?clen=5348807&gir=yes&keepalive=yes&aitags=133%2C134%2C135%2C136%2C137%2C160%2C242%2C243%2C244%2C247%2C248%2C278&mime=video%2Fmp4&key=yt6&itag=133&ei=WqK1WouIH4arxgS3wp34Bg&c=WEB&source=youtube&fvip=2&signature=A795EB4AE165A5785288121B13E45A57D20D9141.37B208E3ECC6A7DC09C42135C6204C3247FC02BC&id=o-AN4EhYYtTRVcbehrgICK0l6t3isae2fVTJxuy7Mzeg2Y&dur=128.728&requiressl=yes&mm=31%2C29&pl=20&mn=sn-oxunxg8pjvn-nw2e%2Csn-bg07dnsl&ms=au%2Crdu&mt=1521852891&mv=m&ip=189.32.180.245&initcwndbps=422500&lmt=1518266670894604&ipbits=0&expire=1521874618&sparams=aitags%2Cclen%2Cdur%2Cei%2Cgir%2Cid%2Cinitcwndbps%2Cip%2Cipbits%2Citag%2Ckeepalive%2Clmt%2Cmime%2Cmm%2Cmn%2Cms%2Cmv%2Cpl%2Crequiressl%2Csource%2Cexpire&ratebypass=yes\", \"player_url\": \"/yts/jsbin/player-vflMfSEyN/en_US/base.js\", \"ext\": \"mp4\", \"height\": 240, \"format_note\": \"240p\", \"vcodec\": \"avc1.4d4015\", \"filesize\": 5348807, \"tbr\": 401.902, \"width\": 426, \"fps\": 30, \"acodec\": \"none\", \"downloader_options\": {\"http_chunk_size\": 10485760}, \"format\": \"133 - 426x240 (240p)\", \"protocol\": \"https\", \"http_headers\": {\"User-Agent\": \"Mozilla/5.0 (X11; Linux x86_64; rv:59.0) Gecko/20100101 Firefox/59.0 (Chrome)\", \"Accept-Charset\": \"ISO-8859-1,utf-8;q=0.7,*;q=0.7\", \"Accept\": \"text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\", \"Accept-Encoding\": \"gzip, deflate\", \"Accept-Language\": \"en-us,en;q=0.5\"}}, {\"format_id\": \"243\", \"url\": \"https://r1---sn-oxunxg8pjvn-nw2e.googlevideo.com/videoplayback?clen=5520996&gir=yes&keepalive=yes&aitags=133%2C134%2C135%2C136%2C137%2C160%2C242%2C243%2C244%2C247%2C248%2C278&mime=video%2Fwebm&key=yt6&itag=243&ei=WqK1WouIH4arxgS3wp34Bg&c=WEB&source=youtube&fvip=2&signature=259AB211A3972121697E3FAECB5FC023F61C94A8.2C8266E380C3B9E7C136050174AE56A22B264500&id=o-AN4EhYYtTRVcbehrgICK0l6t3isae2fVTJxuy7Mzeg2Y&dur=128.728&requiressl=yes&mm=31%2C29&pl=20&mn=sn-oxunxg8pjvn-nw2e%2Csn-bg07dnsl&ms=au%2Crdu&mt=1521852891&mv=m&ip=189.32.180.245&initcwndbps=422500&lmt=1518268923445452&ipbits=0&expire=1521874618&sparams=aitags%2Cclen%2Cdur%2Cei%2Cgir%2Cid%2Cinitcwndbps%2Cip%2Cipbits%2Citag%2Ckeepalive%2Clmt%2Cmime%2Cmm%2Cmn%2Cms%2Cmv%2Cpl%2Crequiressl%2Csource%2Cexpire&ratebypass=yes\", \"player_url\": \"/yts/jsbin/player-vflMfSEyN/en_US/base.js\", \"ext\": \"webm\", \"height\": 360, \"format_note\": \"360p\", \"vcodec\": \"vp9\", \"filesize\": 5520996, \"tbr\": 402.484, \"width\": 640, \"fps\": 30, \"acodec\": \"none\", \"downloader_options\": {\"http_chunk_size\": 10485760}, \"format\": \"243 - 640x360 (360p)\", \"protocol\": \"https\", \"http_headers\": {\"User-Agent\": \"Mozilla/5.0 (X11; Linux x86_64; rv:59.0) Gecko/20100101 Firefox/59.0 (Chrome)\", \"Accept-Charset\": \"ISO-8859-1,utf-8;q=0.7,*;q=0.7\", \"Accept\": \"text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\", \"Accept-Encoding\": \"gzip, deflate\", \"Accept-Language\": \"en-us,en;q=0.5\"}}, {\"format_id\": \"244\", \"url\": \"https://r1---sn-oxunxg8pjvn-nw2e.googlevideo.com/videoplayback?clen=8966947&gir=yes&keepalive=yes&aitags=133%2C134%2C135%2C136%2C137%2C160%2C242%2C243%2C244%2C247%2C248%2C278&mime=video%2Fwebm&key=yt6&itag=244&ei=WqK1WouIH4arxgS3wp34Bg&c=WEB&source=youtube&fvip=2&signature=88D975230390ECE9FCA41C278E4367F221897D1D.76B50A8E844D4C944E6153880E0DBE22E64BF2DB&id=o-AN4EhYYtTRVcbehrgICK0l6t3isae2fVTJxuy7Mzeg2Y&dur=128.728&requiressl=yes&mm=31%2C29&pl=20&mn=sn-oxunxg8pjvn-nw2e%2Csn-bg07dnsl&ms=au%2Crdu&mt=1521852891&mv=m&ip=189.32.180.245&initcwndbps=422500&lmt=1518268918613448&ipbits=0&expire=1521874618&sparams=aitags%2Cclen%2Cdur%2Cei%2Cgir%2Cid%2Cinitcwndbps%2Cip%2Cipbits%2Citag%2Ckeepalive%2Clmt%2Cmime%2Cmm%2Cmn%2Cms%2Cmv%2Cpl%2Crequiressl%2Csource%2Cexpire&ratebypass=yes\", \"player_url\": \"/yts/jsbin/player-vflMfSEyN/en_US/base.js\", \"ext\": \"webm\", \"height\": 480, \"format_note\": \"480p\", \"vcodec\": \"vp9\", \"filesize\": 8966947, \"tbr\": 733.152, \"width\": 854, \"fps\": 30, \"acodec\": \"none\", \"downloader_options\": {\"http_chunk_size\": 10485760}, \"format\": \"244 - 854x480 (480p)\", \"protocol\": \"https\", \"http_headers\": {\"User-Agent\": \"Mozilla/5.0 (X11; Linux x86_64; rv:59.0) Gecko/20100101 Firefox/59.0 (Chrome)\", \"Accept-Charset\": \"ISO-8859-1,utf-8;q=0.7,*;q=0.7\", \"Accept\": \"text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\", \"Accept-Encoding\": \"gzip, deflate\", \"Accept-Language\": \"en-us,en;q=0.5\"}}, {\"format_id\": \"134\", \"url\": \"https://r1---sn-oxunxg8pjvn-nw2e.googlevideo.com/videoplayback?clen=11475754&gir=yes&keepalive=yes&aitags=133%2C134%2C135%2C136%2C137%2C160%2C242%2C243%2C244%2C247%2C248%2C278&mime=video%2Fmp4&key=yt6&itag=134&ei=WqK1WouIH4arxgS3wp34Bg&c=WEB&source=youtube&fvip=2&signature=0BF30CF567BB331794DB4CDD24A1C44CBC826DC5.4EC34DDDC547C55E6A2F860D0C351CF62E27636E&id=o-AN4EhYYtTRVcbehrgICK0l6t3isae2fVTJxuy7Mzeg2Y&dur=128.728&requiressl=yes&mm=31%2C29&pl=20&mn=sn-oxunxg8pjvn-nw2e%2Csn-bg07dnsl&ms=au%2Crdu&mt=1521852891&mv=m&ip=189.32.180.245&initcwndbps=422500&lmt=1518266680389935&ipbits=0&expire=1521874618&sparams=aitags%2Cclen%2Cdur%2Cei%2Cgir%2Cid%2Cinitcwndbps%2Cip%2Cipbits%2Citag%2Ckeepalive%2Clmt%2Cmime%2Cmm%2Cmn%2Cms%2Cmv%2Cpl%2Crequiressl%2Csource%2Cexpire&ratebypass=yes\", \"player_url\": \"/yts/jsbin/player-vflMfSEyN/en_US/base.js\", \"ext\": \"mp4\", \"height\": 360, \"format_note\": \"360p\", \"vcodec\": \"avc1.4d401e\", \"filesize\": 11475754, \"tbr\": 785.487, \"width\": 640, \"fps\": 30, \"acodec\": \"none\", \"downloader_options\": {\"http_chunk_size\": 10485760}, \"format\": \"134 - 640x360 (360p)\", \"protocol\": \"https\", \"http_headers\": {\"User-Agent\": \"Mozilla/5.0 (X11; Linux x86_64; rv:59.0) Gecko/20100101 Firefox/59.0 (Chrome)\", \"Accept-Charset\": \"ISO-8859-1,utf-8;q=0.7,*;q=0.7\", \"Accept\": \"text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\", \"Accept-Encoding\": \"gzip, deflate\", \"Accept-Language\": \"en-us,en;q=0.5\"}}, {\"format_id\": \"135\", \"url\": \"https://r1---sn-oxunxg8pjvn-nw2e.googlevideo.com/videoplayback?clen=19988924&gir=yes&keepalive=yes&aitags=133%2C134%2C135%2C136%2C137%2C160%2C242%2C243%2C244%2C247%2C248%2C278&mime=video%2Fmp4&key=yt6&itag=135&ei=WqK1WouIH4arxgS3wp34Bg&c=WEB&source=youtube&fvip=2&signature=1D69CF528D901735EBDCC69A72289C0A4DC233E3.609AAD6ED8C9FDD40AFCAF1B9EF69F9BC1FA0638&id=o-AN4EhYYtTRVcbehrgICK0l6t3isae2fVTJxuy7Mzeg2Y&dur=128.728&requiressl=yes&mm=31%2C29&pl=20&mn=sn-oxunxg8pjvn-nw2e%2Csn-bg07dnsl&ms=au%2Crdu&mt=1521852891&mv=m&ip=189.32.180.245&initcwndbps=422500&lmt=1518266700552050&ipbits=0&expire=1521874618&sparams=aitags%2Cclen%2Cdur%2Cei%2Cgir%2Cid%2Cinitcwndbps%2Cip%2Cipbits%2Citag%2Ckeepalive%2Clmt%2Cmime%2Cmm%2Cmn%2Cms%2Cmv%2Cpl%2Crequiressl%2Csource%2Cexpire&ratebypass=yes\", \"player_url\": \"/yts/jsbin/player-vflMfSEyN/en_US/base.js\", \"ext\": \"mp4\", \"height\": 480, \"format_note\": \"480p\", \"vcodec\": \"avc1.4d401f\", \"filesize\": 19988924, \"tbr\": 1431.791, \"width\": 854, \"fps\": 30, \"acodec\": \"none\", \"downloader_options\": {\"http_chunk_size\": 10485760}, \"format\": \"135 - 854x480 (480p)\", \"protocol\": \"https\", \"http_headers\": {\"User-Agent\": \"Mozilla/5.0 (X11; Linux x86_64; rv:59.0) Gecko/20100101 Firefox/59.0 (Chrome)\", \"Accept-Charset\": \"ISO-8859-1,utf-8;q=0.7,*;q=0.7\", \"Accept\": \"text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\", \"Accept-Encoding\": \"gzip, deflate\", \"Accept-Language\": \"en-us,en;q=0.5\"}}, {\"format_id\": \"247\", \"url\": \"https://r1---sn-oxunxg8pjvn-nw2e.googlevideo.com/videoplayback?clen=18507713&gir=yes&keepalive=yes&aitags=133%2C134%2C135%2C136%2C137%2C160%2C242%2C243%2C244%2C247%2C248%2C278&mime=video%2Fwebm&key=yt6&itag=247&ei=WqK1WouIH4arxgS3wp34Bg&c=WEB&source=youtube&fvip=2&signature=99095CF3DC308FA7218E8F5C41AB2C29E9ACDA6F.8C7B86BA1126E871782E70FBA1E2C91F87E80440&id=o-AN4EhYYtTRVcbehrgICK0l6t3isae2fVTJxuy7Mzeg2Y&dur=128.728&requiressl=yes&mm=31%2C29&pl=20&mn=sn-oxunxg8pjvn-nw2e%2Csn-bg07dnsl&ms=au%2Crdu&mt=1521852891&mv=m&ip=189.32.180.245&initcwndbps=422500&lmt=1518268920124790&ipbits=0&expire=1521874618&sparams=aitags%2Cclen%2Cdur%2Cei%2Cgir%2Cid%2Cinitcwndbps%2Cip%2Cipbits%2Citag%2Ckeepalive%2Clmt%2Cmime%2Cmm%2Cmn%2Cms%2Cmv%2Cpl%2Crequiressl%2Csource%2Cexpire&ratebypass=yes\", \"player_url\": \"/yts/jsbin/player-vflMfSEyN/en_US/base.js\", \"ext\": \"webm\", \"height\": 720, \"format_note\": \"720p\", \"vcodec\": \"vp9\", \"filesize\": 18507713, \"tbr\": 1519.343, \"width\": 1280, \"fps\": 30, \"acodec\": \"none\", \"downloader_options\": {\"http_chunk_size\": 10485760}, \"format\": \"247 - 1280x720 (720p)\", \"protocol\": \"https\", \"http_headers\": {\"User-Agent\": \"Mozilla/5.0 (X11; Linux x86_64; rv:59.0) Gecko/20100101 Firefox/59.0 (Chrome)\", \"Accept-Charset\": \"ISO-8859-1,utf-8;q=0.7,*;q=0.7\", \"Accept\": \"text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\", \"Accept-Encoding\": \"gzip, deflate\", \"Accept-Language\": \"en-us,en;q=0.5\"}}, {\"format_id\": \"136\", \"url\": \"https://r1---sn-oxunxg8pjvn-nw2e.googlevideo.com/videoplayback?clen=26339828&gir=yes&keepalive=yes&aitags=133%2C134%2C135%2C136%2C137%2C160%2C242%2C243%2C244%2C247%2C248%2C278&mime=video%2Fmp4&key=yt6&itag=136&ei=WqK1WouIH4arxgS3wp34Bg&c=WEB&source=youtube&fvip=2&signature=3F58457936E5537B2A727397CAF7D4F07C8E52FC.C1C77D685D5EBA22552B9BC7BDEE6891C83A6DC7&id=o-AN4EhYYtTRVcbehrgICK0l6t3isae2fVTJxuy7Mzeg2Y&dur=128.728&requiressl=yes&mm=31%2C29&pl=20&mn=sn-oxunxg8pjvn-nw2e%2Csn-bg07dnsl&ms=au%2Crdu&mt=1521852891&mv=m&ip=189.32.180.245&initcwndbps=422500&lmt=1518266747088330&ipbits=0&expire=1521874618&sparams=aitags%2Cclen%2Cdur%2Cei%2Cgir%2Cid%2Cinitcwndbps%2Cip%2Cipbits%2Citag%2Ckeepalive%2Clmt%2Cmime%2Cmm%2Cmn%2Cms%2Cmv%2Cpl%2Crequiressl%2Csource%2Cexpire&ratebypass=yes\", \"player_url\": \"/yts/jsbin/player-vflMfSEyN/en_US/base.js\", \"ext\": \"mp4\", \"height\": 720, \"format_note\": \"720p\", \"vcodec\": \"avc1.4d401f\", \"filesize\": 26339828, \"tbr\": 1941.641, \"width\": 1280, \"fps\": 30, \"acodec\": \"none\", \"downloader_options\": {\"http_chunk_size\": 10485760}, \"format\": \"136 - 1280x720 (720p)\", \"protocol\": \"https\", \"http_headers\": {\"User-Agent\": \"Mozilla/5.0 (X11; Linux x86_64; rv:59.0) Gecko/20100101 Firefox/59.0 (Chrome)\", \"Accept-Charset\": \"ISO-8859-1,utf-8;q=0.7,*;q=0.7\", \"Accept\": \"text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\", \"Accept-Encoding\": \"gzip, deflate\", \"Accept-Language\": \"en-us,en;q=0.5\"}}, {\"format_id\": \"248\", \"url\": \"https://r1---sn-oxunxg8pjvn-nw2e.googlevideo.com/videoplayback?clen=35916248&gir=yes&keepalive=yes&aitags=133%2C134%2C135%2C136%2C137%2C160%2C242%2C243%2C244%2C247%2C248%2C278&mime=video%2Fwebm&key=yt6&itag=248&ei=WqK1WouIH4arxgS3wp34Bg&c=WEB&source=youtube&fvip=2&signature=29BE61209056E6430AA921970FDC012CAE2B4BDE.7DE0DB8AE8633E3FFBBDFC669904874FC2F658B8&id=o-AN4EhYYtTRVcbehrgICK0l6t3isae2fVTJxuy7Mzeg2Y&dur=128.728&requiressl=yes&mm=31%2C29&pl=20&mn=sn-oxunxg8pjvn-nw2e%2Csn-bg07dnsl&ms=au%2Crdu&mt=1521852891&mv=m&ip=189.32.180.245&initcwndbps=422500&lmt=1518269049502884&ipbits=0&expire=1521874618&sparams=aitags%2Cclen%2Cdur%2Cei%2Cgir%2Cid%2Cinitcwndbps%2Cip%2Cipbits%2Citag%2Ckeepalive%2Clmt%2Cmime%2Cmm%2Cmn%2Cms%2Cmv%2Cpl%2Crequiressl%2Csource%2Cexpire&ratebypass=yes\", \"player_url\": \"/yts/jsbin/player-vflMfSEyN/en_US/base.js\", \"ext\": \"webm\", \"height\": 1080, \"format_note\": \"1080p\", \"vcodec\": \"vp9\", \"filesize\": 35916248, \"tbr\": 2622.603, \"width\": 1920, \"fps\": 30, \"acodec\": \"none\", \"downloader_options\": {\"http_chunk_size\": 10485760}, \"format\": \"248 - 1920x1080 (1080p)\", \"protocol\": \"https\", \"http_headers\": {\"User-Agent\": \"Mozilla/5.0 (X11; Linux x86_64; rv:59.0) Gecko/20100101 Firefox/59.0 (Chrome)\", \"Accept-Charset\": \"ISO-8859-1,utf-8;q=0.7,*;q=0.7\", \"Accept\": \"text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\", \"Accept-Encoding\": \"gzip, deflate\", \"Accept-Language\": \"en-us,en;q=0.5\"}}, {\"format_id\": \"137\", \"url\": \"https://r1---sn-oxunxg8pjvn-nw2e.googlevideo.com/videoplayback?clen=37691953&gir=yes&keepalive=yes&aitags=133%2C134%2C135%2C136%2C137%2C160%2C242%2C243%2C244%2C247%2C248%2C278&mime=video%2Fmp4&key=yt6&itag=137&ei=WqK1WouIH4arxgS3wp34Bg&c=WEB&source=youtube&fvip=2&signature=388778FCA04F5212CADE5E5843BF9A51E7DA7600.236E871B1F5F24F10E96CCFB1C0072BBF6984A09&id=o-AN4EhYYtTRVcbehrgICK0l6t3isae2fVTJxuy7Mzeg2Y&dur=128.728&requiressl=yes&mm=31%2C29&pl=20&mn=sn-oxunxg8pjvn-nw2e%2Csn-bg07dnsl&ms=au%2Crdu&mt=1521852891&mv=m&ip=189.32.180.245&initcwndbps=422500&lmt=1518266766621383&ipbits=0&expire=1521874618&sparams=aitags%2Cclen%2Cdur%2Cei%2Cgir%2Cid%2Cinitcwndbps%2Cip%2Cipbits%2Citag%2Ckeepalive%2Clmt%2Cmime%2Cmm%2Cmn%2Cms%2Cmv%2Cpl%2Crequiressl%2Csource%2Cexpire&ratebypass=yes\", \"player_url\": \"/yts/jsbin/player-vflMfSEyN/en_US/base.js\", \"ext\": \"mp4\", \"height\": 1080, \"format_note\": \"1080p\", \"vcodec\": \"avc1.640028\", \"filesize\": 37691953, \"tbr\": 3054.153, \"width\": 1920, \"fps\": 30, \"acodec\": \"none\", \"downloader_options\": {\"http_chunk_size\": 10485760}, \"format\": \"137 - 1920x1080 (1080p)\", \"protocol\": \"https\", \"http_headers\": {\"User-Agent\": \"Mozilla/5.0 (X11; Linux x86_64; rv:59.0) Gecko/20100101 Firefox/59.0 (Chrome)\", \"Accept-Charset\": \"ISO-8859-1,utf-8;q=0.7,*;q=0.7\", \"Accept\": \"text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\", \"Accept-Encoding\": \"gzip, deflate\", \"Accept-Language\": \"en-us,en;q=0.5\"}}, {\"format_id\": \"17\", \"url\": \"https://r1---sn-oxunxg8pjvn-nw2e.googlevideo.com/videoplayback?clen=1305912&gir=yes&mime=video%2F3gpp&key=yt6&signature=DF34BA07F6E1E2FC14A209812F99547739C2F83A.A3AFA3B1FE5AFA482631C13A5776EAEA2F3AF485&ei=WqK1WouIH4arxgS3wp34Bg&c=WEB&source=youtube&fvip=2&itag=17&id=o-AN4EhYYtTRVcbehrgICK0l6t3isae2fVTJxuy7Mzeg2Y&dur=128.824&requiressl=yes&mm=31%2C29&pl=20&mn=sn-oxunxg8pjvn-nw2e%2Csn-bg07dnsl&ms=au%2Crdu&mt=1521852891&mv=m&ip=189.32.180.245&initcwndbps=422500&lmt=1518265873424483&ipbits=0&expire=1521874618&sparams=clen%2Cdur%2Cei%2Cgir%2Cid%2Cinitcwndbps%2Cip%2Cipbits%2Citag%2Clmt%2Cmime%2Cmm%2Cmn%2Cms%2Cmv%2Cpl%2Crequiressl%2Csource%2Cexpire&ratebypass=yes\", \"player_url\": \"/yts/jsbin/player-vflMfSEyN/en_US/base.js\", \"ext\": \"3gp\", \"width\": 176, \"height\": 144, \"acodec\": \"mp4a.40.2\", \"abr\": 24, \"vcodec\": \"mp4v.20.3\", \"resolution\": \"176x144\", \"format_note\": \"small\", \"format\": \"17 - 176x144 (small)\", \"protocol\": \"https\", \"http_headers\": {\"User-Agent\": \"Mozilla/5.0 (X11; Linux x86_64; rv:59.0) Gecko/20100101 Firefox/59.0 (Chrome)\", \"Accept-Charset\": \"ISO-8859-1,utf-8;q=0.7,*;q=0.7\", \"Accept\": \"text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\", \"Accept-Encoding\": \"gzip, deflate\", \"Accept-Language\": \"en-us,en;q=0.5\"}}, {\"format_id\": \"36\", \"url\": \"https://r1---sn-oxunxg8pjvn-nw2e.googlevideo.com/videoplayback?clen=3611941&gir=yes&mime=video%2F3gpp&key=yt6&signature=6614F4F7826446700A1428C53DD99A4FF64D340E.3C38D22E95CF3FD8CB1B35981230D90460E2CF11&ei=WqK1WouIH4arxgS3wp34Bg&c=WEB&source=youtube&fvip=2&itag=36&id=o-AN4EhYYtTRVcbehrgICK0l6t3isae2fVTJxuy7Mzeg2Y&dur=128.824&requiressl=yes&mm=31%2C29&pl=20&mn=sn-oxunxg8pjvn-nw2e%2Csn-bg07dnsl&ms=au%2Crdu&mt=1521852891&mv=m&ip=189.32.180.245&initcwndbps=422500&lmt=1518265890615475&ipbits=0&expire=1521874618&sparams=clen%2Cdur%2Cei%2Cgir%2Cid%2Cinitcwndbps%2Cip%2Cipbits%2Citag%2Clmt%2Cmime%2Cmm%2Cmn%2Cms%2Cmv%2Cpl%2Crequiressl%2Csource%2Cexpire&ratebypass=yes\", \"player_url\": \"/yts/jsbin/player-vflMfSEyN/en_US/base.js\", \"ext\": \"3gp\", \"width\": 320, \"acodec\": \"mp4a.40.2\", \"vcodec\": \"mp4v.20.3\", \"resolution\": \"320x180\", \"height\": 180, \"format_note\": \"small\", \"format\": \"36 - 320x180 (small)\", \"protocol\": \"https\", \"http_headers\": {\"User-Agent\": \"Mozilla/5.0 (X11; Linux x86_64; rv:59.0) Gecko/20100101 Firefox/59.0 (Chrome)\", \"Accept-Charset\": \"ISO-8859-1,utf-8;q=0.7,*;q=0.7\", \"Accept\": \"text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\", \"Accept-Encoding\": \"gzip, deflate\", \"Accept-Language\": \"en-us,en;q=0.5\"}}, {\"format_id\": \"18\", \"url\": \"https://r1---sn-oxunxg8pjvn-nw2e.googlevideo.com/videoplayback?clen=11652498&gir=yes&mime=video%2Fmp4&key=yt6&signature=841D8458AE82A3C5F5F8DE285DFC88BF6CB1ED98.70C28F6729F1ADB47B86B8E0E3DCD4D5A8D5B93E&ei=WqK1WouIH4arxgS3wp34Bg&c=WEB&source=youtube&ratebypass=yes&fvip=2&itag=18&id=o-AN4EhYYtTRVcbehrgICK0l6t3isae2fVTJxuy7Mzeg2Y&dur=128.777&requiressl=yes&mm=31%2C29&pl=20&mn=sn-oxunxg8pjvn-nw2e%2Csn-bg07dnsl&ms=au%2Crdu&mt=1521852891&mv=m&ip=189.32.180.245&initcwndbps=422500&lmt=1518265910630758&ipbits=0&expire=1521874618&sparams=clen%2Cdur%2Cei%2Cgir%2Cid%2Cinitcwndbps%2Cip%2Cipbits%2Citag%2Clmt%2Cmime%2Cmm%2Cmn%2Cms%2Cmv%2Cpl%2Cratebypass%2Crequiressl%2Csource%2Cexpire\", \"player_url\": \"/yts/jsbin/player-vflMfSEyN/en_US/base.js\", \"ext\": \"mp4\", \"width\": 640, \"height\": 360, \"acodec\": \"mp4a.40.2\", \"abr\": 96, \"vcodec\": \"avc1.42001E\", \"resolution\": \"640x360\", \"format_note\": \"medium\", \"format\": \"18 - 640x360 (medium)\", \"protocol\": \"https\", \"http_headers\": {\"User-Agent\": \"Mozilla/5.0 (X11; Linux x86_64; rv:59.0) Gecko/20100101 Firefox/59.0 (Chrome)\", \"Accept-Charset\": \"ISO-8859-1,utf-8;q=0.7,*;q=0.7\", \"Accept\": \"text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\", \"Accept-Encoding\": \"gzip, deflate\", \"Accept-Language\": \"en-us,en;q=0.5\"}}, {\"format_id\": \"43\", \"url\": \"https://r1---sn-oxunxg8pjvn-nw2e.googlevideo.com/videoplayback?clen=14527943&gir=yes&mime=video%2Fwebm&key=yt6&signature=05D17821F8EF4DE37169976CF6518E3D61DBE011.205CD0BA04CCF692E208491CBEAFFF5AF1FAE435&ei=WqK1WouIH4arxgS3wp34Bg&c=WEB&source=youtube&ratebypass=yes&fvip=2&itag=43&id=o-AN4EhYYtTRVcbehrgICK0l6t3isae2fVTJxuy7Mzeg2Y&dur=0.000&requiressl=yes&mm=31%2C29&pl=20&mn=sn-oxunxg8pjvn-nw2e%2Csn-bg07dnsl&ms=au%2Crdu&mt=1521852891&mv=m&ip=189.32.180.245&initcwndbps=422500&lmt=1518268001462879&ipbits=0&expire=1521874618&sparams=clen%2Cdur%2Cei%2Cgir%2Cid%2Cinitcwndbps%2Cip%2Cipbits%2Citag%2Clmt%2Cmime%2Cmm%2Cmn%2Cms%2Cmv%2Cpl%2Cratebypass%2Crequiressl%2Csource%2Cexpire\", \"player_url\": \"/yts/jsbin/player-vflMfSEyN/en_US/base.js\", \"ext\": \"webm\", \"width\": 640, \"height\": 360, \"acodec\": \"vorbis\", \"abr\": 128, \"vcodec\": \"vp8.0\", \"resolution\": \"640x360\", \"format_note\": \"medium\", \"format\": \"43 - 640x360 (medium)\", \"protocol\": \"https\", \"http_headers\": {\"User-Agent\": \"Mozilla/5.0 (X11; Linux x86_64; rv:59.0) Gecko/20100101 Firefox/59.0 (Chrome)\", \"Accept-Charset\": \"ISO-8859-1,utf-8;q=0.7,*;q=0.7\", \"Accept\": \"text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\", \"Accept-Encoding\": \"gzip, deflate\", \"Accept-Language\": \"en-us,en;q=0.5\"}}, {\"format_id\": \"22\", \"url\": \"https://r1---sn-oxunxg8pjvn-nw2e.googlevideo.com/videoplayback?id=o-AN4EhYYtTRVcbehrgICK0l6t3isae2fVTJxuy7Mzeg2Y&dur=128.777&requiressl=yes&mm=31%2C29&pl=20&mn=sn-oxunxg8pjvn-nw2e%2Csn-bg07dnsl&fvip=2&ms=au%2Crdu&mt=1521852891&mv=m&mime=video%2Fmp4&ip=189.32.180.245&key=yt6&signature=2E86F5564E07DD67C352E7FE162B5DF4DC020045.C70C523DB224B0CE1FFFF26C3C00F22BA36E7AF3&ei=WqK1WouIH4arxgS3wp34Bg&initcwndbps=422500&c=WEB&lmt=1518266798247993&source=youtube&ratebypass=yes&ipbits=0&expire=1521874618&sparams=dur%2Cei%2Cid%2Cinitcwndbps%2Cip%2Cipbits%2Citag%2Clmt%2Cmime%2Cmm%2Cmn%2Cms%2Cmv%2Cpl%2Cratebypass%2Crequiressl%2Csource%2Cexpire&itag=22\", \"player_url\": \"/yts/jsbin/player-vflMfSEyN/en_US/base.js\", \"ext\": \"mp4\", \"width\": 1280, \"height\": 720, \"acodec\": \"mp4a.40.2\", \"abr\": 192, \"vcodec\": \"avc1.64001F\", \"resolution\": \"1280x720\", \"format_note\": \"hd720\", \"format\": \"22 - 1280x720 (hd720)\", \"protocol\": \"https\", \"http_headers\": {\"User-Agent\": \"Mozilla/5.0 (X11; Linux x86_64; rv:59.0) Gecko/20100101 Firefox/59.0 (Chrome)\", \"Accept-Charset\": \"ISO-8859-1,utf-8;q=0.7,*;q=0.7\", \"Accept\": \"text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\", \"Accept-Encoding\": \"gzip, deflate\", \"Accept-Language\": \"en-us,en;q=0.5\"}}], \"is_live\": null, \"start_time\": null, \"end_time\": null, \"series\": null, \"season_number\": null, \"episode_number\": null, \"extractor\": \"youtube\", \"webpage_url_basename\": \"NE71LIUJ4gQ\", \"extractor_key\": \"Youtube\", \"playlist\": null, \"playlist_index\": null, \"thumbnails\": [{\"url\": \"https://i.ytimg.com/vi/NE71LIUJ4gQ/maxresdefault.jpg\", \"id\": \"0\"}], \"display_id\": \"NE71LIUJ4gQ\", \"requested_subtitles\": null, \"requested_formats\": [{\"format_id\": \"137\", \"url\": \"https://r1---sn-oxunxg8pjvn-nw2e.googlevideo.com/videoplayback?clen=37691953&gir=yes&keepalive=yes&aitags=133%2C134%2C135%2C136%2C137%2C160%2C242%2C243%2C244%2C247%2C248%2C278&mime=video%2Fmp4&key=yt6&itag=137&ei=WqK1WouIH4arxgS3wp34Bg&c=WEB&source=youtube&fvip=2&signature=388778FCA04F5212CADE5E5843BF9A51E7DA7600.236E871B1F5F24F10E96CCFB1C0072BBF6984A09&id=o-AN4EhYYtTRVcbehrgICK0l6t3isae2fVTJxuy7Mzeg2Y&dur=128.728&requiressl=yes&mm=31%2C29&pl=20&mn=sn-oxunxg8pjvn-nw2e%2Csn-bg07dnsl&ms=au%2Crdu&mt=1521852891&mv=m&ip=189.32.180.245&initcwndbps=422500&lmt=1518266766621383&ipbits=0&expire=1521874618&sparams=aitags%2Cclen%2Cdur%2Cei%2Cgir%2Cid%2Cinitcwndbps%2Cip%2Cipbits%2Citag%2Ckeepalive%2Clmt%2Cmime%2Cmm%2Cmn%2Cms%2Cmv%2Cpl%2Crequiressl%2Csource%2Cexpire&ratebypass=yes\", \"player_url\": \"/yts/jsbin/player-vflMfSEyN/en_US/base.js\", \"ext\": \"mp4\", \"height\": 1080, \"format_note\": \"1080p\", \"vcodec\": \"avc1.640028\", \"filesize\": 37691953, \"tbr\": 3054.153, \"width\": 1920, \"fps\": 30, \"acodec\": \"none\", \"downloader_options\": {\"http_chunk_size\": 10485760}, \"format\": \"137 - 1920x1080 (1080p)\", \"protocol\": \"https\", \"http_headers\": {\"User-Agent\": \"Mozilla/5.0 (X11; Linux x86_64; rv:59.0) Gecko/20100101 Firefox/59.0 (Chrome)\", \"Accept-Charset\": \"ISO-8859-1,utf-8;q=0.7,*;q=0.7\", \"Accept\": \"text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\", \"Accept-Encoding\": \"gzip, deflate\", \"Accept-Language\": \"en-us,en;q=0.5\"}}, {\"format_id\": \"171\", \"url\": \"https://r1---sn-oxunxg8pjvn-nw2e.googlevideo.com/videoplayback?clen=2294642&gir=yes&keepalive=yes&mime=audio%2Fwebm&key=yt6&signature=0169F87DB2F53BE1A2439441B463FD7B31086B40.653300D24CA3181E3222325EB0087741648C77A0&ei=WqK1WouIH4arxgS3wp34Bg&c=WEB&source=youtube&fvip=2&itag=171&id=o-AN4EhYYtTRVcbehrgICK0l6t3isae2fVTJxuy7Mzeg2Y&dur=128.732&requiressl=yes&mm=31%2C29&pl=20&mn=sn-oxunxg8pjvn-nw2e%2Csn-bg07dnsl&ms=au%2Crdu&mt=1521852891&mv=m&ip=189.32.180.245&initcwndbps=422500&lmt=1518268125409651&ipbits=0&expire=1521874618&sparams=clen%2Cdur%2Cei%2Cgir%2Cid%2Cinitcwndbps%2Cip%2Cipbits%2Citag%2Ckeepalive%2Clmt%2Cmime%2Cmm%2Cmn%2Cms%2Cmv%2Cpl%2Crequiressl%2Csource%2Cexpire&ratebypass=yes\", \"player_url\": \"/yts/jsbin/player-vflMfSEyN/en_US/base.js\", \"ext\": \"webm\", \"acodec\": \"vorbis\", \"format_note\": \"DASH audio\", \"abr\": 128, \"filesize\": 2294642, \"tbr\": 159.289, \"vcodec\": \"none\", \"downloader_options\": {\"http_chunk_size\": 10485760}, \"format\": \"171 - audio only (DASH audio)\", \"protocol\": \"https\", \"http_headers\": {\"User-Agent\": \"Mozilla/5.0 (X11; Linux x86_64; rv:59.0) Gecko/20100101 Firefox/59.0 (Chrome)\", \"Accept-Charset\": \"ISO-8859-1,utf-8;q=0.7,*;q=0.7\", \"Accept\": \"text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\", \"Accept-Encoding\": \"gzip, deflate\", \"Accept-Language\": \"en-us,en;q=0.5\"}}], \"format\": \"137 - 1920x1080 (1080p)+171 - audio only (DASH audio)\", \"format_id\": \"137+171\", \"width\": 1920, \"height\": 1080, \"resolution\": null, \"fps\": 30, \"vcodec\": \"avc1.640028\", \"vbr\": null, \"stretched_ratio\": null, \"acodec\": \"vorbis\", \"abr\": 128, \"ext\": \"mp4\"}")
