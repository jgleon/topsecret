# Top secret api

Api desarrollada en Go 1.15.2 que permite encontrar una posición X,Y de acuerdo a las distancias de 3 satelites, además permite determinar el mensaje enviado por los satelites.

## Problema

1. Se debe determinar la posición (x,y) de la nave y el mensaje enviado por los satelites, se conoce la ubicación (x,y) y la distancia entre cada satelite y la nave, el mensaje llega fraccionado por por satelite.

## Solución teórica

Para encontrar la ubicación x, y se utilizó el algoritmo de trilateración.

![trilateración](resource/trilateration-diagram.png?raw=true)