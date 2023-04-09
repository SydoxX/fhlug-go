# fhlug - Go Workshop

Wer öfter mit Cloud- und Containerumgebungen zu tun hat, ist sicherlich schon einmal über die Programmiersprache [Go](https://go.dev/) gestolpert. Dort wird sie verwendet, um vorallem Microservices und Datenbanken zu implementieren. Doch warum wird Go dort so extensiv benutzt, wo kann man Go sonst noch einsetzen und was sind die Vorteile gegenüber anderer Programmiersprachen? Diese Fragen werden in diesem Workshop beantwortet.

Dieser Workshop ist aufgeteilt in zwei Teile: zuerst wird Go erläutert und in der Theorie erklärt, darauffolgend wird eine kurze Codingsession abgehalten, um Go auch in der Praxis zu erfahren.

Der erste Teil behandelt die Grundlagen von Go wie dessen Syntax, aber auch gewisse Eigenheiten der Sprache wie das Ducktyping und die in Go 1.18 dazugekommene Generizität. Im zweiten, praktischen Teil des Workshops werden zwei Programme implementiert, davon eine anlehnend an den vorhergehenden [Rust Workshop](https://fhlug.at/2023/02/rust-workshop/).

* Eine CLI-Anwendung, die JSON/YAML-Dateien verarbeitet.
* Eine Fortune-as-a-Service Anwendung, welche uns kluge Sprüche übers Weg zurückliefert.

Für den praktischen Teil wird der Go Compiler benötigt; endweder [lokal am PC installieren](https://go.dev/doc/install) oder über einen Docker Container (`docker pull golang:1.20.3`) ausführen.

Die Präsentation sowie der erstellte Code können auch auf Github eingesehen werden: https://github.com/SydoxX/fhlug-go
