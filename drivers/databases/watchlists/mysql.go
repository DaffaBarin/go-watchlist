package watchlists

import (
	"go-watchlist/business/watchlists"
	"go-watchlist/drivers/databases/medias"

	"gorm.io/gorm"
)

type MysqlWatchlistRepository struct {
	Conn *gorm.DB
}

func NewMysqlWatchlistRepository(conn *gorm.DB) watchlists.Repository {
	return &MysqlWatchlistRepository{
		Conn: conn,
	}
}

func (rep *MysqlWatchlistRepository) Create(userID int, domain *watchlists.Domain) (watchlists.Domain, error) {
	watchlist := fromDomain(*domain)
	watchlist.UserID = userID
	result := rep.Conn.Create(&watchlist)

	if result.Error != nil {
		return watchlists.Domain{}, result.Error
	}

	return toDomain(watchlist), nil
}

func (rep *MysqlWatchlistRepository) GetAllByUserID(userID int) ([]watchlists.Domain, error) {
	var watchlist []Watchlists
	var unwatchedList []UnwatchedPivot
	var watchedList []WatchedPivot

	result := rep.Conn.Find(&watchlist, "user_id = ?", userID)
	for i := range watchlist {
		unwatched := rep.Conn.Find(&unwatchedList, "watchlists_id = ?", watchlist[i].ID)
		if unwatched.Error != nil {
			return nil, unwatched.Error
		}
		for j := range unwatchedList {
			mediaList := medias.Medias{}
			medID := unwatchedList[j].Medias_id
			res := rep.Conn.Where("id = ?", medID).Take(&mediaList)
			watchlist[i].Unwatched = append(watchlist[i].Unwatched,
				watchlists.MediaStruct{ID: mediaList.ID, Type: mediaList.Type, Name: mediaList.Title, Overview: mediaList.Overview})
			if res.Error != nil {
				return nil, res.Error
			}
		}
		watched := rep.Conn.Find(&watchedList, "watchlists_id = ?", watchlist[i].ID)
		if watched.Error != nil {
			return nil, watched.Error
		}
		for k := range watchedList {
			mediaList := medias.Medias{}
			medID := watchedList[k].Medias_id
			res := rep.Conn.Where("id = ?", medID).Take(&mediaList)
			watchlist[i].Watched = append(watchlist[i].Watched,
				watchlists.MediaStruct{ID: mediaList.ID, Type: mediaList.Type, Name: mediaList.Title, Overview: mediaList.Overview})
			if res.Error != nil {
				return nil, res.Error
			}
		}

	}
	if result.Error != nil {
		return nil, result.Error
	}

	return toListDomain(watchlist), nil
}

func (rep *MysqlWatchlistRepository) GetByID(userID int, id int) (watchlists.Domain, error) {
	var watchlist Watchlists
	var unwatchedList []UnwatchedPivot
	var watchedList []WatchedPivot

	result := rep.Conn.Where("user_id = ?", userID).Where("id = ?", id).First(&watchlist)

	if result.Error != nil {
		return watchlists.Domain{}, result.Error
	}
	unwatched := rep.Conn.Find(&unwatchedList, "watchlists_id = ?", watchlist.ID)
	if unwatched.Error != nil {
		return watchlists.Domain{}, unwatched.Error
	}
	for j := range unwatchedList {
		mediaList := medias.Medias{}
		medID := unwatchedList[j].Medias_id
		res := rep.Conn.Where("id = ?", medID).Take(&mediaList)
		watchlist.Unwatched = append(watchlist.Unwatched,
			watchlists.MediaStruct{ID: mediaList.ID, Type: mediaList.Type, Name: mediaList.Title, Overview: mediaList.Overview})
		if res.Error != nil {
			return watchlists.Domain{}, res.Error
		}
	}
	watched := rep.Conn.Find(&watchedList, "watchlists_id = ?", watchlist.ID)
	if watched.Error != nil {
		return watchlists.Domain{}, watched.Error
	}
	for k := range watchedList {
		mediaList := medias.Medias{}
		medID := watchedList[k].Medias_id
		res := rep.Conn.Where("id = ?", medID).Take(&mediaList)
		watchlist.Watched = append(watchlist.Watched,
			watchlists.MediaStruct{ID: mediaList.ID, Type: mediaList.Type, Name: mediaList.Title, Overview: mediaList.Overview})
		if res.Error != nil {
			return watchlists.Domain{}, res.Error
		}
	}

	return toDomain(watchlist), nil
}

func (rep *MysqlWatchlistRepository) InsertMedia(id int, MediaID int) (watchlists.Domain, error) {
	var watchlist Watchlists
	result := rep.Conn.Create(&UnwatchedPivot{Watchlists_id: id, Medias_id: MediaID})
	if result.Error != nil {
		return watchlists.Domain{}, result.Error
	}

	return toDomain(watchlist), nil
}

func (rep *MysqlWatchlistRepository) UpdateMedia(userID int, watchlistID int, mediaID int) (watchlists.Domain, error) {
	var watchlist Watchlists
	var unwatched UnwatchedPivot
	result := rep.Conn.Where("user_id = ?", userID).Where("id = ?", watchlistID).First(&watchlist)
	if result.Error != nil {
		return watchlists.Domain{}, result.Error
	}
	create := rep.Conn.Create(&WatchedPivot{Watchlists_id: watchlistID, Medias_id: mediaID})
	if create.Error != nil {
		return watchlists.Domain{}, create.Error
	}
	delete := rep.Conn.Where("watchlists_id", watchlistID).Where("medias_id", mediaID).Delete(&unwatched)
	if delete.Error != nil {
		return watchlists.Domain{}, create.Error
	}
	return toDomain(watchlist), nil
}
