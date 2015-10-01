package main

type EmbInfo struct {
    Temp        float32
    Humidity    float32
}

func ReadValues() EmbInfo {
    current := EmbInfo {
        Temp: 20.5,
        Humidity: 26,
    }

    return current
}
