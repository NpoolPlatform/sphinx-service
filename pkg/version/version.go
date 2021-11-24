package version

import (
	npool "github.com/NpoolPlatform/message/npool/trading" //nolint

	cv "github.com/NpoolPlatform/go-service-framework/pkg/version"

	"golang.org/x/xerrors"
)

func Version() (*npool.VersionResponse, error) {
	info, err := cv.GetVersion()
	if err != nil {
		return nil, xerrors.Errorf("get service version error: %w", err)
	}
	return &npool.VersionResponse{
		Info: info + " 1124-2252",
	}, nil
}
