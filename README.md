# Import NotePlan notes into Bear App

Uses the clipboard to import notes into Bear using its `x-callback-url` protocol.

Recursively imports all notes, adding an import tag at the top. The tag is based on the note's directory, so a note called `Drafts/Blog/important_thing.md` will have the tag `#_noteplan/drafts/blog` inserted below the note title.

It ignores folders that start with `@`, which includes the `@Archive`, so you should un-archive any notes you want to transfer.

It runs in the NotePlan `Notes` folder, which is separate from the `Calendar` folder (they are siblings), so your day notes won't be imported. I haven't decided what to do with those, especially since any attachments seem to be in subdirectories based on the note name.

Note: This process **does not check for duplicates!** So if you run it more than once, you will get multiple copies of the notes created. At least you can select and delete all the notes under `_noteplan` if you want to run it again.