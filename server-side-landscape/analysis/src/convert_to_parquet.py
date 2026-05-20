from pathlib import Path

import polars as pl
from polars._typing import SchemaDict

from src.utils.load import load_df_from_json_pickles, load_df_from_pickles

DATA_DIR = Path(__file__).parent
OUTPUT_DIR = Path(__file__).parent.parent / "data" / "measurements"

folder_definitions: list[tuple[str, SchemaDict | None, str | None]] = [
    ("ech_cipher_suites", {"timestamp": pl.Date}, None),
    ("domain_sourcing", None, None),
    # Due to export issues, the data for adaptation is stored as json
    # ("ech_adaptation", {"date": pl.Date}, None),
    ("ech_deactivation", {"current_test_date": pl.Date}, None),
    ("ech_cohort_size", {"test_date": pl.Date}, "*"),
    ("ech_kem", None, None),
    ("ech_public_name", {"test_date": pl.Date}, None),
]


def save_df_to_parquet(df: pl.DataFrame, output_folder: Path, name: str):
    df.write_parquet(output_folder / f"{name}.parquet")


def pickles_to_parquet(
    folder: Path,
    file_endings: str,
    schema_overrides: SchemaDict | None = None,
):
    df = pl.DataFrame(
        load_df_from_pickles(folder.glob(file_endings if file_endings is not None else "*.pickle")),
        schema_overrides=schema_overrides,
    )
    save_df_to_parquet(df, OUTPUT_DIR, folder.name)


def json_to_parquet(
    folder: Path,
    file_endings: str,
    schema_overrides: SchemaDict | None = None,
):
    df = pl.DataFrame(
        load_df_from_json_pickles(folder.glob(file_endings if file_endings is not None else "*.json")),
        schema_overrides=schema_overrides,
    )
    save_df_to_parquet(df, folder.parent, folder.name)


if __name__ == "__main__":
    for (
        folder,
        schema_overrides,
        file_endings,
    ) in folder_definitions:
        pickles_to_parquet(
            DATA_DIR / folder,
            schema_overrides=schema_overrides,
            file_endings=file_endings,
        )

    json_to_parquet(DATA_DIR / "ech_adaptation", "*.pickle", {"date": pl.Date})
