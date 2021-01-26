import * as React from "react";
import { ColDef } from "@material-ui/data-grid";
import {
    Button,
    Checkbox,
    makeStyles,
    Table,
    TableBody,
    TableCell,
    TableContainer,
    TableFooter,
    TablePagination,
    TableRow,
    Typography,
} from "@material-ui/core";
import Paper from "@material-ui/core/Paper";
import PlayArrowIcon from "@material-ui/icons/PlayArrow";
import TablePaginationActions from "./Pagination";
import RemoveSelectedButton from "./RemoveSelectedButton";
import "./VideoTable.css";
import { override } from "../../connections/BackendConnection";

const columns: ColDef[] = [
    { field: "youtubeID", headerName: "ID", width: 130, align: "left" },
    { field: "title", headerName: "Name", width: 500, align: "left" },
    { field: "addedBy", headerName: "Added By", width: 130, align: "left" },
];

const useStyles = makeStyles({
    removeAllButtonContainer: {
        display: "flex",
        flexDirection: "row",
        justifyContent: "space-between",
        alignItems: "center",
        margin: "10px",
        minWidth: "265px",
    },
});

const VideoTable = ({
    videos,
    reloadVideos,
}: {
    videos: any;
    reloadVideos: any;
}) => {
    const rows = videos;
    const [page, setPage] = React.useState(0);
    const [rowsPerPage, setRowsPerPage] = React.useState(5);
    const [selected, setSelected] = React.useState(new Set<number>());
    const classes = useStyles();
    const numRows =
        rowsPerPage !== -1
            ? Math.min(rowsPerPage, rows.length - page * rowsPerPage)
            : rows.length;
    const emptyRows =
        rowsPerPage - Math.min(rowsPerPage, rows.length - page * rowsPerPage);
    const allSelected = numRows > 0 && selected.size === numRows;
    const currentRows =
        rowsPerPage > 0
            ? rows.slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage)
            : rows;

    const handleChangePage = (
        event: React.MouseEvent<HTMLButtonElement> | null,
        newPage: number
    ) => {
        setPage(newPage);
        setSelected(new Set());
    };

    const handleChangeRowsPerPage = (
        event: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
    ) => {
        setRowsPerPage(parseInt(event.target.value, 10));
        setPage(0);
        setSelected(new Set());
    };

    const handleSelectAll = () => {
        if (allSelected) {
            setSelected(new Set());
        } else {
            setSelected(new Set(currentRows.map((rw: any) => rw.id)));
        }
    };

    const handleOverride = (id: number) => {
        override(id);
    };

    return (
        <Paper>
            <div className={classes.removeAllButtonContainer}>
                <Typography variant="h5">Videos</Typography>
                <RemoveSelectedButton
                    selected={selected}
                    reloadVideos={reloadVideos}
                />
            </div>
            <div className="table-container">
                <Table aria-label="custom pagination table">
                    <TableBody>
                        <TableRow>
                            <TableCell padding="checkbox">
                                <Checkbox
                                    indeterminate={
                                        selected.size > 0 &&
                                        selected.size < numRows
                                    }
                                    checked={allSelected}
                                    onChange={handleSelectAll}
                                />
                            </TableCell>
                            {columns.map(col => (
                                <TableCell
                                    key={col.headerName}
                                    style={{ width: col.width }}
                                    align={col.align}
                                >
                                    {col.headerName}
                                </TableCell>
                            ))}
                            <TableCell align="left"></TableCell>
                        </TableRow>
                        {currentRows.map((row: any) => {
                            const isSelected = selected.has(row.id);
                            const handleSelect = () => {
                                if (!isSelected) {
                                    const newSelected = new Set(
                                        selected.add(row.id)
                                    );
                                    setSelected(newSelected);
                                } else {
                                    const newSelected = new Set(selected);
                                    newSelected.delete(row.id);
                                    setSelected(newSelected);
                                }
                            };
                            return (
                                <TableRow key={row.title}>
                                    <TableCell padding="checkbox">
                                        <Checkbox
                                            checked={isSelected}
                                            onChange={handleSelect}
                                        />
                                    </TableCell>
                                    {columns.map(col => (
                                        <TableCell
                                            key={col.field}
                                            style={{ width: col.width }}
                                            align={col.align}
                                        >
                                            {
                                                // @ts-ignore
                                                row[col.field]
                                            }
                                        </TableCell>
                                    ))}
                                    <TableCell align="left">
                                        <Button
                                            variant="contained"
                                            color="primary"
                                            onClick={() => {
                                                handleOverride(row.id);
                                            }}
                                            startIcon={<PlayArrowIcon />}
                                        >
                                            Play
                                        </Button>
                                    </TableCell>
                                </TableRow>
                            );
                        })}
                        {emptyRows > 0 && (
                            <TableRow style={{ height: 53 * emptyRows }}>
                                <TableCell colSpan={6} />
                            </TableRow>
                        )}
                    </TableBody>
                    <TableFooter>
                        <TableRow>
                            <TablePagination
                                rowsPerPageOptions={[
                                    5,
                                    10,
                                    25,
                                    { label: "All", value: -1 },
                                ]}
                                colSpan={5}
                                count={rows.length}
                                rowsPerPage={rowsPerPage}
                                page={page}
                                SelectProps={{
                                    inputProps: {
                                        "aria-label": "rows per page",
                                    },
                                    native: true,
                                }}
                                onChangePage={handleChangePage}
                                onChangeRowsPerPage={handleChangeRowsPerPage}
                                ActionsComponent={TablePaginationActions}
                            />
                        </TableRow>
                    </TableFooter>
                </Table>
            </div>
        </Paper>
    );
};

export default VideoTable;
