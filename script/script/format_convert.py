import json

TARGET_FILES = ['/Users/ryotaroakagawa/PearOpenDataProject_Ver2/script/data/train_annotations_multi.json',
                '/Users/ryotaroakagawa/PearOpenDataProject_Ver2/script/data/validation_annotations_multi.json']
ANNOTATION_FILE = '/Users/ryotaroakagawa/PearOpenDataProject_Ver2/script/data/annotation.json'


def main():
    json_data = {"images": []}
    for file in TARGET_FILES:
        with open(file) as f:
            multi_data = json.load(f)
            for image in multi_data["images"]:
                json_data['images'].append(image)

    with open(ANNOTATION_FILE) as f:
        data = json.load(f)

    # print(data)

    pear_list = []
    id_list1 = []
    id_list2 = []

    for pears in data['pears']:
        anno_id = pears['id']
        str_anno_id = str(anno_id)
        if not str_anno_id in id_list1:
            id_list1.append(str_anno_id)
            pear_list.append(pears)

    for jsn_image in json_data['images']:
        filename = jsn_image['file_name']
        filename_split = filename.strip().split('_')

        # pear_id のフォーマットが year . ripe . image_id
        pear_id = multi_data['info']['year'] + '00' + filename_split[0].zfill(3)

        if not pear_id in id_list1:
            id_list2.append(pear_id)

            pear_list.append(
                {'id': int(pear_id), 'year': multi_data['info']['year'], 'shape_type_id': 10,
                 'grading_id_deterioration': 0,
                 'grading_id': 0, 'images': [], 'comment': "等級判定の実施なし"})

        if id_list2 == []:
            pass
        else:
            pear_list[id_list2.index(pear_id)]['images'].append(jsn_image)

    pear_dict = {'pears': pear_list}

    with open('/Users/ryotaroakagawa/PearOpenDataProject_Ver2/script/data/annotation_original.json',
              'w') as f:
        json.dump(pear_dict, f, indent=4, ensure_ascii=False)


if __name__ == '__main__':
    main()
