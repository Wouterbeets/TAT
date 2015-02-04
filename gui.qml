import QtQuick 2.0
import QtQuick.Controls 1.3
import GoExtensions 1.0

Rectangle {
	id : window
	width: 800; height: 600
	color: "white"
	Rectangle {
		color: "grey"
		width: window.width
		height: 100
		Text {
			text : 	"Thuis aan Tafel"
			anchors.centerIn: parent
		}
		Button {
			id: "test"
			text: "button"
		}
	}
	Klant {
		id: klant
	
	}
}
