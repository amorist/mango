# Workflow to create distribution archive
    workflow "Create Archive" {
        on = "push"
        resolves = ["archive"]
    }

    # Filter for tag
    action "tag" {
        uses = "actions/bin/filter@master"
        args = "tag"
    }

    # Create Release ZIP archive
    action "archive" {
        uses = "lubusIN/actions/archive@master"
        env = {
                ZIP_FILENAME = "archive-filename"
            }
    }
