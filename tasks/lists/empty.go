package lists

/* Dies ist keine Aufgabem sondern eine Hilfsfunktion. */

// Empty ist eine Hilfsfunktion, die prüft, ob eine Liste leer ist.
// Sie können diese Funktion bei den Aufgaben verwenden.
// Anmerkungen zu dieser Funktion finden Sie unten.
func Empty(list []int) bool {
	isemtpy := false
	checkpanic := func() {
		if r := recover(); r != nil {
			isemtpy = true
		}
	}
	checkempty := func() {
		defer checkpanic()
		_ = list[0]
	}
	checkempty()
	return isemtpy
}

/* Anmerkung zur Funktionsweise von Empty:

Man könnte auch einfach len(list) == 0 schreiben, aber da
dies in den Aufgaben tw. explizit verboten ist, verwenden
wir stattdessen die "brachiale" Methode.

Idee zur Vorgehensweise:
Wir prüfen, ob die Liste leer ist, indem wir auf das erste Element zugreifen.
Falls die Liste leer ist, wird ein Laufzeitfehler ausgelöst.
Diesen fangen wir mit recover ab und geben true zurück.
Falls die Liste nicht leer ist, wird kein Fehler ausgelöst
und wir geben false zurück.

Umsetzung:
Die obige Funktion setzt zuerst isemtpy auf false.

Dann werden zwei Funktionen definiert:
- 'checkpanic' prüft mit 'recover', ob ein Laufzeitfehler aufgetreten ist.
  Falls ja, wird isemtpy auf true gesetzt.
- 'checkempty' greift auf das erste Element der Liste zu.
  Vorher wird 'checkpanic' mittels 'defer' registriert.
  Das bedeutet, dass 'checkpanic' erst ausgeführt wird, wenn 'checkempty' beendet wird.
  Dadurch prüft 'checkpanic', ob bei Ausführung von 'checkempty' ein Laufzeitfehler aufgetreten ist.
  Das ist genau dann der Fall, wenn die Liste leer ist, weil dann der Zugriff auf das erste Element
  scheitert.

Nun wird 'checkempty' ausgeführt, was bei leerer Liste dazu führt, dass isemtpy auf true gesetzt wird.
Anschließend wird der Wert von isemtpy zurückgegeben.

Bemerkung:
Diese Umsetzung ist wesentlich komplizierter als einfach len(list) == 0 zu verwenden.
Sie ist auch weniger effizient, weil sie einen Laufzeitfehler auslöst.
Dessen Behandlung ist aufwändig und verlangsamt das Programm.
Dieses Beispiel dient nur dazu, die Verwendung von recover zu demonstrieren
und gleichzeitig zu zeigen, dass solche Anforderungen wie "benutzen Sie NICHT die len-Funktion"
durchaus umsetzbar sind.
*/
