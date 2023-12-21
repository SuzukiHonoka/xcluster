package server

import "xcluster/internal/database"

type Groups []*Group

func (gs Groups) Delete() error {
	var err error
	// delete group in sequence
	for _, g := range gs {
		// NOTE: DO NOT INLINE THE DELETE OPERATION SINCE MULTI-TABLE ASSOCIATED
		if err = g.Delete(); err != nil {
			return err
		}
	}
	return nil
}

func (gs Groups) HasGroupID(gid GroupID) bool {
	for _, g := range gs {
		if g.ID == gid {
			return true
		}
	}
	return false
}

func (gs Groups) GetServers() (Servers, error) {
	// fill up ids
	groupIDs := make([]GroupID, 0, len(gs))
	for _, g := range gs {
		groupIDs = append(groupIDs, g.ID)
	}
	// actual query, inline query for performance
	var servers Servers
	if err := database.DB.Find(&servers, "server_group_id IN ?", groupIDs).Error; err != nil {
		return nil, err
	}
	return servers, nil
}
