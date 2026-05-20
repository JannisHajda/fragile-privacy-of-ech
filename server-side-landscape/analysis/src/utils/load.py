import pickle
from pathlib import Path

import pandas as pd


def convert_json_to_dataframe(data: dict[str, list[tuple[str, float]]]) -> pd.DataFrame:
    # Initialize an empty list to store the rows
    rows = [{"date": date, "metric": key, "value": value} for key, values in data.items() for date, value in values]

    # Create a DataFrame from the rows
    df = pd.DataFrame(rows)

    return df


def load_df_from_pickles(files: list[str] | list[Path], folder: Path | None = None) -> pd.DataFrame:
    # Load the dataframes using list comprehension
    if folder is None:
        dfs = [pickle.load(Path(filename).open("rb")) for filename in files]
    else:
        dfs = [pickle.load((folder / filename).open("rb")) for filename in files]

    # Concatenate all the dataframes into a single dataframe
    df = pd.concat(dfs, ignore_index=True)

    return df


def load_df_from_json_pickles(files: list[Path]) -> pd.DataFrame:
    # Load the dataframes using list comprehension
    dfs = [convert_json_to_dataframe(pickle.load(filename.open("rb"))) for filename in files]

    # Concatenate all the dataframes into a single dataframe
    df = pd.concat(dfs, ignore_index=True)

    return df


def load_df_from_pickle(filename: Path) -> pd.DataFrame:
    # Load the dataframe
    df = pickle.load(filename.open("rb"))

    return df
