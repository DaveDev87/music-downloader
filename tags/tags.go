package tags

// tag, err := id3v2.Open("./chenzo2.mp3", id3v2.Options{Parse: true})
// if err != nil {
// 	log.Fatal("Error while opening mp3 file: ", err)
// }
// defer tag.Close()

// artwork, err := ioutil.ReadFile("zanki-album.jpg")
// if err != nil {
// 	log.Fatal("Error while reading artwork file", err)
// }

// pic := id3v2.PictureFrame{
// 	Encoding:    id3v2.EncodingUTF8,
// 	MimeType:    "image/jpeg",
// 	PictureType: id3v2.PTFrontCover,
// 	Description: "Front cover",
// 	Picture:     artwork,
// }
// tag.AddAttachedPicture(pic)

// // tag.SetArtist("Zutomayo")
// // tag.SetTitle("Zanki")
// // tag.SetAlbum("Zanki [single]")
// // tag.SetGenre("Pop")
// // tag.SetYear("2022")

// // fmt.Println(tag.Artist())
// // fmt.Println(tag.Title())
// // fmt.Println(tag.Album())
// // fmt.Println(tag.Genre())
// // fmt.Println(tag.Year())

// if err := tag.Save(); err != nil {
// 	log.Fatal("Error while saving a tag: ", err)
// }
