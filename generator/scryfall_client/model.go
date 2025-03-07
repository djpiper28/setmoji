package scryfallclient

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
)

const (
	UserAgentHeader = "github.com_djpiper28_setmoji/1.0"
	AcceptHeader    = "*/*"
	BaseUrl         = "https://api.scryfall.com"
	SetsUrl         = BaseUrl + "/sets"
)

type Set struct {
	Object        string    `json:"object"`
	Id            uuid.UUID `json:"id"`
	Code          string    `json:"code"`
	MtgoCode      string    `json:"mtg_code"`
	ArenaCode     string    `json:"arena_code"`
	TcgPlayerId   int       `json:"tcgplayer_id"`
	Name          string    `json:"name"`
	SetType       string    `json:"set_type"`
	ReleasedAt    string    `json:"released_at"`
	BlockCode     string    `json:"block_code"`
	Block         string    `json:"block"`
	ParentSetCode string    `json:"parent_set_code"`
	CardCount     int       `json:"card_count"`
	PrintedSize   int       `json:"printed_size"`
	Digital       bool      `json:"digital"`
	FoilOnly      bool      `json:"foil_only"`
	NonFoilOnly   bool      `json:"nonfoil_only"`
	ScryfallUri   string    `json:"scryfall_uri"`
	Uri           string    `json:"uri"`
	IconSvgUri    string    `json:"icon_svg_uri"`
	SearchUri     string    `json:"search_uri"`
}

type setsResp struct {
	Data []Set `json:"data"`
}

func GetSets() ([]Set, error) {
	client := http.Client{Timeout: time.Second * 10}

	req, err := http.NewRequest("GET", SetsUrl, nil)
	if err != nil {
		return nil, errors.Join(errors.New("Cannot create request"), err)
	}

	req.Header.Set("Accept", AcceptHeader)
	req.Header.Set("User-Agent", UserAgentHeader)

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Join(errors.New("Cannot get sets"), err)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Join(errors.New("Cannot read sets"), err)
	}

	var sets setsResp
	err = json.Unmarshal(bytes, &sets)
	if err != nil {
		return nil, errors.Join(errors.New("Cannot parse sets"), err)
	}

	return sets.Data, nil
}

func (s *Set) GetSvg() ([]byte, error) {
	client := http.Client{Timeout: time.Second * 10}

	req, err := http.NewRequest("GET", s.IconSvgUri, nil)
	if err != nil {
		return nil, errors.Join(errors.New("Cannot create request"), err)
	}

	req.Header.Set("Accept", AcceptHeader)
	req.Header.Set("User-Agent", UserAgentHeader)

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Join(errors.New("Cannot get set SVG"), err)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Join(errors.New("Cannot read set SVG"), err)
	}

	return bytes, nil
}
