import * as React from "react";
import DeleteIcon from "@material-ui/icons/Delete";
import {
    Button,
    CircularProgress,
    Dialog,
    DialogActions,
    DialogContent,
    DialogContentText,
    DialogTitle,
    useMediaQuery,
    useTheme,
} from "@material-ui/core";

const RemoveSelectedButton = ({
    selected,
    setSelected,
}: {
    selected: Set<string>;
    setSelected: React.Dispatch<React.SetStateAction<Set<string>>>;
}) => {
    const [open, setOpen] = React.useState(false);
    const [loading, setLoading] = React.useState(false);
    const theme = useTheme();
    const fullScreen = useMediaQuery(theme.breakpoints.down("sm"));
    const handleClickOpen = () => {
        setOpen(true);
    };

    const handleClose = () => {
        if (!loading) {
            setOpen(false);
        }
    };

    const handleDelete = () => {
        setLoading(true);
    };
    return (
        <div>
            <Button
                color="secondary"
                variant="contained"
                startIcon={<DeleteIcon />}
                onClick={handleClickOpen}
                disabled={selected.size === 0}
            >
                Remove Selected
            </Button>
            <Dialog
                fullScreen={fullScreen}
                open={open}
                onClose={handleClose}
                aria-labelledby="responsive-dialog-title"
            >
                <DialogTitle id="responsive-dialog-title">
                    {"Delete Videos"}
                </DialogTitle>
                <DialogContent>
                    <DialogContentText>
                        Are you sure you want to delete <b>{selected.size}</b>{" "}
                        videos?
                    </DialogContentText>
                </DialogContent>
                <div
                    style={{
                        height: 50,
                        display: "flex",
                        flexDirection: "column",
                        alignItems: "center",
                    }}
                >
                    {loading && <CircularProgress />}
                </div>
                <DialogActions>
                    <Button
                        autoFocus
                        onClick={handleClose}
                        color="primary"
                        variant="outlined"
                        disabled={loading}
                    >
                        Cancel
                    </Button>
                    <Button
                        onClick={handleDelete}
                        color="secondary"
                        variant="contained"
                        autoFocus
                    >
                        Delete
                    </Button>
                </DialogActions>
            </Dialog>
        </div>
    );
};

export default RemoveSelectedButton;
