import json

out = []

with open('raw.json') as fi:
    data = json.load(fi)
    for v in data:
        out.append({
            'id': v['snippet']['resourceId']['videoId'],
            'title': v['snippet']['title'],
        })

with open('ost.json', 'w') as fo:
    json.dump(out, fo, indent=2)
