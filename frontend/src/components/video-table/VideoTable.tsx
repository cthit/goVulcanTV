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
} from "@material-ui/core";
import Paper from "@material-ui/core/Paper";
import PlayArrowIcon from "@material-ui/icons/PlayArrow";
import TablePaginationActions from "./Pagination";
import RemoveSelectedButton from "./RemoveSelectedButton";

const columns: ColDef[] = [
    { field: "id", headerName: "ID", width: 130, align: "left" },
    { field: "name", headerName: "Name", width: 500, align: "left" },
    { field: "addedBy", headerName: "Added By", width: 130, align: "right" },
];

const rows = [
    {
        id: "99_Abbuf3cQ",
        name: "Australia's Bushfire-Hunting Satellites",
        description: "dasdasda",
        addedBy: "Swexbe",
    },
    {
        id: "zHL9GP_B30E",
        name: "Illusions of Time",
        description: "dasdasda",
        addedBy: "Vidde",
    },
    {
        id: "6w3wr691uss",
        name: "Earth's Deadliest [Computer] Virus",
        description: "dasdasda",
        addedBy: "Santa",
    },
];

const useStyles = makeStyles({
    table: {
        minWidth: 500,
    },
    removeAllButtonContainer: {
        display: "flex",
        flexDirection: "column",
        alignItems: "flex-end",
        margin: "10px",
    },
});

const VideoTable = () => {
    const [page, setPage] = React.useState(0);
    const [rowsPerPage, setRowsPerPage] = React.useState(5);
    const [selected, setSelected] = React.useState(new Set<string>());
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
            setSelected(new Set(currentRows.map(rw => rw.id)));
        }
    };

    return (
        <TableContainer component={Paper}>
            <div className={classes.removeAllButtonContainer}>
                <RemoveSelectedButton
                    selected={selected}
                    setSelected={setSelected}
                />
            </div>
            <Table
                className={classes.table}
                aria-label="custom pagination table"
            >
                <TableBody>
                    <TableRow>
                        <TableCell padding="checkbox">
                            <Checkbox
                                indeterminate={
                                    selected.size > 0 && selected.size < numRows
                                }
                                checked={allSelected}
                                onChange={handleSelectAll}
                            />
                        </TableCell>
                        {columns.map(col => (
                            <TableCell
                                style={{ width: col.width }}
                                align={col.align}
                            >
                                {col.headerName}
                            </TableCell>
                        ))}
                        <TableCell align="right"></TableCell>
                    </TableRow>
                    {currentRows.map(row => {
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
                            <TableRow key={row.name}>
                                <TableCell padding="checkbox">
                                    <Checkbox
                                        checked={isSelected}
                                        onChange={handleSelect}
                                    />
                                </TableCell>
                                {columns.map(col => (
                                    <TableCell
                                        style={{ width: col.width }}
                                        align={col.align}
                                    >
                                        {
                                            // @ts-ignore
                                            row[col.field]
                                        }
                                    </TableCell>
                                ))}
                                <TableCell align="right">
                                    <Button
                                        variant="contained"
                                        color="primary"
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
                            colSpan={3}
                            count={rows.length}
                            rowsPerPage={rowsPerPage}
                            page={page}
                            SelectProps={{
                                inputProps: { "aria-label": "rows per page" },
                                native: true,
                            }}
                            onChangePage={handleChangePage}
                            onChangeRowsPerPage={handleChangeRowsPerPage}
                            ActionsComponent={TablePaginationActions}
                        />
                    </TableRow>
                </TableFooter>
            </Table>
        </TableContainer>
    );
};

export default VideoTable;