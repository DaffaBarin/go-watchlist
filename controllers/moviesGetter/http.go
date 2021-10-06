package moviesGetter

import (
	request "go-watchlist/controllers/medias/request"
	"net/http"
	"strconv"
	"time"

	tmdb "github.com/cyruzin/golang-tmdb"
)

func TransformMovie(req request.MediasCreate) (request.Medias, error) {
	apiKey := "599c657837f4bddd29290c19f7940640"
	tmdbClient, err := tmdb.Init(apiKey)
	if err != nil {
		return request.Medias{}, err
	}
	customClient := http.Client{
		Timeout: time.Second * 5,
		Transport: &http.Transport{
			MaxIdleConns:    10,
			IdleConnTimeout: 15 * time.Second,
		},
	}
	tmdbClient.SetClientConfig(customClient)
	tmdbClient.SetSessionID(apiKey)
	tmdbClient.SetClientAutoRetry()
	mediaIDint, _ := strconv.Atoi(req.Media_Id)
	media, err := tmdbClient.GetMovieDetails(mediaIDint, nil)
	if err != nil {
		return request.Medias{}, err
	}
	var genres string
	if len(media.Genres) > 1 {
		for i := range media.Genres {
			genres = genres + "," + media.Genres[i].Name
		}
		genres = genres[1:]
	} else if len(media.Genres) == 1 {
		genres = media.Genres[0].Name
	} else {
		genres = ""
	}
	movie := request.Medias{
		ID:                int(media.ID),
		Type:              "Movie",
		Title:             media.Title,
		Poster:            media.PosterPath,
		Overview:          media.Overview,
		Original_Language: media.OriginalLanguage,
		Genres:            genres,
		Year:              media.ReleaseDate[:4],
		Duration:          media.Runtime,
	}
	return movie, nil
}

func TransformTV(req request.MediasCreate) (request.Medias, error) {
	apiKey := "599c657837f4bddd29290c19f7940640"
	tmdbClient, err := tmdb.Init(apiKey)
	if err != nil {
		return request.Medias{}, err
	}
	customClient := http.Client{
		Timeout: time.Second * 5,
		Transport: &http.Transport{
			MaxIdleConns:    10,
			IdleConnTimeout: 15 * time.Second,
		},
	}
	tmdbClient.SetClientConfig(customClient)
	tmdbClient.SetSessionID(apiKey)
	tmdbClient.SetClientAutoRetry()
	mediaIDint, _ := strconv.Atoi(req.Media_Id)

	media, err := tmdbClient.GetTVDetails(mediaIDint, nil)
	if err != nil {
		return request.Medias{}, err
	}
	var genres string
	if len(media.Genres) > 1 {
		for i := range media.Genres {
			genres = genres + "," + media.Genres[i].Name
		}
		genres = genres[1:]
	} else if len(media.Genres) == 1 {
		genres = media.Genres[0].Name
	} else {
		genres = ""
	}
	TV := request.Medias{
		ID:                int(media.ID),
		Type:              "TV",
		Title:             media.OriginalName,
		Poster:            media.PosterPath,
		Overview:          media.Overview,
		Original_Language: media.OriginalLanguage,
		Genres:            genres,
		Year:              media.FirstAirDate[:4],
		Duration:          media.EpisodeRunTime[0],
	}
	return TV, nil
}
