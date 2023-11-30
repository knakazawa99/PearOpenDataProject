import json

ANNOTATION_FILE = '/Users/ryotaroakagawa/PearOpenDataProject_Ver2/script/data/annotation.json'


def main():
    with open(ANNOTATION_FILE) as f:
        data = json.load(f)

    img_list = []
    anno_list = []
    pear_id_list = []

    for pears in data['pears']:
        img_id = pears['id']
        images = pears['images']
        str_img_id = str(img_id)
        if not str_img_id in pear_id_list:
            pear_id_list.append(str_img_id)
            img_list.append(images)

    new_img_list = [item for sublist in img_list for item in sublist]

    id_list = [item['id'] for item in new_img_list]

    for annotations in data['annotations']:
        anno_id = annotations['image_id']

        if anno_id in id_list:
            anno_list.append(annotations)

    coco_dict = {'images': new_img_list, 'annotations': anno_list, 'categories': data['categories']}

    with open('/Users/ryotaroakagawa/PearOpenDataProject_Ver2/script/data/annotation_coco.json',
              'w') as f:
        json.dump(coco_dict, f, indent=4, ensure_ascii=False)


if __name__ == '__main__':
    main()
