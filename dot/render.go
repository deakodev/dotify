package dot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

type payload struct {
	Graph  string `json:"graph"`
	Layout string `json:"layout"`
	Format string `json:"format"`
}

func Render(graph *Graph, outPath, format string) error {
	payload := payload{graph.String(), "dot", format}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	res, err := http.Post("https://quickchart.io/graphviz", "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return fmt.Errorf("graphviz post failed: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("graphviz bad response: %s", res.Status)
	}

	fileName := fmt.Sprintf("%s.%s", graph.Name(), payload.Format)
	filePath := path.Join(outPath, fileName)
	outFile, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, res.Body)
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	fmt.Printf("Rendered: %s\n", fileName)
	return nil
}
