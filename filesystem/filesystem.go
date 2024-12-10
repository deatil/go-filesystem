package filesystem

import(
    "io"
    "os"
    "errors"

    "github.com/deatil/go-filesystem/filesystem/util"
    "github.com/deatil/go-filesystem/filesystem/config"
    "github.com/deatil/go-filesystem/filesystem/interfaces"
)

/**
 * 文件管理器
 *
 * @create 2021-8-1
 * @author deatil
 */
type Filesystem struct {
    // 适配器
    adapter interfaces.Adapter

    // 配置
    config interfaces.Config
}

// new 文件管理器
func New(adapters interfaces.Adapter, conf ...map[string]any) *Filesystem {
    fs := &Filesystem{
        adapter: adapters,
    }

    if len(conf) > 0{
        fs.config = fs.PrepareConfig(conf[0])
    }

    return fs
}

// 设置配置
func (this *Filesystem) WithConfig(conf interfaces.Config) {
    this.config = conf
}

// 获取配置
func (this *Filesystem) GetConfig() interfaces.Config {
    return this.config
}

// 提前设置配置
func (this *Filesystem) PrepareConfig(settings map[string]any) interfaces.Config {
    conf := config.New(settings)

    return conf
}

// 设置适配器
func (this *Filesystem) WithAdapter(adapters interfaces.Adapter) *Filesystem {
    this.adapter = adapters
    return this
}

// 获取适配器
func (this *Filesystem) GetAdapter() interfaces.Adapter {
    return this.adapter
}

// 获取文件系统
func (this *Filesystem) GetFilesystem() *Filesystem {
    return this
}

// 判断
func (this *Filesystem) Has(path string) bool {
    path = util.NormalizePath(path)

    if len(path) == 0 {
        return false
    }

    return this.adapter.Has(path)
}

// 写入文件
func (this *Filesystem) Write(path string, contents string, conf ...map[string]any) (bool, error) {
    path = util.NormalizePath(path)

    var newConf map[string]any
    if len(conf) > 0 {
        newConf = conf[0]
    }

    configs := this.PrepareConfig(newConf)

    if _, err := this.adapter.Write(path, contents, configs); err != nil {
        return false, err
    }

    return true, nil
}

// 写入数据流
func (this *Filesystem) WriteStream(path string, resource io.Reader, conf ...map[string]any) (bool, error) {
    path = util.NormalizePath(path)

    var newConf map[string]any
    if len(conf) > 0 {
        newConf = conf[0]
    }

    configs := this.PrepareConfig(newConf)

    if _, err := this.adapter.WriteStream(path, resource, configs); err != nil {
        return false, err
    }

    return true, nil
}

// 更新
func (this *Filesystem) Put(path string, contents string, conf ...map[string]any) (bool, error) {
    path = util.NormalizePath(path)

    var newConf map[string]any
    if len(conf) > 0 {
        newConf = conf[0]
    }

    configs := this.PrepareConfig(newConf)

    if this.Has(path) {
        if _, err := this.adapter.Update(path, contents, configs); err != nil {
            return false, err
        }

        return true, nil
    }

    if _, err := this.adapter.Write(path, contents, configs); err != nil {
        return false, err
    }

    return true, nil
}

// 更新数据流
func (this *Filesystem) PutStream(path string, resource io.Reader, conf ...map[string]any) (bool, error) {
    path = util.NormalizePath(path)

    var newConf map[string]any
    if len(conf) > 0 {
        newConf = conf[0]
    }

    configs := this.PrepareConfig(newConf)

    if this.Has(path) {
        if _, err := this.adapter.UpdateStream(path, resource, configs); err != nil {
            return false, err
        }

        return true, nil
    }

    if _, err := this.adapter.WriteStream(path, resource, configs); err != nil {
        return false, err
    }

    return true, nil
}

// 读取并删除
func (this *Filesystem) ReadAndDelete(path string) (any, error) {
    path = util.NormalizePath(path)

    contents, err := this.Read(path)
    if err != nil {
        return nil, err
    }

    this.Delete(path)

    return contents, nil
}

// 更新字符
func (this *Filesystem) Update(path string, contents string, conf ...map[string]any) (bool, error) {
    path = util.NormalizePath(path)

    var newConf map[string]any
    if len(conf) > 0 {
        newConf = conf[0]
    }

    configs := this.PrepareConfig(newConf)

    if _, err := this.adapter.Update(path, contents, configs); err != nil {
        return false, err
    }

    return true, nil
}

// 更新数据流
func (this *Filesystem) UpdateStream(path string, resource io.Reader, conf ...map[string]any) (bool, error) {
    path = util.NormalizePath(path)

    var newConf map[string]any
    if len(conf) > 0 {
        newConf = conf[0]
    }

    configs := this.PrepareConfig(newConf)

    if _, err := this.adapter.WriteStream(path, resource, configs); err != nil {
        return false, err
    }

    return true, nil
}

// 文件到字符
func (this *Filesystem) Read(path string) (string, error) {
    path = util.NormalizePath(path)
    object, err := this.adapter.Read(path)

    if err != nil {
        return "", err
    }

    return object["contents"].(string), nil
}

// 读取成数据流
func (this *Filesystem) ReadStream(path string) (*os.File, error) {
    path = util.NormalizePath(path)
    object, err := this.adapter.ReadStream(path)

    if err != nil {
        return nil, err
    }

    return object["stream"].(*os.File), nil
}

// 重命名
func (this *Filesystem) Rename(path string, newpath string) (bool, error) {
    path = util.NormalizePath(path)
    newpath = util.NormalizePath(newpath)

    if err := this.adapter.Rename(path, newpath); err != nil {
        return false, err
    }

    return true, nil
}

// 复制
func (this *Filesystem) Copy(path string, newpath string) (bool, error) {
    path = util.NormalizePath(path)
    newpath = util.NormalizePath(newpath)

    if err := this.adapter.Copy(path, newpath); err != nil {
        return false, err
    }

    return true, nil
}

// 删除
func (this *Filesystem) Delete(path string) (bool, error) {
    path = util.NormalizePath(path)

    if err := this.adapter.Delete(path); err != nil {
        return false, err
    }

    return true, nil
}

// 删除文件夹
func (this *Filesystem) DeleteDir(dirname string) (bool, error) {
    dirname = util.NormalizePath(dirname)
    if dirname == "" {
        return false, errors.New("文件夹路径错误")
    }

    if err := this.adapter.DeleteDir(dirname); err != nil {
        return false, err
    }

    return true, nil
}

// 创建文件夹
func (this *Filesystem) CreateDir(dirname string, conf ...map[string]any) (bool, error) {
    dirname = util.NormalizePath(dirname)

    var newConf map[string]any
    if len(conf) > 0 {
        newConf = conf[0]
    }

    configs := this.PrepareConfig(newConf)

    if _, err := this.adapter.CreateDir(dirname, configs); err != nil {
        return false, err
    }

    return true, nil
}

// 列表
func (this *Filesystem) ListContents(dirname string, recursive ...bool) ([]map[string]any, error) {
    dirname = util.NormalizePath(dirname)

    result, err := this.adapter.ListContents(dirname, recursive...)
    if err != nil {
        return nil, err
    }

    return result, nil
}

// 类型
func (this *Filesystem) GetMimetype(path string) (string, error) {
    path = util.NormalizePath(path)
    object, err := this.adapter.GetMimetype(path)

    if err != nil {
        return "", err
    }

    return object["mimetype"].(string), nil
}

// 时间戳
func (this *Filesystem) GetTimestamp(path string) (int64, error) {
    path = util.NormalizePath(path)
    object, err := this.adapter.GetTimestamp(path)

    if err != nil {
        return 0, err
    }

    return object["timestamp"].(int64), nil
}

// 权限
func (this *Filesystem) GetVisibility(path string) (string, error) {
    path = util.NormalizePath(path)
    object, err := this.adapter.GetVisibility(path)

    if err != nil {
        return "", err
    }

    return object["visibility"], nil
}

// 大小
func (this *Filesystem) GetSize(path string) (int64, error) {
    path = util.NormalizePath(path)
    object, err := this.adapter.GetSize(path)

    if err != nil {
        return 0, err
    }

    return object["size"].(int64), nil
}

// 设置权限
func (this *Filesystem) SetVisibility(path string, visibility string) (bool, error) {
    path = util.NormalizePath(path)

    if _, err := this.adapter.SetVisibility(path, visibility); err != nil {
        return false, err
    }

    return true, nil
}

// 信息数据
func (this *Filesystem) GetMetadata(path string) (map[string]any, error) {
    path = util.NormalizePath(path)

    if info, err := this.adapter.GetMetadata(path); err != nil {
        return nil, err
    } else {
        return info, nil
    }
}

// 获取
// Get("file.txt").(*fllesystem.File).Read()
// Get("/file").(*fllesystem.Directory).Read()
func (this *Filesystem) Get(path string, handler ...func(*Filesystem, string) any) any {
    path = util.NormalizePath(path)

    if len(handler) > 0 {
        return handler[0](this, path)
    }

    data, _ := this.GetMetadata(path)

    if data != nil && data["type"] == "file" {
        file := &File{}
        file.SetFilesystem(this)
        file.SetPath(path)

        return file
    } else {
        dir := &Directory{}
        dir.SetFilesystem(this)
        dir.SetPath(path)

        return dir
    }
}
