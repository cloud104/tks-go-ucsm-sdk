#!/usr/bin/env python3
"""Extrai metadados de classes do ucsmsdk para JSON (um arquivo por classe).

Uso:
    python3 extract.py --classes ../classes.txt --out ../../../schema
"""

import argparse
import json
import sys
from importlib.metadata import version as pkg_version
from pathlib import Path

from ucsmsdk.ucscoremeta import MoPropertyMeta
from ucsmsdk.ucscoreutils import load_class

# Mapa número -> texto montado a partir das constantes do próprio SDK,
# para não depender dos valores numéricos.
_ACCESS_LABELS = {
    "NAMING": "naming",
    "CREATE_ONLY": "create-only",
    "READ_ONLY": "ro",
    "READ_WRITE": "rw",
    "INTERNAL": "internal",
}
ACCESS = {getattr(MoPropertyMeta, k): v for k, v in _ACCESS_LABELS.items()}


def read_classes(path: Path) -> list[str]:
    classes = []
    for line in path.read_text().splitlines():
        line = line.strip()
        if line and not line.startswith("#"):
            classes.append(line)
    return classes


def prop_to_dict(p) -> dict:
    if p.access not in ACCESS:
        raise ValueError(f"nível de acesso desconhecido: {p.access} ({p.xml_attribute})")

    d = {
        "xml": p.xml_attribute,
        "type": p.field_type,
        "access": ACCESS[p.access],
        "version": str(p.version) if p.version else None,
    }

    r = p.restriction
    if r is not None:
        if r.value_set:
            d["values"] = sorted(r.value_set)
        if r.range_val:
            d["ranges"] = sorted(r.range_val)
        if r.pattern:
            d["pattern"] = r.pattern
        if r.min_length is not None:
            d["min_length"] = r.min_length
        if r.max_length is not None:
            d["max_length"] = r.max_length
    return d


def class_to_dict(class_id: str, sdk_version: str) -> dict:
    try:
        cls = load_class(class_id)
    except Exception as e:  # o SDK levanta tipos variados
        raise SystemExit(f"erro ao carregar a classe '{class_id}': {e}")
    if cls is None:
        raise SystemExit(f"classe não encontrada no ucsmsdk: '{class_id}'")

    meta = cls.mo_meta
    return {
        "class": meta.xml_attribute,
        "rn": meta.rn,
        "version": str(meta.version) if meta.version else None,
        "parents": sorted(meta.parents),
        "children": sorted(meta.children),
        "sdk_version": sdk_version,
        "props": sorted(
            (prop_to_dict(p) for p in cls.prop_meta.values()),
            key=lambda d: d["xml"],
        ),
    }


def main() -> int:
    ap = argparse.ArgumentParser(description=__doc__)
    ap.add_argument("--classes", required=True, type=Path)
    ap.add_argument("--out", required=True, type=Path)
    args = ap.parse_args()

    sdk_version = pkg_version("ucsmsdk")
    args.out.mkdir(parents=True, exist_ok=True)

    for class_id in read_classes(args.classes):
        data = class_to_dict(class_id, sdk_version)
        dest = args.out / f"{data['class']}.json"
        dest.write_text(json.dumps(data, indent=2, sort_keys=True, ensure_ascii=False) + "\n")
        print(f"{data['class']}: {len(data['props'])} props -> {dest}")

    return 0


if __name__ == "__main__":
    sys.exit(main())